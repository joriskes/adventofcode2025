package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const RUN_EXAMPLE = false

func main() {
	fmt.Println("====== Day 3 ======")
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

	for scanner.Scan() {
		t := scanner.Text()

		tmp := strings.Split(t, "")
		values := make([]int, 0, len(tmp))
		for _, raw := range tmp {
			v, err := strconv.Atoi(raw)
			if err != nil {
				log.Print(err)
				continue
			}
			values = append(values, v)
		}

		lineSum := 0
		s:=0
		for i := 0; i < len(values); i += 1 {
			for j := i + 1; j < len(values); j += 1 {
				s = (values[i] * 10) + values[j]
				if lineSum < s {
					lineSum = s
				}
			}
		}
		part1 += lineSum
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 3 time:", time.Since(start))
}
