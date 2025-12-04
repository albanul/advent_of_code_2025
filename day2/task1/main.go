package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day2/task1/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	arr := make([]Range, 0)

	for scanner.Scan() {
		line := scanner.Text()

		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			if len(r) == 0 {
				continue
			}

			s := strings.Split(r, "-")
			arr = append(arr, Range{s[0], s[1]})
		}
	}

	sum := 0

	for _, r := range arr {
		l := len(r.Start)

		var x string

		if l%2 == 0 {
			x = r.Start[:l/2]
		} else {
			x = "1"
			for i := 0; i < l/2; i++ {
				x += "0"
			}
		}

		fmt.Print(r, x, ":")

		for {
			candidate := x + x

			candidateInt, _ := strconv.Atoi(candidate)
			startInt, _ := strconv.Atoi(r.Start)
			endInt, _ := strconv.Atoi(r.End)

			if candidateInt > endInt {
				break
			}

			xi, _ := strconv.Atoi(x)
			x = strconv.Itoa(xi + 1)

			if candidateInt < startInt {
				continue
			}

			fmt.Print(candidate, ",")

			sum += candidateInt

		}

		fmt.Println()
	}

	fmt.Println("sum:", sum)
}

type Range struct {
	Start, End string
}
