package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const RUN_EXAMPLE = false

type pos struct {
	x, y, z   int
}
type circuit struct {
	indexes[] int
}

type edge struct {
	from, to int
	distance float64
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func pow2(x int) int {
	return x * x
}

func dist(pos1, pos2 pos) float64 {
	return math.Sqrt(float64(pow2(pos1.x-pos2.x) + pow2(pos1.y-pos2.y) + pow2(pos1.z-pos2.z)))
}

func findCircuit(circuits []circuit, index int) int {
	for ci, c := range circuits {
		for _, i := range c.indexes {
			if i == index {
				return ci
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("====== Day 8 ======")
	start := time.Now()
	fileName := "input.txt"
	if RUN_EXAMPLE {
		fileName = "example.txt"
		fmt.Println("Running on example data")
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	points := make([]pos, 0)
	edges := make([]edge, 0)
	circuits := make([]circuit, 0)

	for scanner.Scan() {
		t := scanner.Text()
		strs := strings.Split(t, ",")
		points = append(points, pos{x: atoi(strs[0]), y: atoi(strs[1]), z: atoi(strs[2])})
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, edge{from: i, to: j, distance: dist(points[i], points[j])})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	todo := 1000
	if RUN_EXAMPLE {
		todo = 10
	}

	for i := 0; i < todo; i++ {
		e := edges[i]
		//fmt.Println("Processing edge", points[e.from], "->", points[e.to])
		fromCI := findCircuit(circuits, e.from)
		toCI := findCircuit(circuits, e.to)

		if fromCI == -1 {
			if toCI == -1 {
				circuits = append(circuits, circuit{indexes: []int{e.from, e.to}})
			} else {
				circuits[toCI].indexes = append(circuits[toCI].indexes, e.from)
			}
		} else {
			if toCI == -1 {
				circuits[fromCI].indexes = append(circuits[fromCI].indexes, e.to)
			} else {
				//fmt.Println("skipping edge, already in circuit")
				if fromCI != toCI {
					// Merge circuits
					fci := circuits[fromCI].indexes
					circuits[toCI].indexes = append(circuits[toCI].indexes, fci...)
					circuits = append(circuits[:fromCI], circuits[fromCI+1:]...)
				}
			}
		}
		//fmt.Println(circuits)
	}


	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].indexes) > len(circuits[j].indexes)
	})

	part1 = 1
	for i:=0; i<3; i++ {
		part1 *= len(circuits[i].indexes)
	}


	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 8 time:", time.Since(start))
}
