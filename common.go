package gomcts

// GameResult - number representing a game result
type GameResult int8

// RolloutPolicy - function signature determining the next game state during Monte Carlo Tree Search rollout
type RolloutPolicy func(GameState) GameState

// GameState - state of the game interface
type GameState interface {
	EvaluateGame() (GameResult, bool)
	GetLegalGameStates() []GameState
}

// MCTSNode - Monte Carlo Tree Search node interface
type MCTSNode interface {
	UCTBestChild() MCTSNode
	Rollout(RolloutPolicy) GameResult
	Backpropagate(GameResult)
}

