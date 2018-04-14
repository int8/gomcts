package gomcts

// addChild - adding a node to given MCTS Node
func (node *MonteCarloTreeSearchGameNode) addChild(child *MonteCarloTreeSearchGameNode) {
	node.children = append(node.children, *child)
}

// GetParent - adding a node to given MCTS Node
func (node *MonteCarloTreeSearchGameNode) getParent() *MonteCarloTreeSearchGameNode {
	return node.parent
}
