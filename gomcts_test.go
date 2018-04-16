package gomcts

import "testing"

func TestTicTacToeGameStateInitialization(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
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
	state := createTicTacToeInitialGameState(3)
	action := TicTacToeBoardGameAction{xCoord:1, yCoord:1, value:1}
	nextState := action.ApplyTo(state)

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
	state := createTicTacToeInitialGameState(3)
	_, gameEnded := state.EvaluateGame()
	if gameEnded {
		t.Errorf("Game state is evaluated as ended but should not")
	}
}


func TestNumberOfLegalActionsOfTicTacToeGameState(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	states := state.GetLegalNextGameStates()
	if len(states) != 9 {
		t.Errorf("There should be 9 actions to perform but is %v", len(states))
	}

	newState := states[0]
	states = newState.GetLegalNextGameStates()
	if len(states) != 8 {
		t.Errorf("There should be 8 legal actions to perform after first move but is %v", len(states))
	}
}

func TestLegalGameStateZeroIfGameEnded(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 4
	state.board = [][]int8 {
		{ 1, 0, 0},
		{-1, 1, 0},
		{ 0,-1, 1},
	}

	states := state.GetLegalNextGameStates()
	if len(states) > 0 {
		t.Errorf("Game is ended but state has legal game states to go to")
	}
}

func TestOutOfBoardMovePanic(t *testing.T) {
	defer func() {
		r := recover();
		if r == nil {
			t.Errorf("The code did not panic but should")
		}
	}()

	state := createTicTacToeInitialGameState(3)
	action := TicTacToeBoardGameAction{xCoord: 4, yCoord: 4, value: 100}
	action.ApplyTo(state)
}


func TestAlreadyOccupiedSquareMovePanic(t *testing.T) {
	defer func() {
		r := recover();
		if r == nil {
			t.Errorf("The code did not panic but should")
		}
	}()

	state := createTicTacToeInitialGameState(3)
	action := TicTacToeBoardGameAction{xCoord: 1, yCoord: 1, value: 1}
	nextState := action.ApplyTo(state)
	action.ApplyTo(nextState)
}


func TestGameEvaluationShouldBeNotEnded(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 1
	state.board = [][]int8 {
		{ 1,  0, -1},
		{-1, -1,  1},
		{ 1,  -1, 1},
	}

	_, ended := state.EvaluateGame()

	if ended {
		t.Errorf("Game should be not ended but is")
	}

	if state.ended {
		t.Errorf("Game should be not ended but is")
	}

}

func TestGameEvaluationShouldBeDraw(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 0
	state.board = [][]int8 {
		{ 1,  1, -1},
		{-1, -1,  1},
		{ 1,  -1, 1},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(0) {
		t.Errorf("Result should be a draw but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}

	if state.ended {
		t.Errorf("Game should be ended but is not")
	}

}


func TestGameEvaluationShouldResultFirstPlayerWinningBecauseOfFirstDiagonal(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 4
	state.board = [][]int8 {
		{ 1, 0, 0},
		{-1, 1, 0},
		{ 0,-1, 1},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(1) {
		t.Errorf("Result should be a 1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}


func TestGameEvaluationShouldResultFirstPlayerWinningBecauseOfSecondDiagonal(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 4
	state.board = [][]int8 {
		{ 0, 0, 1},
		{-1, 1, 0},
		{ 1,-1, 0},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(1) {
		t.Errorf("Result should be a 1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}

func TestGameEvaluationShouldResultSecondPlayerWinningBecauseOfFirstDiagonal(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 3
	state.board = [][]int8 {
		{ -1, 0, 1},
		{ 1, -1, 0},
		{ 0, 1, -1},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(-1) {
		t.Errorf("Result should be a -1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}

func TestGameEvaluationShouldResultSecondPlayerWinningBecauseOfSecondDiagonal(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 3
	state.board = [][]int8 {
		{ 1, 0, -1},
		{1, -1, 0},
		{ -1,1, 0},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(-1) {
		t.Errorf("Result should be a -1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}


func TestGameEvaluationShouldResultFirstPlayerWinningBecauseOfFirstRow(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 4
	state.board = [][]int8 {
		{ 1, 1, 1},
		{-1, 0, 0},
		{ 0,-1, 0},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(1) {
		t.Errorf("Result should be a 1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}

func TestGameEvaluationShouldResultSecondPlayerWinningBecauseOfFirstRow(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 3
	state.board = [][]int8 {
		{ -1, -1, -1},
		{1, 0, 0},
		{ 0,1, 1},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(-1) {
		t.Errorf("Result should be a -1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}



func TestGameEvaluationShouldResultFirstPlayerWinningBecauseOfFirstColumn(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 4
	state.board = [][]int8 {
		{ 1, 0, 0},
		{ 1, 0, -1},
		{ 1,-1, 0},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(1) {
		t.Errorf("Result should be a 1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}

func TestGameEvaluationShouldResultSecondPlayerWinningBecauseOfFirstColumn(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 3
	state.board = [][]int8 {
		{ -1, 1, 1},
		{-1, 0, 0},
		{ -1,1, 0},
	}

	result, ended := state.EvaluateGame()
	if result != GameResult(-1) {
		t.Errorf("Result should be a -1 but is %v", result)
	}

	if !ended {
		t.Errorf("Game should be ended but is not")
	}
}


func TestRolloutTerminates(t *testing.T) {
	defer func() {
		r := recover();
		if r != nil {
			t.Errorf("The code panics: %v", r)
		}
	}()

	state := createTicTacToeInitialGameState(3)
	node := NewMCTSNode(nil, state)
	node.Rollout(DefaultRolloutPolicy)
}


func TestUCTBestChildNewNode(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	node := NewMCTSNode(nil, state)
	if len(node.untriedGameStates) != 9 {
		t.Errorf("There should be 9 untried game states at the beginning of 3x3 tic tac toe game")
	}
	node.UCTBestChild(1.4)

	if len(node.untriedGameStates) != 8 {
		t.Errorf("There should be 8 untried game states after first UCTBestChild for 3x3 tic tac toe game but is %v", len(node.untriedGameStates))
	}
}


func TestUCTBestChildNewNode9Times(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	node := NewMCTSNode(nil, state)
	if len(node.untriedGameStates) != 9 {
		t.Errorf("There should be 9 untried game states at the beginning of 3x3 tic tac toe game")
	}
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)
	node.UCTBestChild(1.4)

	if len(node.untriedGameStates) != 0 {
		t.Errorf("There should be 0 untried game states after first UCTBestChild for 3x3 tic tac toe game but is %v", len(node.untriedGameStates))
	}

	if !node.IsFullyExpanded() {
		t.Errorf("Node should be fully expanded now but is not")
	}

	node.UCTBestChild(1.4)
}


func TestNodeIsTerminal(t *testing.T) {
	state := createTicTacToeInitialGameState(3)
	state.emptySquares = 3
	state.board = [][]int8 {
		{ -1, 1, 1},
		{-1, 0, 0},
		{ -1,1, 0},
	}

	node := NewMCTSNode(nil, state)
	if !node.IsTerminal() {
		t.Errorf("Node should be terminal but is not")
	}
}








