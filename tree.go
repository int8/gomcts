package gomcts

// addChild - adding a node to given MCTS Node
func (node *monteCarloTreeSearchGameNode) addChild(child *monteCarloTreeSearchGameNode) {
	node.children = append(node.children, child)
}

func (node *monteCarloTreeSearchGameNode) isRoot() bool {
	return node.parent == nil
}
