package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const DIGIT_LENGTH = 12

type Bank struct {
	Batteries []Battery
}

type Battery struct {
	Value, Position int
}

func main() {
	f, _ := os.Open("day3/task2/input.txt")
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
	jArr := make([]int, 0)
	jArr = append(jArr, b.Batteries[0].Value)

	last := jArr[0]

	for i := 1; i < len(b.Batteries); i++ {
		val := b.Batteries[i].Value

		if val > last {
			j := len(jArr) - 1

			digitsLeft := len(b.Batteries) - i

			for (j+digitsLeft >= DIGIT_LENGTH) && j >= 0 && jArr[j] < val {
				j--
			}

			jArr = jArr[:j+1]
			if len(jArr) < DIGIT_LENGTH {
				jArr = append(jArr, val)
			}
			continue
		}

		if len(jArr) == DIGIT_LENGTH {
			continue
		}

		jArr = append(jArr, b.Batteries[i].Value)

		last = val
	}

	joltage := jArr[0]

	for i := 1; i < len(jArr); i++ {
		joltage = joltage*10 + jArr[i]
	}

	return joltage
}
