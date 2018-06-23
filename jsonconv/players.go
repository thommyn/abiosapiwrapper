package jsonconv

func NewPlayersFromSeries() JsonConverter {
	return NewRostersSubNode("players")
}
