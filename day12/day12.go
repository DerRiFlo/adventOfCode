package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	small = iota
	big   = iota
)

type cave struct {
	neighbours []int
	name       string
	size       int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findCave(name string, caves []cave) int {
	for i, v := range caves {
		if v.name == name {
			return i
		}
	}
	return -1
}

func isCapitalLetter(letter rune) bool {
	if letter >= 'A' && letter <= 'Z' {
		return true
	}
	return false
}

func parseInputFile(path string) []cave {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var caves []cave

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		tmp := strings.Split(text, "-")

		startIdx := findCave(tmp[0], caves)
		if startIdx == -1 {
			var start cave
			start.name = tmp[0]
			if isCapitalLetter(rune(start.name[0])) {
				start.size = big
			} else {
				start.size = small
			}
			caves = append(caves, start)
			startIdx = len(caves) - 1
		}

		endIdx := findCave(tmp[1], caves)
		if endIdx == -1 {
			var end cave
			end.name = tmp[1]
			if isCapitalLetter(rune(end.name[0])) {
				end.size = big
			} else {
				end.size = small
			}
			caves = append(caves, end)
			endIdx = len(caves) - 1
		}

		caves[startIdx].neighbours = append(caves[startIdx].neighbours, endIdx)
		caves[endIdx].neighbours = append(caves[endIdx].neighbours, startIdx)

	}
	check(scanner.Err())
	return caves
}

func contains(arr []int, i int) bool {
	for _, c := range arr {
		if c == i {
			return true
		}
	}
	return false
}

func printRoute(caves []cave, route []int) {
	for i, s := range route {
		c := caves[s]
		fmt.Print(c.name)
		if i < len(route)-1 {
			fmt.Print("->")
		}
	}
	fmt.Println("")
}

func findRoutes(caves []cave, currCave int, route []int) [][]int {
	route = append(route, currCave)
	var foundRoutes [][]int

	if caves[currCave].name == "end" {
		foundRoutes = append(foundRoutes, route)
		return foundRoutes
	}

	for i := 0; i < len(caves[currCave].neighbours); i++ {
		c := caves[currCave].neighbours[i]

		if caves[c].name == "start" {
			continue
		}
		if caves[c].size == big {
			foundRoutes = append(foundRoutes, findRoutes(caves, c, route)...)
		} else if !contains(route, c) {
			foundRoutes = append(foundRoutes, findRoutes(caves, c, route)...)
		}
	}
	return foundRoutes
}

func findRoutes2(caves []cave, currCave int, route []int, duplicateUsed bool) [][]int {
	route = append(route, currCave)
	var foundRoutes [][]int

	if caves[currCave].name == "end" {
		foundRoutes = append(foundRoutes, route)
		return foundRoutes
	}

	for i := 0; i < len(caves[currCave].neighbours); i++ {
		c := caves[currCave].neighbours[i]

		if caves[c].name == "start" {
			continue
		}
		if caves[c].size == big {
			foundRoutes = append(foundRoutes, findRoutes2(caves, c, route, duplicateUsed)...)
		} else if !contains(route, c) {
			foundRoutes = append(foundRoutes, findRoutes2(caves, c, route, duplicateUsed)...)
		} else if contains(route, c) && !duplicateUsed {
			foundRoutes = append(foundRoutes, findRoutes2(caves, c, route, true)...)
		}
	}
	return foundRoutes
}

func main() {
	fmt.Println("Hello, 世界aa")
	caves := parseInputFile("input12")

	var startCaveIdx int
	for i, c := range caves {
		if c.name == "start" {
			startCaveIdx = i
		}
	}
	var route []int
	ans := findRoutes(caves, startCaveIdx, route)
	/*for _, route := range ans {
		printRoute(caves, route)
	}*/
	fmt.Println(len(ans))

	ans = findRoutes2(caves, startCaveIdx, route, false)
	/*for _, route := range ans {
		printRoute(caves, route)
	}*/
	fmt.Println(len(ans))

}
