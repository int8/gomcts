package gomcts

// TicTacToeBoardGameAction - action on a board game
type TicTacToeBoardGameAction struct {
	xCoord uint8
	yCoord uint8
	value  int8
}

func (a TicTacToeBoardGameAction) Apply(s TicTacToeGameState) TicTacToeGameState {
	if s.nextToMove != a.value {
		panic("*hands slapped*,  not your turn")
	}

	if !s.legalActions[a] {
		panic("*hands slapped*,  action illegal")
	}

	s.board[a.xCoord][a.yCoord] = a.value
	return s
}

type TicTacToeGameState struct {
	nextToMove   int8
	board        [][]int8
	legalActions map[TicTacToeBoardGameAction]bool
}

func CreateTicTacToeInitialGameState(n uint8) TicTacToeGameState {

	legalActions := make(map[TicTacToeBoardGameAction]bool)
	board := make([][]int8, n)
	for i := range board {
		board[i] = make([]int8, n)
		for j := range board[i] {
			legalActions[TicTacToeBoardGameAction{xCoord: uint8(i), yCoord: uint8(j), value: 0}] = true
			legalActions[TicTacToeBoardGameAction{xCoord: uint8(i), yCoord: uint8(j), value: 1}] = true
			legalActions[TicTacToeBoardGameAction{xCoord: uint8(i), yCoord: uint8(j), value: -1}] = true
		}
	}
	state := TicTacToeGameState{nextToMove: 1, board: board, legalActions: legalActions}
	return state
}

func (s TicTacToeGameState) Move(a TicTacToeBoardGameAction) TicTacToeGameState {
	nextState := a.Apply(s)
	nextState.nextToMove *= -1
	delete(nextState.legalActions, a)
	return nextState
}

func (s TicTacToeGameState) EvaluateGame() (GameResult, bool) {
	return GameResult(0), true
}

func (s TicTacToeGameState) GetLegalActions() []TicTacToeBoardGameAction {
	actions := make([]TicTacToeBoardGameAction, 0, len(s.legalActions))
	for k := range s.legalActions {
		actions = append(actions, k)
	}
	return actions
}
