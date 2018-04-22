package gomcts

import (
	"crypto/rand"
	"math/big"
)

// RolloutPolicy - function signature determining the next action during Monte Carlo Tree Search rollout
type RolloutPolicy func(GameState) Action

// DefaultRolloutPolicy - default rollout policy, picks action randomly (w.r.t uniform random dist)
func DefaultRolloutPolicy(state GameState) Action {
	actions := state.GetLegalActions()
	numberOfActions := int64(len(actions))
	actionIndex, _ := rand.Int(rand.Reader, big.NewInt(numberOfActions))
	return actions[actionIndex.Int64()]
}
