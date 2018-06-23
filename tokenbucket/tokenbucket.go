package tokenbucket

import (
	"fmt"
)

type TokenBucket interface {
	GetNextNewTokenTime() int64
	GetAvailableTokens() int64
	ConsumeOneToken() error
	Consume(tokens int64) error
}

type StandardTokenBucket struct {
	timer Timer
	metric Metric
	availableTokens int64
	timePerToken int64
	nextNewTokenTime int64
	burstTokens int64
}

func NewStandardTokenBucket(timer Timer, metric Metric, timePerToken int64, burstTokens int64) TokenBucket {
	return &StandardTokenBucket{
		timer:            timer,
		metric:           metric,
		availableTokens:  burstTokens,
		timePerToken:     timePerToken,
		nextNewTokenTime: 0,
		burstTokens:      burstTokens,
	}
}

func (tb *StandardTokenBucket) GetNextNewTokenTime() int64 {
	return tb.nextNewTokenTime
}

func (tb *StandardTokenBucket) GetAvailableTokens() int64 {
	return tb.availableTokens
}

// Consume one token
func (tb *StandardTokenBucket) ConsumeOneToken() error {
	return tb.Consume(1)
}

// Consume specified number of tokens
func (tb *StandardTokenBucket) Consume(tokens int64) error {
	if tokens <= 0 {
		return fmt.Errorf("tokens to consume must be a positive integer")
	}

	// updates token bucket with new tokens according
	// to elapsed time and check if there are enough tokens
	// to consume specified amount
	tb.update()
	if tokens > tb.availableTokens {
		return fmt.Errorf("not enough available tokens to consume")
	}

	// consume tokens
	tb.availableTokens -= tokens

	return nil
}

// Update available tokens
func (tb *StandardTokenBucket) update() {
	// get number of created tokens since last update
	time := tb.timer.Now()
	tokensCreated := tb.getTokensCreated(time)
	if tokensCreated == 0 {
		return
	}

	// update number of available tokent
	newAvailableTokens := tb.availableTokens + tokensCreated
	if newAvailableTokens > tb.burstTokens {
		newAvailableTokens = tb.burstTokens
	}
	tb.availableTokens = newAvailableTokens

	// set next time when the next (new) token is available
	tb.nextNewTokenTime += (tokensCreated+1)*tb.timePerToken
}

func (tb *StandardTokenBucket) getTokensCreated(time int64) int64 {
	if time < tb.nextNewTokenTime {
		return 0
	}

	elapsedTicks := tb.metric.Diff(tb.nextNewTokenTime, time)
	tokensCreated := elapsedTicks/tb.timePerToken

	return tokensCreated
}
