## Monte Carlo Tree Search

Implementation of basic [Monte Carlo Tree Search algorithm](https://int8.io/monte-carlo-tree-search-beginners-guide/). 

Install with 
```bash
go get github.com/int8/gomcts
```

To use it for your sum-zero two players game you need to provide implementation of ```GameState``` and ```Action``` interfaces

```go
// Action - game action interface
type Action interface{
	ApplyTo(GameState) GameState
}

// GameState - state of the game interface
type GameState interface {
	EvaluateGame() (GameResult, bool)
	GetLegalActions() []Action
	IsGameEnded() bool
	NextToMove() int8
}
```

You can use ```DefaultRolloutPolicy``` (actions chosen randomly) or implement your own Rollout Policy as a function with the following signature:

```go
func YourCustomRolloutPolicy(state GameState) Action {
	...
}
```


There is a built-in tic-tac-toe game implementation available through 
 ```TicTacToeBoardGameAction``` and ```TicTacToeGameState``` types

To play with it go for something like:
```go
package main 
import "github.com/int8/gomcts"

func main() {
	initialState := gomcts.CreateTicTacToeInitialGameState(3)
	chosenAction:= gomcts.MonteCarloTreeSearch(initialState, gomcts.DefaultRolloutPolicy, 100)
	// use chosenAction further
}   
``` 
