package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const RUN_EXAMPLE = false

func main() {
	fmt.Println("====== Day 1 ======")
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

	dial := 50;

	for scanner.Scan() {
		t := scanner.Text()
		side := t[:1]
		dist, err := strconv.Atoi(t[1:])
		if err != nil {
			panic(err)
		}

		for i:=0; i<dist; i++ {
			if side == "L" {
				dial--
			} else {
				dial++
			}
			if dial < 0 {
				dial = dial + 100
			}
			if dial > 99 {
				dial = dial - 100
			}
			if dial == 0 {
				part2++
			}
		}
		if dial == 0 {
			part1++
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 1 time:", time.Since(start))
}
