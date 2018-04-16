package gomcts

import "fmt"


// TicTacToeBoardGameAction - action on a board game
type TicTacToeBoardGameAction struct {
	xCoord uint8
	yCoord uint8
	value  int8
}

func (a TicTacToeBoardGameAction) ApplyTo(s TicTacToeGameState) TicTacToeGameState {
	s.board = copy2DInt8Slice(s.board)
	if s.nextToMove != a.value {
		panic("*hands slapped*,  not your turn")
	}

	if s.emptySquares == 0 {
		panic("*hands slapped*,  no empty squares on board")
	}

	if s.board[a.xCoord][a.yCoord] != 0 {
		fmt.Println(a, s)
		panic("*hands slapped*,  action illegal - square already occupied")
	}

	s.board[a.xCoord][a.yCoord] = a.value
	s.nextToMove *= -1
	s.emptySquares--
	return s
}


type TicTacToeGameState struct {
	nextToMove   int8
	board        [][]int8
	emptySquares uint16
	ended        bool
	result       GameResult
}

func createTicTacToeInitialGameState(boardSize uint8) TicTacToeGameState {
	board := initialize2DInt8Slice(boardSize)
	state := TicTacToeGameState{nextToMove: 1, board: board, emptySquares: uint16(boardSize) * uint16(boardSize)}
	return state
}


func (s TicTacToeGameState) IsGameEnded() bool {
	_, ended := s.EvaluateGame()
	return ended
}

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

func (s TicTacToeGameState) GetLegalNextGameStates() []GameState {
	if !s.IsGameEnded() {
		states := make([]GameState, 0, s.emptySquares)
		for i := range s.board {
			for j := range s.board[i] {
				if s.board[i][j] == 0 {
					states = append(states, TicTacToeBoardGameAction{xCoord: uint8(i), yCoord: uint8(j), value: s.nextToMove}.ApplyTo(s))
				}
			}
		}
		return states
	}
	return nil
}
