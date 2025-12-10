package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	f, err := os.Open("day6/task2/input.txt")
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
	numbers := make([]string, 0)
	problems := make([]Problem, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line[:1] == string(Multiplication) || line[:1] == string(Addition) {
			var p *Problem
			i := 0
			for i = range line {
				ch := line[i : i+1]
				if ch == string(Multiplication) || ch == string(Addition) {
					if p != nil {
						problems = append(problems, *p)
					}

					p = &Problem{make([]int, 0), Sign(ch)}
				}

				s := strings.TrimSpace(numbers[i])
				if s != "" {
					num, _ := strconv.Atoi(s)
					p.Numbers = append(p.Numbers, num)
				}
			}

			for j := i + 1; j < len(numbers); j++ {
				s := strings.TrimSpace(numbers[j])
				if s != "" {
					num, _ := strconv.Atoi(s)
					p.Numbers = append(p.Numbers, num)
				}
			}

			problems = append(problems, *p)

			break
		}

		for i := range line {
			ch := line[i : i+1]

			if len(numbers) <= i {
				numbers = append(numbers, ch)
			} else {
				numbers[i] += ch
			}
		}

	}

	return problems
}
