package gomcts


func emptyBoard2D(n uint8) [][]int8 {
	board := make([][]int8, n)
	for i := range board {
		board[i] = make([]int8, n)
	}
	return board
}

func copyBoard(board [][]int8) [][]int8 {
	newBoard := emptyBoard2D(uint8(len(board)))
	for i:= range board {
		for j:= range board[i] {
			newBoard[i][i] = board[i][j]
		}
	}
	return newBoard
}

