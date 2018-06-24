package tokenbucket

import "testing"

type testClock struct {
	now int64
}

func (t testClock) Now() int64 {
	return t.now
}

func Test_Consume_InitialStateConsumeOne_ReturnBurstTokensMinusOne(t *testing.T) {
	var ticksPerToken int64 = 10
	var burstTokens int64 = 100
	var expectedAvailableTokens int64 = 99

	clock := new(testClock)
	metric := NewTimeMetric()
	bucket := NewStandardTokenBucket(clock, metric, ticksPerToken, burstTokens)

	bucket.Consume(1)

	if bucket.GetAvailableTokens() != expectedAvailableTokens {
		t.Errorf(
			"Number of vailable tokens for token bucken in initial state is incorrect, got: %d, want: %d.",
			bucket.GetAvailableTokens(),
			burstTokens)
	}
}

func Test_Consume_TimeElapsed_ReturnCorrectNumberOfTokensAndNewTime(t *testing.T) {
	var ticksPerToken int64 = 10
	var burstTokens int64 = 100
	var timeDelta int64 = 200
	var expectedAvailableTokens int64 = 30
	var expectedNextNewTokenTime int64 = 210

	clock := new(testClock)
	metric := NewTimeMetric()
	bucket := NewStandardTokenBucket(clock, metric, ticksPerToken, burstTokens)

	bucket.Consume(50)
	clock.now += timeDelta
	bucket.Consume(40)

	// check so that available tokens is correct
	if bucket.GetAvailableTokens() != expectedAvailableTokens {
		t.Errorf(
			"Number of new tokens is incorrect, got: %d, want: %d.",
			bucket.GetAvailableTokens(),
			expectedAvailableTokens)
	}

	// check so that next new token time is correct
	if bucket.GetNextNewTokenTime() != expectedNextNewTokenTime {
		t.Errorf(
			"Expected next new token time is incorrect, got: %d, want: %d.",
			bucket.GetNextNewTokenTime(),
			expectedNextNewTokenTime)
	}
}

func Test_GetNewTokens_TooManyTokensConsumed_Error(t *testing.T) {
	var ticksPerToken int64 = 10
	var burstTokens int64 = 100

	clock := new(testClock)
	metric := NewTimeMetric()
	tb := NewStandardTokenBucket(clock, metric, ticksPerToken, burstTokens)

	err := tb.Consume(120)

	if err == nil {
		t.Errorf("Error should be returned if trying to consume more tokens than available.")
	}
}
