package gomcts

import (
	"fmt"
	"math"
)

// MonteCarloTreeSearchGameNode - MCTS tree node struct
type MonteCarloTreeSearchGameNode struct {
	parent         *MonteCarloTreeSearchGameNode
	value          GameState
	children       []MonteCarloTreeSearchGameNode
	untriedActions []Action
	q              float64
	n              int64
}

// NewMCTSNode - function initializing new MonteCarloTreeSearchGameNode
func NewMCTSNode(parentNode *MonteCarloTreeSearchGameNode, state GameState) MonteCarloTreeSearchGameNode {
	node := MonteCarloTreeSearchGameNode{parent: parentNode, value: state}
	node.children = make([]MonteCarloTreeSearchGameNode, 0, 0)
	node.untriedActions = state.GetLegalActions()
	return node
}

func (node *MonteCarloTreeSearchGameNode) UCTBestChild(c float64) (child *MonteCarloTreeSearchGameNode, isLeaf bool) {
	if node.IsTerminal() {
		return nil, true
	}

	if node.IsFullyExpanded() {
		chosenIndex := 0
		maxValue := 0.0
		for i, child := range node.children {
			if (child.q/float64(child.n))+c*math.Sqrt(2*math.Log(float64(node.n))/float64(child.n)) > maxValue {
				chosenIndex = i
			}
		}
		return &node.children[chosenIndex], false
	}
	action := node.untriedActions[0]
	node.untriedActions = node.untriedActions[1:]
	childNode := NewMCTSNode(node, action.ApplyTo(node.value))
	node.addChild(&childNode)
	return &childNode, true
}

func (node *MonteCarloTreeSearchGameNode) Rollout(policy RolloutPolicy) GameResult {
	currentState := node.value
	for {
		result, ended := currentState.EvaluateGame()
		if ended {
			return result
		} else {
			currentState = policy(currentState).ApplyTo(currentState)
		}
	}
}

func (node *MonteCarloTreeSearchGameNode) Backpropagate(result GameResult) {
	node.q += float64(result)
	node.n++
	if node.parent != nil {
		node.parent.Backpropagate(result)
	}
}

func (node *MonteCarloTreeSearchGameNode) IsTerminal() bool {
	_, ended := node.value.EvaluateGame()
	return ended
}

func (node *MonteCarloTreeSearchGameNode) IsFullyExpanded() bool {
	return len(node.untriedActions) == 0 && !node.IsTerminal()
}

func BestNextGameState(node *MonteCarloTreeSearchGameNode, n int) GameState {
	for i := 0; i < n; i++ {
		v := TreePolicy(node)
		reward := v.Rollout(DefaultRolloutPolicy)
		v.Backpropagate(reward)
	}
	best, _ := node.UCTBestChild(0.0)
	return best.value
}

func TreePolicy(node *MonteCarloTreeSearchGameNode) *MonteCarloTreeSearchGameNode {
	for {
		if node.IsTerminal() {
			return node
		}
		node, leaf := node.UCTBestChild(1.4)
		fmt.Println(node)
		if leaf {
			return node
		}
	}
}
