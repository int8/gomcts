package gomcts

// TicTacToeBoardGameAction - action on a tic tac toe board game
type TicTacToeBoardGameAction struct {
	xCoord uint8
	yCoord uint8
	value  int8
}

// ApplyTo - TicTacToeBoardGameAction implementation of ApplyTo method of Action interface
func (a TicTacToeBoardGameAction) ApplyTo(s GameState) GameState {
	ticTacToeGameState := s.(TicTacToeGameState)
	ticTacToeGameState.board = copy2DInt8Slice(ticTacToeGameState.board)
	if ticTacToeGameState.nextToMove != a.value {
		panic("*hands slapped*,  not your turn")
	}

	if ticTacToeGameState.board[a.xCoord][a.yCoord] != 0 {
		panic("*hands slapped*,  action illegal - square already occupied")
	}

	ticTacToeGameState.board[a.xCoord][a.yCoord] = a.value
	ticTacToeGameState.nextToMove *= -1
	ticTacToeGameState.emptySquares--
	return ticTacToeGameState
}

// TicTacToeGameState - tic tac toe game state
type TicTacToeGameState struct {
	nextToMove   int8
	board        [][]int8
	emptySquares uint16
	ended        bool
	result       GameResult
}

// CreateTicTacToeInitialGameState - initializes tic tac toe game state
func CreateTicTacToeInitialGameState(boardSize uint8) TicTacToeGameState {
	board := initialize2DInt8Slice(boardSize)
	state := TicTacToeGameState{nextToMove: 1, board: board, emptySquares: uint16(boardSize) * uint16(boardSize)}
	return state
}

// IsGameEnded - TicTacToeGameState implementation of IsGameEnded method of GameState interface
func (s TicTacToeGameState) IsGameEnded() bool {
	_, ended := s.EvaluateGame()
	return ended
}

// EvaluateGame - TicTacToeGameState implementation of EvaluateGame method of GameState interface
func (s TicTacToeGameState) EvaluateGame() (result GameResult, ended bool) {

	defer func() {
		s.result = result
		s.ended = ended
	}()

	if s.ended {
		return s.result, s.ended
	}

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
		diag2Sum += int16(s.board[boardSize-1-i][i])
	}

	if diag1Sum == int16(boardSize) || diag2Sum == int16(boardSize) {
		return GameResult(1), true
	}

	if diag1Sum == -int16(boardSize) || diag2Sum == -int16(boardSize) {
		return GameResult(-1), true
	}

	for i := range s.board {
		if rowSums[i] == int16(boardSize) || colSums[i] == int16(boardSize) {
			return GameResult(1), true
		}
		if rowSums[i] == -int16(boardSize) || colSums[i] == -int16(boardSize) {
			return GameResult(-1), true
		}
	}

	if s.emptySquares == 0 {
		return GameResult(0), true
	}

	return GameResult(0), false
}

// GetLegalActions - TicTacToeGameState implementation of GetLegalActions method of GameState interface
func (s TicTacToeGameState) GetLegalActions() []Action {
	actions := make([]Action, 0, s.emptySquares)
	if !s.IsGameEnded() {
		for i := range s.board {
			for j := range s.board[i] {
				if s.board[i][j] == 0 {
					actions = append(actions, TicTacToeBoardGameAction{xCoord: uint8(i), yCoord: uint8(j), value: s.nextToMove})
				}
			}
		}
	}
	return actions
}

// NextToMove - TicTacToeGameState implementation of NextToMove method of GameState interface
func (s TicTacToeGameState) NextToMove() int8 {
	return s.nextToMove
}
