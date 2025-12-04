package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func main() {
	f, _ := os.Open("day2/task2/input.txt")
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

			start, _ := strconv.Atoi(s[0])
			end, _ := strconv.Atoi(s[1])

			arr = append(arr, Range{start, end})
		}
	}

	sum := 0

	for _, r := range arr {
		sum += calculateSumForRange(r)
	}

	fmt.Println("sum:", sum)
}

func calculateSumForRange(r Range) int {
	sum := 0

	fmt.Print(r, ": ")
	for i := r.Start; i <= r.End; i++ {
		if isInvalidId(i) {
			sum += i
			fmt.Print(i, ",")
		}
	}

	fmt.Println()

	return sum
}

func isInvalidId(candidate int) bool {
	cs := strconv.Itoa(candidate)

	for i := 1; i <= len(cs)/2; i++ {
		x := cs[:i]
		s := x

		for len(s) < len(cs) {
			s += x
		}

		if s != cs {
			continue
		}

		return true
	}

	return false
}
