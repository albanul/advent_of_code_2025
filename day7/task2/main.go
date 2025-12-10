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

	beams := map[int]int{strings.Index(line, "S"): 1}

	for scanner.Scan() {
		line = scanner.Text()

		nBeams := map[int]int{}
		splitter := map[int]bool{}

		for i := range line {
			ch := line[i : i+1]

			v, ok := beams[i]
			if ch == "^" && ok {
				splitter[i] = true
				delete(beams, i)

				lV := nBeams[i-1]

				if i > 0 {
					nBeams[i-1] = lV + v
				}

				rV := nBeams[i+1]

				if i < len(line)-1 {
					nBeams[i+1] = rV + v
				}
			}
		}

		for k, v := range beams {
			if _, ok := nBeams[k]; ok {
				nBeams[k] += v
			} else {
				nBeams[k] = v
			}
		}

		/*for i := 0; i < len(line); i++ {
			v, ok := nBeams[i]

			if !ok {
				if _, ok := splitter[i]; ok {
					fmt.Print("^")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(v)
			}
		}*/
		//fmt.Println()

		beams = nBeams
	}

	sum := 0

	for _, v := range beams {
		sum += v
	}

	fmt.Println(sum)
}
