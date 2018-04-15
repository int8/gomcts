package gomcts

// MonteCarloTreeSearchGameNode - MCTS tree node struct
type MonteCarloTreeSearchGameNode struct {
	stats NodeStatistics
	parent   *MonteCarloTreeSearchGameNode
	value    GameState
	children []MonteCarloTreeSearchGameNode
}

// newMCTSNode - function initializing new MonteCarloTreeSearchGameNode
func newMCTSNode(parentNode *MonteCarloTreeSearchGameNode, state GameState) MonteCarloTreeSearchGameNode {
	stats := NodeStatistics{N:0, Q:0}
	node := MonteCarloTreeSearchGameNode{stats, parentNode, state, make([]MonteCarloTreeSearchGameNode, 0, 0)}
	return node
}


func (n MonteCarloTreeSearchGameNode) GetStatistics() NodeStatistics {
	return n.stats
}

func (n MonteCarloTreeSearchGameNode) Rollout(policy RolloutPolicy) GameResult {
	currentState := n.value
	for {
		result, ended := currentState.EvaluateGame()
		if ended {
			return result
		} else {
			currentState = policy(currentState)
		}
	}
}

func (n MonteCarloTreeSearchGameNode) Backpropagate(result GameResult)  {
	n.stats.Q += float64(result)
	n.stats.N++
	if n.parent != nil {
		n.parent.Backpropagate(result)
	}
}


