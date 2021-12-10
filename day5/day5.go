package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type vector struct {
	start point
	end   point
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const dim = 1000

func getPoint(s string) point {
	var p point
	var err error
	fields := strings.Split(s, ",")

	p.x, err = strconv.Atoi(strings.TrimSpace(fields[0]))
	check(err)
	p.y, err = strconv.Atoi(strings.TrimSpace(fields[1]))
	check(err)
	return p
}

func getVector(s string) vector {
	fields := strings.Split(s, "->")
	var v vector
	v.start = getPoint(fields[0])
	v.end = getPoint(fields[1])
	return v
}

func vectorStraight(v vector) bool {
	if v.start.x == v.end.x || v.start.y == v.end.y {
		return true
	}
	return false
}

func parseInputFile(path string) []vector {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var vectors []vector
	for scanner.Scan() {
		vec := getVector(scanner.Text())
		vectors = append(vectors, vec)
	}
	check(scanner.Err())
	return vectors
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	return x / abs(x)
}

func addLine(v vector, board *[dim][dim]int) {
	if v.end.x != v.start.x {
		s := sgn(v.end.x - v.start.x)
		dist := abs(v.end.x - v.start.x)
		for i := 0; i <= dist; i++ {
			x := v.start.x + s*i
			board[x][v.start.y] += 1
		}
	}

	if v.end.y != v.start.y {
		s := sgn(v.end.y - v.start.y)
		dist := abs(v.end.y - v.start.y)
		for i := 0; i <= dist; i++ {
			y := v.start.y + s*i
			board[v.start.x][y] += 1
		}
	}
}

func printBoard(board *[dim][dim]int) {
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			fmt.Print(board[x][y])
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func countBoard(board *[dim][dim]int) int {
	score := 0
	for y := 0; y < dim; y++ {
		for x := 0; x < dim; x++ {
			if board[x][y] > 1 {
				score++
			}
		}
	}
	return score
}

func main() {
	fmt.Println("Hello, 世界aa")
	vectors := parseInputFile("input5")
	board := [dim][dim]int{}

	for index, element := range vectors {
		if vectorStraight(element) {
			fmt.Println(element)
			fmt.Println(index)
			addLine(element, &board)
		}
	}
	score := countBoard(&board)
	//printBoard(&board)
	fmt.Println("score:")
	fmt.Println(score)

}
