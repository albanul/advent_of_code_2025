package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Bank struct {
	Batteries []Battery
}

type Battery struct {
	Value, Position int
}

func main() {
	f, _ := os.Open("day3/task1/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	banks := make([]Bank, 0)

	for scanner.Scan() {

		line := scanner.Text()

		arr := make([]Battery, 0)

		for i := 0; i < len(line); i++ {
			v := line[i : i+1]
			vInt, _ := strconv.Atoi(v)
			arr = append(arr, Battery{vInt, i})
		}

		banks = append(banks, Bank{arr})
	}

	sum := 0

	for i, b := range banks {
		j := calculateJoltage(b)
		fmt.Println(i+1, ":", j)
		sum += j
	}

	fmt.Println(sum)
}

func calculateJoltage(b Bank) int {
	seq := func(yield func(Battery) bool) {
		for _, x := range b.Batteries {
			if !yield(x) {
				return
			}
		}
	}
	sorted := slices.SortedFunc(seq, func(a, b Battery) int {
		c := -cmp.Compare(a.Value, b.Value)

		if c == 0 {
			c = cmp.Compare(a.Position, b.Position)
		}

		return c
	})

	first, second := sorted[0], sorted[1]

	// max at the end
	if first.Position == len(b.Batteries)-1 {
		return second.Value*10 + first.Value
	}

	// find the next max battery staying after the first max battery
	for i := 1; ; i++ {
		if sorted[i].Position > first.Position {
			second = sorted[i]
			break
		}
	}

	return first.Value*10 + second.Value
}
