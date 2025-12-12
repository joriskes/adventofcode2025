package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const RUN_EXAMPLE = false

type pos struct {
	x, y int
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func surfaceArea(p1, p2 pos) int {
	return int(math.Abs(float64(p1.x - p2.x) + 1) * (math.Abs(float64(p1.y - p2.y)) + 1))
}

func main() {
	fmt.Println("====== Day 9 ======")
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
	for scanner.Scan() {
		t := scanner.Text()
		strs := strings.Split(t, ",")
		points = append(points, pos{x: atoi(strs[0]), y: atoi(strs[1])})
	}

	highest := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			s := surfaceArea(points[i], points[j])
			if s > highest {
				highest = s
			}
		}
	}

	part1 = highest

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 9 time:", time.Since(start))
}
