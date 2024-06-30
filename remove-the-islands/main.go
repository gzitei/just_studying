package main

type (
	Board [][]int
	Memo  map[[2]int]bool
)

var memo = Memo{}

func checkDirection(board Board, row, col int, increment map[byte]int) bool {
	if _, ok := memo[[2]int{row, col}]; ok {
	} else if col == 0 || row == 0 || col == len(board[row])-1 || row == len(board)-1 {
		memo[[2]int{row, col}] = board[row][col] == 1
	} else if board[row][col] == 0 {
		memo[[2]int{row, col}] = false
	} else {
		col, row = col+increment['x'], row+increment['y']
		return checkDirection(board, row, col, increment)
	}
	return memo[[2]int{row, col}]
}

func connectsToBorder(board Board, row, col int) bool {
	if v, ok := memo[[2]int{row, col}]; ok {
		return v
	}
	var connects bool
	directions := []map[byte]int{
		{'x': -1, 'y': 0},
		{'x': 1, 'y': 0},
		{'x': 0, 'y': -1},
		{'x': 0, 'y': 1},
	}
	for _, increment := range directions {
		connects = connects || checkDirection(board, row, col, increment)
	}
	memo[[2]int{row, col}] = connects
	return connects
}

func removeIslands(board Board) Board {
	result := make(Board, len(board))
	for row := 0; row < len(board); row++ {
		result[row] = make([]int, len(board[row]))
		for col := 0; col < len(board[row]); col++ {
			connects := connectsToBorder(board, row, col)
			if board[row][col] == 1 && !connects {
				result[row][col] = 0
			} else {
				result[row][col] = board[row][col]
			}
		}
	}
	return result
}

func main() {
}
