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


var memos = make(map[int]int)
func bfsLines(x int, field []string, lineIndex int) int {
	if lineIndex >= len(field) {
		return 1
	}
	if memos[lineIndex * 1000 + x] != 0 {
		return memos[lineIndex * 1000 + x]
	}
	res:=0
	nextLine := field[lineIndex]
	if nextLine[x] == '^' {
		res += bfsLines(x-1, field, lineIndex+1) + bfsLines(x+1, field, lineIndex+1)
	} else {
		res += bfsLines(x, field, lineIndex+1)
	}
	memos[lineIndex * 1000 + x] = res
	return res
}

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
	field := make([]string, 0)
	for scanner.Scan() {
		field = append(field, scanner.Text())
	}

	for lineIndex, line := range field {
		chars := strings.Split(line, "")
		for x, c := range chars {
			switch c {
			case "S":
				beamList[x] = true
				part2 = bfsLines(x, field, lineIndex + 1)
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
