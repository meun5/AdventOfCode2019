package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calculate(mass int) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var results []int

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		r, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		// Part 1
		// t := calculate(r)

		//Part 2
		t := recurse(r)

		results = append(results, t)
		fmt.Printf("Result: %d\n", t)
	}

	result := 0

	for _, item := range results {
		result += item
	}

	fmt.Println("Result |")
	fmt.Println(result)
}

func recurse(input int) int {
	i := calculate(input)
	r := 0

	fmt.Printf("Initial i is %d | r is %d\n", i, r)

	for ; i > 0; i = calculate(i) {
		fmt.Printf("i is %d | r is %d\n", i, r)
		r += i
	}

	fmt.Printf("Final i is %d | r is %d\n", i, r)

	return r
}
