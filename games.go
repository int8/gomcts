package gomcts

// GameResult - number representing a game result
type GameResult int8

type Action interface{
	ApplyTo(GameState) GameState
}

// GameState - state of the game interface
type GameState interface {
	EvaluateGame() (GameResult, bool)
	GetLegalActions() []Action
	IsGameEnded() bool
}
