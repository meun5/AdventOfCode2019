package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	// set "github.com/deckarep/golang-set"
)

// Point is where we are in the graph
type Point struct {
	x int
	y int
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	c := string(content)

	h := make([]map[Point]int, 5)

	for g, c := range strings.Split(c, "\n") {
		f := strings.Split(c, ",")

		h[g] = make(map[Point]int)
		r := h[g]

		k, y, u := 0, 0, 0
		for _, x := range f {
			switch x[0:1] {
			case "U":
				t, _ := strconv.Atoi(x[1:len(x)])

				for i := 0; i < t; i++ {
					y++
					r[Point{
						x: k,
						y: y,
					}] = u
					u++
				}
				break
			case "D":
				t, _ := strconv.Atoi(x[1:len(x)])

				for i := 0; i < t; i++ {
					y--
					r[Point{
						x: k,
						y: y,
					}] = u
					u++
				}
				break

			case "R":
				t, _ := strconv.Atoi(x[1:len(x)])

				for i := 0; i < t; i++ {
					k++
					r[Point{
						x: k,
						y: y,
					}] = u
					u++
				}
				break

			case "L":
				t, _ := strconv.Atoi(x[1:len(x)])

				for i := 0; i < t; i++ {
					k--
					r[Point{
						x: k,
						y: y,
					}] = u
					u++
				}
				break
			}
		}

		// fmt.Println(r)
	}

	e := intersect(h[0], h[1])

	z := make([]int, len(e))

	p := 0
	for o := range e {
		z[p] = Abs(o.x) + Abs(o.y)

		p++
	}

	sort.Ints(z)

	fmt.Printf("Part 1: %d\n", z[0])

	z = make([]int, len(e))

	p = 0
	for o := range e {
		z[p] = h[0][o] + h[1][o]

		p++
	}

	sort.Ints(z)

	// add two as I am not counting 0, 0 on the wires
	fmt.Printf("Part 2: %d\n", z[0]+2)
}

func intersect(map1, map2 map[Point]int) map[Point]int {
	v := make(map[Point]int)

	if len(map1) < len(map2) {
		for m := range map2 {
			if _, ok := map1[m]; ok {
				v[m] = 1
			}
		}
	} else {
		for m := range map1 {
			if _, ok := map2[m]; ok {
				v[m] = 1
			}
		}
	}

	return v
}

// Abs in integer absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
