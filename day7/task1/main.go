package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day7/task1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()

	beams := map[int]bool{strings.Index(line, "S"): true}

	splitCount := 0

	for scanner.Scan() {
		line = scanner.Text()

		nBeams := map[int]bool{}

		for i := range line {
			ch := line[i : i+1]

			if ch == "^" && beams[i] {
				delete(beams, i)

				if i > 0 && !nBeams[i-1] {
					nBeams[i-1] = true
				}

				if i < len(line)-1 && !nBeams[i+1] {
					nBeams[i+1] = true
				}

				splitCount++
			}
		}

		for k, v := range beams {
			nBeams[k] = v
		}

		beams = nBeams
	}

	fmt.Println(splitCount)
}
