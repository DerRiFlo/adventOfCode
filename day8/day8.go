package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type display struct {
	connections []string
	data        []string
	codes       [10]string
}

func alphasort(input string, depth int) string {
	str := []rune(input)
	for x := range str {
		y := x + 1
		for y = range str {
			if str[x] < str[y] {
				str[x], str[y] = str[y], str[x]
			}
		}
	}
	return string(str)
}

func parseInputFile(path string) []display {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var displays []display
	for scanner.Scan() {
		text := scanner.Text()
		tmp := strings.Split(text, "|")

		var disp display
		disp.connections = strings.Split(strings.TrimSpace(tmp[0]), " ")
		for i := range disp.connections {
			disp.connections[i] = alphasort(disp.connections[i], 1)
		}

		disp.data = strings.Split(strings.TrimSpace(tmp[1]), " ")
		for i := range disp.data {
			disp.data[i] = alphasort(disp.data[i], 1)
		}

		displays = append(displays, disp)
	}
	check(scanner.Err())
	return displays

}

func stringAinB(a, b string) bool {
	retval := true
	rA := []rune(a)
	for _, r := range rA {
		retval = retval && strings.ContainsRune(b, r)
	}
	return retval
}

func numRunesEqual(a, b string) int {
	retval := 0
	rA := []rune(a)
	for _, r := range rA {
		if strings.ContainsRune(b, r) {
			retval++
		}
	}
	return retval
}

func decode(disp *display) {
	for _, code := range disp.connections {
		if len(code) == 2 {
			disp.codes[1] = code
		} else if len(code) == 3 {
			disp.codes[7] = code
		} else if len(code) == 4 {
			disp.codes[4] = code
		} else if len(code) == 7 {
			disp.codes[8] = code
		}
	}

	for _, code := range disp.connections {
		// 0, 6, 9
		if len(code) == 6 {
			//all segments of 4 must be on in 9
			if stringAinB(disp.codes[4], code) {
				disp.codes[9] = code
			} else if stringAinB(disp.codes[1], code) {
				disp.codes[0] = code
			} else if numRunesEqual(code, disp.codes[4]) == 3 {
				disp.codes[6] = code // 4 and 6 have 3 equal segments
			}
		}
	}

	for _, code := range disp.connections {
		//2,3,5,6
		if len(code) == 5 {
			//all segments of 1 in 3
			if stringAinB(disp.codes[1], code) {
				disp.codes[3] = code
			} else if stringAinB(code, disp.codes[9]) {
				disp.codes[5] = code //all segnments of 5 in 9
			} else if numRunesEqual(code, disp.codes[4]) == 2 {
				disp.codes[2] = code // 4 and 2 have 2 equal segments
			}
		}
	}
	fmt.Println(disp.codes)
}

func findCode(codes []string, s string) int {
	for i, c := range codes {
		if s == c {
			return i
		}
	}
	return -1
}

func decodeOutput(disp display) int {
	var output []int
	for _, s := range disp.data {
		output = append(output, findCode(disp.codes[:], s))
	}

	retval := 0
	maxPow := len(output) - 1
	for i, val := range output {
		retval += int(math.Pow10(maxPow-i)) * val
	}
	return retval
}

func main() {
	fmt.Println("Hello, 世界aa")
	displays := parseInputFile("input8")

	fmt.Println("---------A----------")
	counter := 0
	for _, disp := range displays {
		for _, data := range disp.data {
			l := len(data)
			if (l == 2) || (l == 3) || (l == 4) || (l == 7) {
				counter++

				fmt.Println(data)
			}
		}
	}
	fmt.Print(displays)

	fmt.Println("---------B----------")
	out := 0
	for _, disp := range displays {
		decode(&disp)
		out += decodeOutput(disp)
	}

	fmt.Println(out)
}
