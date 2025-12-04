package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("day4/task1/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	m := make([][]bool, 0)

	for scanner.Scan() {
		text := scanner.Text()

		w := len(text)
		line := make([]bool, w)

		for j := 0; j < w; j++ {
			line[j] = text[j:j+1] == "@"
		}

		m = append(m, line)
	}

	count := 0

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if !m[i][j] {
				//fmt.Print(".")
				continue
			}

			nc := 0

			// left
			if j > 0 && m[i][j-1] {
				nc++
			}

			// top-left
			if j > 0 && i > 0 && m[i-1][j-1] {
				nc++
			}

			// top
			if i > 0 && m[i-1][j] {
				nc++
			}

			// top-right
			if i > 0 && j < len(m[i])-1 && m[i-1][j+1] {
				nc++
			}

			// right
			if j < len(m[i])-1 && m[i][j+1] {
				nc++
			}

			// right-bottom
			if i < len(m)-1 && j < len(m[i])-1 && m[i+1][j+1] {
				nc++
			}

			// down
			if i < len(m)-1 && m[i+1][j] {
				nc++
			}

			// left-bottom
			if i < len(m)-1 && j > 0 && m[i+1][j-1] {
				nc++
			}

			if nc < 4 {
				count++
				//fmt.Print("x")
			} else {
				//fmt.Print("@")
			}
		}
		//fmt.Println()
	}

	fmt.Println(count)
}
