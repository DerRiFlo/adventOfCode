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
	var fish = make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		j, err := strconv.Atoi(tmp[i])
		check(err)
		fish[i] = j
	}
	return fish
}

func iterate(fish []int, start int, end int, c chan []int) {
	var newFish = []int{}

	for i := start; i < end; i++ {
		if fish[i] == 0 {
			fish[i] = 6
			newFish = append(newFish, 8)
		} else {
			fish[i] = fish[i] - 1
		}

	}
	c <- newFish
}

func main() {
	fmt.Println("Hello, 世界aa")
	fish := parseInputFile("input6")

	part1 := int(math.Floor(float64(len(fish) / 4.0)))
	part2 := int(math.Floor(float64(len(fish) / 2.0)))
	part3 := int(math.Floor(float64(len(fish) * 3.0 / 4.0)))

	for i := 0; i < 256; i++ {
		fmt.Print(i)
		fmt.Print("->")
		fmt.Println(len(fish))
		c := make(chan []int)

		go iterate(fish, 0, part1, c)
		go iterate(fish, part1, part2, c)
		go iterate(fish, part2, part3, c)
		go iterate(fish, part3, len(fish), c)

		append1, append2, append3, append4 := <-c, <-c, <-c, <-c
		fish = append(fish, append1...)
		fish = append(fish, append2...)
		fish = append(fish, append3...)
		fish = append(fish, append4...)
	}
	fmt.Print("Total fish: ")
	fmt.Println(len(fish))
}
