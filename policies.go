package gomcts

import "math/rand"

// RolloutPolicy - function signature determining the next action during Monte Carlo Tree Search rollout
type RolloutPolicy func(GameState) Action

func DefaultRolloutPolicy(state GameState) Action {
	actions := state.GetLegalActions()
	actionIndex := rand.Intn(len(actions))
	return actions[actionIndex]
}