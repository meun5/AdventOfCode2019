package main

import (
	"math"
	"strconv" // Comment out for Part 1
	"fmt"
)

const (
	PWD_MIN = 128392
	PWD_MAX = 643281
)

func digit(num, place int) int {
    r := num % int(math.Pow(10, float64(place)))
    return r / int(math.Pow(10, float64(place-1)))
}

func main()  {
	c := 0

	for x := PWD_MIN; x < PWD_MAX; x++ {
		var isValid bool
		
		// Part 1
		for i := 1; i < 6; i++ {
			if digit(x, i) == digit(x, i+1) {
				isValid = true
				break
			}
		}

		for i := 1; i < 6; i++ {
			if digit(x, i) < digit(x, i+1) {
				isValid = false
				break
			}
		}

		// Part 2
		if isValid {
			isValid = func(x int) bool {
				counts := map[rune]int{}
				for _, s := range strconv.Itoa(x) {
					counts[s]++
				}
			
				for _, v := range counts {
					if v == 2 {
						return true
					}
				}
			
				return false
			}(x)
		}
		/**/
		
		if isValid {
			c++
		}
	}

	fmt.Println(c)
}