package gomcts

type NodeStatistics struct {
	Q int8
	N int8
}

type GameFinisher interface {
	IsGameOver() bool
}

type GameMover interface {
	Move() GameState
}

type GameResultCalculator interface {
	GameResult() int8
}

type GameState interface {
	GameMover
	GameFinisher
	GameResultCalculator
}

type StatisticsHolder interface {
	GetStatistics() NodeStatistics
}

type Expander interface {
	Expand()
}

type PlayoutPerformer interface {
	Playout()
}

type Backpropagator interface {
	Backpropagate()
}

type MCTSNode interface {
	StatisticsHolder
	Expander
	PlayoutPerformer
	Backpropagator
}

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
