package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInputFile(path string) []int {
	data, err := os.ReadFile(path)
	check(err)

	tmp := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")
	var pos = make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		j, err := strconv.Atoi(tmp[i])
		check(err)
		pos[i] = j
	}
	return pos
}

//Thanks Wikipedia :)
//https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF
func seriesVal(n int) int {
	return n * (n + 1) / 2
}

func getFuel(pos int, positions []int) int {
	fuel := 0
	for i := 0; i < len(positions); i++ {
		fuel += seriesVal(int(math.Abs(float64(positions[i] - pos))))
	}
	return fuel
}

func maxSlice(slc []int) int {
	max := slc[0]
	for i := 0; i < len(slc); i++ {
		if slc[i] > max {
			max = slc[i]
		}
	}
	return max
}

func main() {
	fmt.Println("Hello, 世界aa")
	positions := parseInputFile("input7")

	maxpos := maxSlice(positions)
	minFuel := 999999999
	for pos := 1; pos <= maxpos; pos++ {
		fuel := getFuel(pos, positions)
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	fmt.Println(minFuel)
}
