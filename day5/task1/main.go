package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

type Node struct {
	Left, Right *Node
	Range       Range
}

func main() {
	f, err := os.Open("day5/task1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ranges, ids := getRangesAndIDs(f)
	ranges = compactRanges(ranges)

	//log.Println(ranges)

	root := createTree(ranges)

	count := 0

	for _, id := range ids {
		if isInTree(root, id) {
			count++
		}
	}

	fmt.Println(count)
}

func getRangesAndIDs(f *os.File) ([]Range, []int) {
	scanner := bufio.NewScanner(f)

	ranges := make([]Range, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		s := strings.Split(line, "-")

		start, _ := strconv.Atoi(s[0])
		end, _ := strconv.Atoi(s[1])

		ranges = append(ranges, Range{start, end})
	}

	ids := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		v, _ := strconv.Atoi(line)

		ids = append(ids, v)
	}
	return ranges, ids
}

func compactRanges(ranges []Range) []Range {
	seq := func(yield func(Range) bool) {
		for _, r := range ranges {
			if !yield(r) {
				break
			}
		}
	}

	sorted := slices.SortedFunc(seq, func(a, b Range) int {
		return cmp.Compare(a.Start, b.Start)
	})

	result := make([]Range, 0)

	r := Range{sorted[0].Start, sorted[0].End}

	for i := 0; i < len(sorted)-1; i++ {
		nr := sorted[i+1]

		if nr.Start <= r.End {
			if nr.End <= r.End {
				continue
			}

			r.End = nr.End
		} else {
			result = append(result, r)
			r = nr
		}
	}

	result = append(result, r)

	return result
}

func createTree(ranges []Range) *Node {
	if len(ranges) == 0 {
		return nil
	}

	if len(ranges) == 1 {
		return &Node{nil, nil, ranges[0]}
	}

	full := Range{ranges[0].Start, ranges[len(ranges)-1].End}

	i := len(ranges) / 2

	lRanges := ranges[:i]
	rRanges := ranges[i:]

	left := createTree(lRanges)
	right := createTree(rRanges)

	return &Node{left, right, full}
}

func isInTree(node *Node, id int) bool {
	if node == nil {
		return false
	}

	if id < node.Range.Start || id > node.Range.End {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return node.Range.Start <= id && id <= node.Range.End
	}

	return isInTree(node.Left, id) || isInTree(node.Right, id)
}
