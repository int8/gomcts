
[![Build Status](https://travis-ci.org/int8/gomcts.svg?branch=master)](https://travis-ci.org/int8/gomcts)
[![GoDoc Reference](https://godoc.org/github.com/int8/gomcts/vm?status.svg)](http://godoc.org/github.com/int8/gomcts)
[![Go Report Card](https://goreportcard.com/badge/github.com/int8/gomcts)](https://goreportcard.com/report/github.com/int8/gomcts)

## Monte Carlo Tree Search

Implementation of basic [Monte Carlo Tree Search algorithm](https://int8.io/monte-carlo-tree-search-beginners-guide/). 

#### Installation 
Install with 
```bash
go get github.com/int8/gomcts
```


#### Usage 

The central routine is ```MonteCarloTreeSearch(GameState, RolloutPolicy, int)``` which consumes GameState, RolloutPolicy and performs requested number of MCTS simulations 

To use it for your perfect-information sum-zero strictly competitive two players game (board games such as go/chess/checkers/tictactoe) you need to provide implementation of ```GameState``` and ```Action``` interfaces

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

where ```GameResult``` is just ```float64``` alias 
```go
type GameResult float64
```

As current implementation expects sum-zero two players game GameResult is supposed to reflect it. For the same reason ```NextToMove()``` is expected to return ```1``` or ```-1```.    

You can use ```DefaultRolloutPolicy``` (actions chosen randomly) or implement your own Rollout Policy as a function with the following signature:

```go
func YourCustomRolloutPolicy(state GameState) Action {
	...
}
```


#### Examples 

##### Tic-Tac-Toe 
There is a built-in tic-tac-toe game implementation [available](https://github.com/int8/gomcts/blob/master/tictactoe.go) through 
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
