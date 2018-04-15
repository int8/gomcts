package gomcts

import "testing"

func TestTicTacToeGameStateInitialization(t *testing.T) {
	state := CreateTicTacToeInitialGameState(3)
	if state.emptySquares != 9 {
		t.Errorf("state.emptySquares should be 9, but it is %v", state.emptySquares)
	}

	if state.nextToMove != 1 {
		t.Errorf("state.nextToMove should be 1, but it is %v", state.nextToMove)
	}

	if len(state.board) != 3 {
		t.Errorf("len(state.board) should be 3, but it is %v", len(state.board))
	}

	if len(state.board[0]) != 3 {
		t.Errorf("len(state.board[0]) should be 3, but it is %v", len(state.board[0]))
	}
}


func TestMoveProducesTicTacToeGameStateCorrectly(t *testing.T) {
	state := CreateTicTacToeInitialGameState(3)
	action := TicTacToeBoardGameAction{xCoord:1, yCoord:1, value:1}
	nextState := state.Move(action)

	if nextState.emptySquares != 8 {
		t.Errorf("state.emptySquares should be 8, but it is %v", state.emptySquares)
	}

	if nextState.nextToMove != -1 {
		t.Errorf("state.nextToMove should be -1, but it is %v", state.nextToMove)
	}

	if &(nextState.board[0][0]) == &(state.board[0][0]) {
		t.Errorf("state.board[0][0] and nextState.board[0][0] refer to the same memory location - but should not")
	}

	if nextState.board[1][1] != 1 {
		t.Errorf("nextState.board[1][1] should be 1 but is %v", nextState.board[1][1])
	}


	if state.board[1][1] != 0 {
		t.Errorf("state.board[1][1] should remain 0 but is %v", state.board[1][1])
	}
}

func TestEmptyTicTacToeGameStateEvaluation(t *testing.T) {
	state := CreateTicTacToeInitialGameState(3)
	_, gameEnded := state.EvaluateGame()
	if gameEnded {
		t.Errorf("Game state is evaluated as ended but should not")
	}
}


func TestLegalActionsOfTicTacToeGameState(t *testing.T) {
	state := CreateTicTacToeInitialGameState(3)
	actions := state.GetLegalActions()
	if len(actions) != 18 {
		t.Errorf("There should be 18 legal actions to perform but is %v", len(actions))
	}

	newState := state.Move(actions[0])
	actions = newState.GetLegalActions()
	if len(actions) != 16 {
		t.Errorf("There should be 16 legal actions to perform after first move but is %v", len(actions))
	}

}
