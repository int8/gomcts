package gomcts

func initialize2DInt8Slice(n uint8) [][]int8 {
	board := make([][]int8, n)
	for i := range board {
		board[i] = make([]int8, n)
	}
	return board
}

func copy2DInt8Slice(board [][]int8) [][]int8 {
	newBoard := initialize2DInt8Slice(uint8(len(board)))
	for i := range board {
		for j := range board[i] {
			newBoard[i][j] = board[i][j]
		}
	}
	return newBoard
}
