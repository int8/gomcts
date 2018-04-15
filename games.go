package gomcts

type TicTacToeGameState struct {
	nextToMove   int8
	board        [][]int8
	emptySquares uint16
}

// TicTacToeBoardGameAction - action on a board game
type TicTacToeBoardGameAction struct {
	xCoord uint8
	yCoord uint8
	value  int8
}

func (a TicTacToeBoardGameAction) Apply(s TicTacToeGameState) TicTacToeGameState {
	s.board = copyBoard(s.board)
	if s.nextToMove != a.value {
		panic("*hands slapped*,  not your turn")
	}

	if s.emptySquares == 0 {
		panic("*hands slapped*,  no empty squares on board")
	}

	if s.board[a.xCoord][a.yCoord] != 0 {
		panic("*hands slapped*,  action illegal - square already occupied")
	}

	s.board[a.xCoord][a.yCoord] = a.value
	return s
}

func CreateTicTacToeInitialGameState(n uint8) TicTacToeGameState {
	board := emptyBoard2D(n)
	state := TicTacToeGameState{nextToMove: 1, board: board, emptySquares: uint16(n) * uint16(n)}
	return state
}

func (s TicTacToeGameState) Move(a TicTacToeBoardGameAction) TicTacToeGameState {
	nextState := a.Apply(s)
	nextState.nextToMove *= -1
	nextState.emptySquares--
	return nextState
}

func (s TicTacToeGameState) EvaluateGame() (GameResult, bool) {
	boardSize := len(s.board)
	rowSums := make([]int16, boardSize)
	colSums := make([]int16, boardSize)
	diag1Sum := int16(0)
	diag2Sum := int16(0)

	for i := range s.board {
		for j := range s.board[i] {
			rowSums[i] += int16(s.board[i][j])
			colSums[j] += int16(s.board[i][j])
		}
		diag1Sum += int16(s.board[i][i])
		diag2Sum += int16(s.board[boardSize - 1 - i][i])
 	}

	if diag1Sum == int16(boardSize) || diag2Sum == int16(boardSize)  {
		return GameResult(1), true
	}

	if diag1Sum == -int16(boardSize) || diag2Sum == -int16(boardSize)  {
		return GameResult(-1), true
	}

 	for i := range s.board {
		if rowSums[i] == int16(boardSize) || colSums[i] == int16(boardSize)  {
			return GameResult(1), true
		}
		if rowSums[i] == -int16(boardSize) || colSums[i] == -int16(boardSize)  {
			return GameResult(-1), true
		}
	}

	if s.emptySquares == 0 {
		return GameResult(0), true
	}

	return GameResult(0), false
}

func (s TicTacToeGameState) GetLegalActions() []TicTacToeBoardGameAction {
	actions := make([]TicTacToeBoardGameAction, 0, s.emptySquares*2)
	for i := range s.board {
		for j := range s.board[i] {
			if s.board[i][j] == 0 {
				actions = append(actions, TicTacToeBoardGameAction{xCoord: uint8(i), yCoord: uint8(j), value: 1})
				actions = append(actions, TicTacToeBoardGameAction{xCoord: uint8(i), yCoord: uint8(j), value: -1})
			}
		}
	}
	return actions
}
