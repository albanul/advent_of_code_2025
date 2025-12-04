package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction string

const (
	LEFT  Direction = "L"
	RIGHT Direction = "R"
)

func main() {
	f, _ := os.Open("day1/task1/input.txt")
	defer f.Close()

	r := bufio.NewScanner(f)

	cmds := make([]Command, 0)

	for r.Scan() {
		line := r.Text()

		d := Direction(line[0])
		vStr := line[1:]

		v, _ := strconv.Atoi(vStr)

		cmd := Command{Direction: d, Value: v}
		cmds = append(cmds, cmd)
	}

	current := 50
	count := 0

	for _, cmd := range cmds {
		if cmd.Direction == LEFT {
			current -= cmd.Value
		} else if cmd.Direction == RIGHT {
			current += cmd.Value
		}

		current %= 100

		if current < 0 {
			current = 100 + current
		}

		if current == 0 {
			count++
		}
	}

	fmt.Println(count)
}

type Command struct {
	Direction Direction
	Value     int
}
