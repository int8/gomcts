package gomcts

// MonteCarloTreeSearchGameNode - MCTS tree node struct
type MonteCarloTreeSearchGameNode struct {
	parent   *MonteCarloTreeSearchGameNode
	value    GameState
	children []MonteCarloTreeSearchGameNode
}

// newMCTSNode - function initializing new MonteCarloTreeSearchGameNode
func newMCTSNode(parentNode MonteCarloTreeSearchGameNode, state GameState) MonteCarloTreeSearchGameNode {
	node := MonteCarloTreeSearchGameNode{nil, state, make([]MonteCarloTreeSearchGameNode, 0, 0)}
	return node
}
