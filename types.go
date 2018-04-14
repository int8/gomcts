package gomcts

// GameResult - number representing a game result
type GameResult int8

// NodeStatistics - statistics gathered while MCTS simulations
type NodeStatistics struct {
	Q float64
	N int64
}

// Action - action to be performed
type Action interface {
	Apply(GameState) GameState
}

// GameState - state of the game interface
type GameState interface {
	Move(Action) GameState
	EvaluateGame() (GameResult, bool)
	GetLegalActions() []Action
}

// MCTSNode - Monte Carlo Tree Search node interface
type MCTSNode interface {
	GetStatistics() NodeStatistics
	BestChild() []MCTSNode
	Rollout(RolloutPolicy) GameResult
	Backpropagate()
}

// RolloutPolicy - function signature determining the next game state during Monte Carlo Tree Search rollout
type RolloutPolicy func(GameState) GameState

// MonteCarloTreeSearchGameNode - MCTS tree node struct
type MonteCarloTreeSearchGameNode struct {
	parent   *MonteCarloTreeSearchGameNode
	value    GameState
	children []MonteCarloTreeSearchGameNode
}
