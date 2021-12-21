package main

import (
	"bufio"
	"fmt"
	"math"
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

func getDims(path string) (int, int) {
	/*file, err := os.Open(path)
	defer file.Close()
	check(err)
	tmp, err := file.Stat()
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	//x := len(text)
	//y := int(tmp.Size()) / x*/
	return 100, 100
}

func parseInputFile(path string) [][]int {
	xDim, yDim := getDims(path)
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

func isMin(xPos int, yPos int, board [][]int) bool {
	xDim := len(board)
	yDim := len(board[0])
	val := board[xPos][yPos]

	if (xPos+1 < xDim) && (board[xPos+1][yPos] <= val) {
		return false
	}
	if (xPos > 0) && (board[xPos-1][yPos] <= val) {
		return false
	}
	if (yPos+1 < yDim) && (board[xPos][yPos+1] <= val) {
		return false
	}
	if (yPos > 0) && (board[xPos][yPos-1] <= val) {
		return false
	}
	return true
}

func getMinSum(board [][]int) int {
	xDim := len(board)
	yDim := len(board[0])
	sum := 0
	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			if isMin(x, y, board) {
				sum += board[x][y] + 1
				fmt.Printf("[%v,%v] %v\n", x, y, board[x][y])
			}
		}
	}
	return sum
}

func getNeighbours(board [][]int, pos point) []point {
	xDim := len(board)
	yDim := len(board[0])
	var neighbours []point

	if pos.x+1 < xDim {
		n := point{pos.x + 1, pos.y}
		neighbours = append(neighbours, n)
	}
	if pos.x > 0 {
		n := point{pos.x - 1, pos.y}
		neighbours = append(neighbours, n)
	}
	if pos.y+1 < yDim {
		n := point{pos.x, pos.y + 1}
		neighbours = append(neighbours, n)
	}
	if pos.y > 0 {
		n := point{pos.x, pos.y - 1}
		neighbours = append(neighbours, n)
	}

	return neighbours
}

func contains(pointList []point, p point) bool {
	for _, v := range pointList {
		if p == v {
			return true
		}
	}

	return false
}

func calcPool(board [][]int, pos point, processedPoints []point) []point {
	processedPoints = append(processedPoints, pos)
	neighbours := getNeighbours(board, pos)
	for _, n := range neighbours {
		if !contains(processedPoints, n) && board[n.x][n.y] < 9 {
			processedPoints = calcPool(board, n, processedPoints)
		}
	}
	return processedPoints
}

func getMin(arr []int) (int, int) {
	min := math.MaxInt
	minPos := -1

	for idx, i := range arr {
		if i < min {
			min = i
			minPos = idx
		}
	}
	return min, minPos
}

func calcPools(board [][]int) int {
	xDim := len(board)
	yDim := len(board[0])
	threeLargest := [3]int{0, 0, 0}

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			if isMin(x, y, board) {
				p := point{x, y}
				var pool []point
				pool = calcPool(board, p, pool)
				fmt.Printf("[%v,%v] %v\n", x, y, pool)
				min, minPos := getMin(threeLargest[:])
				if len(pool) > min {
					threeLargest[minPos] = len(pool)
				}
			}
		}
	}
	return threeLargest[0] * threeLargest[1] * threeLargest[2]
}

func main() {
	fmt.Println("Hello, 世界aa")
	board := parseInputFile("input9")

	sum := getMinSum(board)
	fmt.Println(sum)

	sum = calcPools(board)
	fmt.Println(sum)

}
