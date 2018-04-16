package gomcts

import "math/rand"



func DefaultRolloutPolicy(state GameState) GameState {
	states := state.GetLegalGameStates()
	stateIndex := rand.Intn(len(states))
	return states[stateIndex]
}

// MonteCarloTreeSearchGameNode - MCTS tree node struct
type MonteCarloTreeSearchGameNode struct {
	parent   *MonteCarloTreeSearchGameNode
	value    GameState
	children []MonteCarloTreeSearchGameNode
	untriedGameStates []GameState
	q        float64
	n        int64
}

// NewMCTSNode - function initializing new MonteCarloTreeSearchGameNode
func NewMCTSNode(parentNode *MonteCarloTreeSearchGameNode, state GameState) MonteCarloTreeSearchGameNode {
	node := MonteCarloTreeSearchGameNode{parent: parentNode, value: state}
	node.children = make([]MonteCarloTreeSearchGameNode, 0, 0)
	node.untriedGameStates = state.GetLegalGameStates()
	return node
}

func (node MonteCarloTreeSearchGameNode) Rollout(policy RolloutPolicy) GameResult {
	currentState := node.value
	for {
		result, ended := currentState.EvaluateGame()
		if ended {
			return result
		} else {
			currentState = policy(currentState)
		}
	}
}


func (node MonteCarloTreeSearchGameNode) Backpropagate(result GameResult) {
	node.q += float64(result)
	node.n++
	if node.getParent() != nil {
		node.getParent().Backpropagate(result)
	}
}


func (node MonteCarloTreeSearchGameNode) IsTerminal() bool {
	_, ended := node.value.EvaluateGame()
	return ended
}

func (node MonteCarloTreeSearchGameNode) IsFullyExpanded() bool {
	return len(node.untriedGameStates) == 0 && !node.IsTerminal()
}

func (node MonteCarloTreeSearchGameNode) IsLeaf() bool {
	return node.n == 0
}

