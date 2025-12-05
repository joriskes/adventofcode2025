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

 func isFresh(freshIngredients []struct{from int; to int}, ingredient int) bool  {
	 for _, v := range freshIngredients {
		 if ingredient >= v.from  && ingredient <= v.to {
			 return true
		 }
	 }
	 return false
 }

func main() {
	fmt.Println("====== Day 5 ======")
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

	freshIngredients := make([]struct{from int; to int}, 0)
	inAvailable := true

	for scanner.Scan() {
		t := scanner.Text()
		if inAvailable {
			if t == "" {
				inAvailable = false
				continue
			}
			s := strings.Split(t, "-")
			from, _ := strconv.Atoi(s[0])
			to, _ := strconv.Atoi(s[1])
			freshIngredients = append(freshIngredients, struct{from int; to int}{from:from , to:to})
		} else {
			ingredient, _ := strconv.Atoi(t)
			if isFresh(freshIngredients, ingredient) {
				part1++
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("⏱️ Day 5 time:", time.Since(start))
}
