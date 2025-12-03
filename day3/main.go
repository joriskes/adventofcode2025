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

func getBiggestJoltage(values []int, numberOfBatteries int, indexToSearchFrom int) int  {
	if numberOfBatteries == 0 {
		return 0
	}
	biggestJoltage := 0
	biggestIndex := 0

	for i := indexToSearchFrom; i <= len(values)-numberOfBatteries; i++ {
		if(values[i] > biggestJoltage) {
			biggestJoltage = values[i]
			biggestIndex = i
		}
	}
	biggestJoltage = biggestJoltage * int(math.Pow(10,float64(numberOfBatteries - 1)))
	return biggestJoltage + getBiggestJoltage(values, numberOfBatteries - 1, biggestIndex + 1)
}

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
		part1 += getBiggestJoltage(values, 2, 0)
		part2 += getBiggestJoltage(values, 12, 0)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 3 time:", time.Since(start))
}
