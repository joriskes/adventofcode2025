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

type homeworkAssignment struct {
	values []int
	symbol string
}

func solveHomework(homework []homeworkAssignment) int {
	res := 0
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
		res += answer
	}
	return res
}

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

	removeDupeSpacesRE := regexp.MustCompile(`\s+`)

	homework := make([]homeworkAssignment, 0)
	l := 0
	linesPart2 := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}
		linesPart2 = append(linesPart2, t)

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

	part1 = solveHomework(homework)
	// Part 2

	symbolLine := linesPart2[len(linesPart2)-1]
	symbolLine = symbolLine[1:] // strip first
	for i:=0; i < len(homework); i++ {
		startLen := len(symbolLine)
		index := strings.IndexAny(symbolLine, "+*")
		colWidth := len(symbolLine)
		if index > -1 {
			symbolLine = symbolLine[index+1:]
			colWidth = startLen - len(symbolLine)
		} else {
			colWidth+=1
		}

		strs := []string{}
		for l := 0; l < len(linesPart2) - 1; l++ {
			strs = append(strs, linesPart2[l][0:colWidth])
			linesPart2[l] = linesPart2[l][colWidth:]
		}
		homework[i].values = homework[i].values[:0]
		for c := 0; c < len(strs[0]); c++ {
			built := ""
			for _, s := range strs {
				built = built + string(s[c])
			}
			n, _ := strconv.Atoi(strings.TrimSpace(built))
			if n != 0 {
				homework[i].values = append(homework[i].values, n)
			}
		}
	}

	part2 = solveHomework(homework)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 6 time:", time.Since(start))
}
