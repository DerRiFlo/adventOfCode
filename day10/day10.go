package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var openRunes [4]rune = [4]rune{'(', '{', '[', '<'}
var closeRunes [4]rune = [4]rune{')', '}', ']', '>'}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find(list []rune, r rune) int {
	for idx, v := range list {
		if r == v {
			return idx
		}
	}
	return -1
}

func parseLine(lineS string) int {
	closePoints := [4]int{3, 1197, 57, 25137}

	var chunkStack []int

	line := []rune(lineS)
	for _, r := range line {
		openIdx := find(openRunes[:], r)
		if openIdx > -1 {
			chunkStack = append(chunkStack, openIdx)
			continue
		}
		closeIdx := find(closeRunes[:], r)
		if closeIdx == chunkStack[len(chunkStack)-1] {
			chunkStack = chunkStack[:len(chunkStack)-1]
		} else {
			return closePoints[closeIdx]
		}
	}
	return 0
}

func completeLine(lineS string) []rune {

	var chunkStack []int
	var completionRunes []rune

	line := []rune(lineS)
	for _, r := range line {
		openIdx := find(openRunes[:], r)
		if openIdx > -1 {
			chunkStack = append(chunkStack, openIdx)
			continue
		} else {
			chunkStack = chunkStack[:len(chunkStack)-1]
		}

	}

	for i := len(chunkStack) - 1; i >= 0; i-- {
		c := chunkStack[i]
		completionRunes = append(completionRunes, closeRunes[c])

	}

	return completionRunes
}

func calcCompeltionScore(completionRunes []rune) int {
	completionPoints := [4]int{1, 3, 2, 4}
	score := 0

	for _, r := range completionRunes {
		idx := find(closeRunes[:], r)
		s := completionPoints[idx]
		score = score*5 + s
	}
	return score
}

func parseInputFile(path string) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNum := 0
	pointSum := 0
	var completionScores []int

	for scanner.Scan() {
		text := scanner.Text()
		points := parseLine(text)
		if points > 0 {
			pointSum += points
			fmt.Printf("Corrupt Line %v: %v points\n", lineNum, points)
		} else {
			completionString := completeLine(text)
			pts := calcCompeltionScore(completionString)
			completionScores = append(completionScores, pts)
			fmt.Printf("Incomplete Line %v: %v points, %v\n", lineNum, pts, string(completionString))
		}

		lineNum++
	}
	fmt.Printf("Total points: %v\n", pointSum)

	sort.Ints(completionScores)
	idx := int(math.Floor(float64(len(completionScores)) / 2.0))
	medianScore := completionScores[idx]
	fmt.Printf("Median completion score: %v", medianScore)
}
func main() {
	fmt.Println("Hello, 世界aa")
	parseInputFile("input10")

}
