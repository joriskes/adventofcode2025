package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const RUN_EXAMPLE = false

func main() {
	fmt.Println("====== Day 6 ======")
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

	type homeworkAssignment struct {
		values []int
		symbol string
	}

	removeDupeSpacesRE := regexp.MustCompile(`\s+`)

	homework := make([]homeworkAssignment, 0)
	l := 0
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}

		t = strings.TrimSpace(removeDupeSpacesRE.ReplaceAllString(t, " "))
		spl := strings.Split(t, " ")
		for i, v := range spl {
			n, err := strconv.Atoi(v)
			if l == 0 {
				homework = append(homework, homeworkAssignment{values: []int{}, symbol: "?"})
			}
			if err == nil {
				homework[i].values = append(homework[i].values, n)
			} else {
				homework[i].symbol = v
			}
		}
		l++;
	}

	for _, v := range homework {
		answer := v.values[0]
		for _, n := range v.values[1:] {
			if v.symbol == "*" {
				answer *= n
			}
			if v.symbol == "+" {
				answer += n
			}
		}
		part1 += answer
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 6 time:", time.Since(start))
}
