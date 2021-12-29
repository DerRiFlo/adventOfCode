package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type fold struct {
	val int
	ax  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNumber(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func parseInputFile(path string) ([]point, []fold) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var points []point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 || !isNumber(rune(text[0])) {
			break
		}

		tmp := strings.Split(text, ",")
		var newPoint point
		newPoint.x, err = strconv.Atoi(tmp[0])
		check(err)
		newPoint.y, err = strconv.Atoi(tmp[1])
		check(err)

		points = append(points, newPoint)
	}

	var folds []fold
	for scanner.Scan() {
		text := scanner.Text()
		tmp := strings.Split(text, " ")
		tmp = strings.Split(tmp[2], "=")

		var newFold fold
		newFold.ax = tmp[0]
		newFold.val, err = strconv.Atoi(tmp[1])
		check(err)
		folds = append(folds, newFold)
	}
	check(scanner.Err())
	return points, folds
}

func execFold(points *[]point, f fold) {
	for i := 0; i < len(*points); i++ {
		if f.ax == "x" && (*points)[i].x > f.val {
			(*points)[i].x = f.val - ((*points)[i].x - f.val)
		}
		if f.ax == "y" && (*points)[i].y > f.val {
			(*points)[i].y = f.val - ((*points)[i].y - f.val)
		}
	}
}

func unique(points []point) []point {
	keys := make(map[point]bool)
	list := []point{}
	for _, entry := range points {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func getDims(points []point) (int, int) {
	xMax, yMax := 0, 0

	for _, p := range points {
		if p.x > xMax {
			xMax = p.x
		}
		if p.y > yMax {
			yMax = p.y
		}
	}
	xMax++
	yMax++

	return xMax, yMax
}

func printPaper(points []point) {

	xMax, yMax := getDims(points)
	var board = make([][]bool, xMax)
	for i := range board {
		board[i] = make([]bool, yMax)
	}

	for _, p := range points {
		board[p.x][p.y] = true
	}

	fmt.Println("----------")
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if board[x][y] {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	fmt.Println("Hello, 世界aa")
	points, folds := parseInputFile("input13")

	execFold(&points, folds[0])
	uPoints := unique(points)
	fmt.Printf("Unique Points: %v\n", len(uPoints))

	for i := 1; i < len(folds); i++ {
		execFold(&points, folds[i])
	}
	printPaper(points)
}
