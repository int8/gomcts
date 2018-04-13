package gomcts

type GameResult int8

type NodeStatistics struct {
	Q float64
	N int64
}

type GameState interface {
	Move() GameState
	IsGameOver() bool
	EvaluateGame() GameResult
}

type MCTSNode interface {
	GetStatistics() NodeStatistics
	Expand() []MCTSNode
	Rollout(RolloutPolicy) GameResult
	Backpropagate()
}

type RolloutPolicy func(GameState) GameResult

type MonteCarloTreeSearchGameNode struct {
	parent   *MonteCarloTreeSearchGameNode
	value    GameState
	children []MonteCarloTreeSearchGameNode
}

func NewMonteCarloTreeSearchGameNode(parentNode MonteCarloTreeSearchGameNode, state GameState) MonteCarloTreeSearchGameNode {
	node := MonteCarloTreeSearchGameNode{nil, state, make([]MonteCarloTreeSearchGameNode, 0, 0)}
	return node
}

func (node *MonteCarloTreeSearchGameNode) AddChild(child *MonteCarloTreeSearchGameNode) {
	node.children = append(node.children, *child)
}
