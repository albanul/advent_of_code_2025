package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Sign string

const (
	Multiplication Sign = "*"
	Addition       Sign = "+"
)

type Problem struct {
	Numbers []int
	Sign    Sign
}

func main() {
	f, err := os.Open("day6/task1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	problems := getProblems(scanner)

	//fmt.Println(problems)

	sum := 0

	for i := 0; i < len(problems); i++ {
		problem := problems[i]
		numbers := problem.Numbers

		var s int

		if problem.Sign == Multiplication {
			s = numbers[0]
			for j := 1; j < len(numbers); j++ {
				s *= numbers[j]
			}
		} else if problem.Sign == Addition {
			s = 0
			for j := 0; j < len(numbers); j++ {
				s += numbers[j]
			}
		}

		sum += s
	}

	fmt.Println(sum)
}

func getProblems(scanner *bufio.Scanner) []Problem {
	problems := make([]Problem, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		split := strings.Split(line, " ")

		i := 0

		if unicode.IsNumber(rune(split[0][0])) {
			for _, s := range split {
				v, err := strconv.Atoi(s)
				if err != nil {
					continue
				}

				if len(problems) <= i {
					problems = append(problems, Problem{Numbers: make([]int, 0)})
				}

				problems[i].Numbers = append(problems[i].Numbers, v)
				i++
			}
		} else {
			for _, s := range split {
				if s == "" {
					continue
				}

				problems[i].Sign = Sign(s)
				i++
			}
		}
	}

	return problems
}
