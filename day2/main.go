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

func isValidIdPart1(id int) bool {
	sId := strconv.Itoa(id)
	l := len(sId)
	if l % 2 != 0 {
		return true
	}
	l = l / 2
	return sId[:l] != sId[l:]
}

func isValidIdPart2(id int) bool {
	sId := strconv.Itoa(id)
	l := len(sId)
	for i := 1; i <= int(math.Ceil(float64(l / 2))); i++ {
		expected := l / i
		if l % i != 0 {
			continue
		}
		if strings.Count(sId, sId[0:i]) == expected {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("====== Day 2 ======")
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

	scanner.Scan()
	t := scanner.Text()
	groups := strings.Split(t, ",")

	for _, group := range groups {
		idRange := strings.Split(group, "-");
		from, _ := strconv.Atoi(idRange[0])
		to, _ := strconv.Atoi(idRange[1])

		for n := from; n <= to; n++ {
			if !isValidIdPart1(n) {
				part1 += n
			}
			if !isValidIdPart2(n) {
				part2 += n
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 2 time:", time.Since(start))
}
