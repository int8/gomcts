package gomcts

// GameResult - number representing a game result
type GameResult float64

type Action interface {
	ApplyTo(GameState) GameState
}

// GameState - state of the game interface
type GameState interface {
	EvaluateGame() (GameResult, bool)
	GetLegalActions() []Action
	IsGameEnded() bool
	NextToMove() int8
}
