package gomcts

// GameResult - number representing a game result
type GameResult int8

// RolloutPolicy - function signature determining the next game state during Monte Carlo Tree Search rollout
type RolloutPolicy func(GameState) GameState

// GameState - state of the game interface
type GameState interface {
	EvaluateGame() (GameResult, bool)
	GetLegalNextGameStates() []GameState
	IsGameEnded() bool
}

// MCTSNode - Monte Carlo Tree Search node interface
type MCTSNode interface {
	GameState() GameState
	UCTBestChild(float64) (child MCTSNode, isLeaf bool)
	Rollout(RolloutPolicy) GameResult
	Backpropagate(GameResult)
	IsFullyExpanded() bool
	IsTerminal() bool
}

func BestNextGameState(node MCTSNode, n int) GameState {
	for i := 0; i < n; i++ {
		v := TreePolicy(node)
		reward := v.Rollout(DefaultRolloutPolicy)
		v.Backpropagate(reward)
	}
	best, _ := node.UCTBestChild(0.0)
	return best.GameState()
}

func TreePolicy(node MCTSNode) MCTSNode {
	for {
		if node.IsTerminal() {
			return node
		}
		node, leaf := node.UCTBestChild(1.4)
		if leaf {
			return node
		}
	}
}
