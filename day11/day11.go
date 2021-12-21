package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDims() (int, int) {
	return 10, 10
}

func printBoard(board [][]int) {
	xDim := len(board)
	yDim := len(board[0])

	for line := 0; line < yDim; line++ {
		for col := 0; col < xDim; col++ {
			fmt.Printf("%v ", board[col][line])
		}
		fmt.Print("\n")
	}
}

func contains(arr []point, p point) bool {
	for _, v := range arr {
		if v == p {
			return true
		}
	}
	return false
}

func parseInputFile(path string) [][]int {
	xDim, yDim := getDims()
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var board = make([][]int, xDim)
	for i := range board {
		board[i] = make([]int, yDim)
	}

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		text := scanner.Text()
		num := []rune(text)

		for col, n := range num {
			board[col][line], err = strconv.Atoi(string(n))
			check(err)
		}
		line++
	}
	check(scanner.Err())
	return board

}

func increaseBoard(board [][]int) {
	xDim := len(board)
	yDim := len(board[0])

	for col := 0; col < xDim; col++ {
		for line := 0; line < yDim; line++ {
			board[col][line] = board[col][line] + 1

		}
	}
}

func increaseIfValid(board [][]int, x int, y int) {
	xDim := len(board)
	yDim := len(board[0])

	if x >= 0 && y >= 0 && x < xDim && y < yDim {
		board[x][y] = board[x][y] + 1
	}
}

func increaseAdjecent(board [][]int, pos point) {
	x := pos.x
	y := pos.y

	increaseIfValid(board, x+1, y)
	increaseIfValid(board, x+1, y+1)
	increaseIfValid(board, x, y+1)
	increaseIfValid(board, x-1, y+1)
	increaseIfValid(board, x-1, y)
	increaseIfValid(board, x+1, y-1)
	increaseIfValid(board, x, y-1)
	increaseIfValid(board, x-1, y-1)

}

func flash(board [][]int, alreadyFlashed []point) (int, []point) {
	xDim := len(board)
	yDim := len(board[0])

	flashCount := 0

	for col := 0; col < xDim; col++ {
		for line := 0; line < yDim; line++ {
			pos := point{col, line}

			if (board[col][line] > 9) && !contains(alreadyFlashed, pos) {
				flashCount++
				alreadyFlashed = append(alreadyFlashed, pos)
				increaseAdjecent(board, pos)
			}
		}
	}
	return flashCount, alreadyFlashed
}

func round(board [][]int) int {
	increaseBoard(board)
	flashCount := 0
	var alreadyFlashed []point
	for true {
		var newFlashes int
		newFlashes, alreadyFlashed = flash(board, alreadyFlashed)
		flashCount += newFlashes
		if newFlashes == 0 {
			break
		}
	}

	for _, flash := range alreadyFlashed {
		board[flash.x][flash.y] = 0
	}
	fmt.Printf("Flash count: %v\n", flashCount)
	return flashCount
}

func main() {
	fmt.Println("Hello, 世界aa")
	board := parseInputFile("input11")
	printBoard(board)

	totalFlashCount := 0
	for step := 1; step <= 100; step++ {
		fmt.Printf("-------Round %v---\n", step)
		totalFlashCount += round(board)
		fmt.Printf("Total flash count: %v\n", totalFlashCount)
		printBoard(board)
	}

	//part B
	fmt.Print("#################PART B############")
	board = parseInputFile("input11")
	printBoard(board)

	totalFlashCount = 0
	for step := 1; true; step++ {
		fmt.Printf("-------Round %v---\n", step)
		newFlashes := round(board)
		totalFlashCount += newFlashes
		fmt.Printf("Total flash count: %v\n", totalFlashCount)
		printBoard(board)
		if newFlashes == 100 {
			break
		}
	}
}
