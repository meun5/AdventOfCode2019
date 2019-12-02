package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func wave(t *[]string) {
	r := (*t)
	for i := 0; i < len(r); {
		// q := s.Pop()

		// fmt.Printf("Have %s\n", r[i])

		switch r[i] {
		case "1":
			op1, _ := strconv.Atoi(r[i+1])
			op2, _ := strconv.Atoi(r[i+2])
			op3, _ := strconv.Atoi(r[i+3])

			i1, _ := strconv.Atoi(r[op1])
			i2, _ := strconv.Atoi(r[op2])
			r[op3] = strconv.Itoa(i1 + i2)

			i += 4

			break
		case "2":
			op1, _ := strconv.Atoi(r[i+1])
			op2, _ := strconv.Atoi(r[i+2])
			op3, _ := strconv.Atoi(r[i+3])

			i1, _ := strconv.Atoi(r[op1])
			i2, _ := strconv.Atoi(r[op2])
			r[op3] = strconv.Itoa(i1 * i2)

			i += 4

			break
		case "99":
			return
		default:
			i++
			break
		}
	}
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	f := string(content)
	t := strings.Split(f, ",")

	// Part 1
	// wave(&t)
	// fmt.Printf("T0 is %s\n", t[0])

	// Part 2
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			fmt.Printf("I: %d, J: %d\n", i, j)

			// Encode into string
			t[1] = strconv.Itoa(i)
			t[2] = strconv.Itoa(j)

			wave(&t)

			fmt.Printf("T0 is %s\n", t[0])

			if t[0] != "19690720" {
				// Reset t
				t = strings.Split(f, ",")
			} else {
				fmt.Println("T0 is 19690720")

				fmt.Printf("I is %d, J is %d", i, j)

				return
			}
		}
	}
}
