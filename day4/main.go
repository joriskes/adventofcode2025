package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const RUN_EXAMPLE = false

func getGrid(grid [][]int, x int, y int) int {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0
	}
	return grid[y][x]
}

func sumNeighbours(grid [][]int, x int, y int) int {
	return getGrid(grid, x-1, y-1) + getGrid(grid, x, y-1) + getGrid(grid, x+1, y-1) +
		getGrid(grid, x-1, y) + getGrid(grid, x+1, y) + getGrid(grid, x-1, y+1) +
		getGrid(grid, x, y+1) + getGrid(grid, x+1, y+1)
}

func main() {
	fmt.Println("====== Day 4 ======")
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

	grid := make([][]int, 0)
	for scanner.Scan() {
		t := scanner.Text()
		line := strings.Split(t, "")
		rolls := make([]int, len(line))
		for i, v := range line {
			if v == "@" {
				rolls[i] = 1
			} else {
				rolls[i] = 0
			}
		}

		grid = append(grid, rolls)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 && sumNeighbours(grid, x, y) < 4 {
				part1++
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 4 time:", time.Since(start))
}
