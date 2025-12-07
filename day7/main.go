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

func main() {
	fmt.Println("====== Day 7 ======")
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

	beamList := make(map[int]bool)

	for scanner.Scan() {
		t := scanner.Text()
		chars := strings.Split(t, "")
		for x, c := range chars {
			switch c {
			case "S":
				beamList[x] = true
			case "^":
				if beamList[x] {
					part1++
					beamList[x-1] = true
					beamList[x+1] = true
					beamList[x] = false
				}
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 7 time:", time.Since(start))
}
