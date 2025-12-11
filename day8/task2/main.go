package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

type Point struct {
	X, Y, Z int
}

type Edge struct {
	A, B     *Point
	Distance float64
}

func main() {
	f, err := os.Open("day8/task2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := ParsePoints(f)

	edges, circuits := getEdgesAndCircuits(points)
	sorted := getSortedEdges(edges)

	circuitSets := map[*[]*Point]bool{}
	for _, c := range circuits {
		circuitSets[c] = true
	}

	result := 0

	for i := 0; i < len(sorted); i++ {
		edge := sorted[i]

		c1 := circuits[*edge.A]
		c2 := circuits[*edge.B]

		c3 := make([]*Point, 0)

		c3 = append(c3, *c1...)

		for _, p := range *c2 {
			if !slices.Contains(*c1, p) {
				c3 = append(c3, p)
			}
		}

		for _, p := range c3 {
			c := circuits[*p]
			delete(circuitSets, c)

			circuits[*p] = &c3
		}

		delete(circuitSets, c1)
		delete(circuitSets, c2)
		circuitSets[&c3] = true

		if len(c3) == len(points) {
			result = edge.A.X * edge.B.X
			break
		}
	}

	fmt.Println(result)
}

func getEdgesAndCircuits(points []Point) ([]Edge, map[Point]*[]*Point) {
	edges := make([]Edge, 0)
	circuit := make(map[Point]*[]*Point)

	for i := 0; i < len(points); i++ {
		p1 := &points[i]

		circuit[*p1] = &[]*Point{p1}

		for j := i + 1; j < len(points); j++ {
			p2 := &points[j]

			d := Edge{p1, p2, calculateDistance(*p1, *p2)}
			edges = append(edges, d)
		}
	}
	return edges, circuit
}

func ParsePoints(f *os.File) []Point {
	points := make([]Point, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		p := Point{}
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)

		points = append(points, p)
	}
	return points
}

func calculateDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2) + math.Pow(float64(p1.Z-p2.Z), 2))
}

func getSortedEdges(edges []Edge) []Edge {
	seq := func(yield func(Edge) bool) {
		for _, d := range edges {
			if !yield(d) {
				break
			}
		}
	}

	sorted := slices.SortedFunc(seq, func(a, b Edge) int {
		return cmp.Compare(a.Distance, b.Distance)
	})
	return sorted
}
