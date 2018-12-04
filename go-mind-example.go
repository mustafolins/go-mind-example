package main

import (
	"fmt"
	"log"

	"github.com/stevenmiller888/go-mind"
)

var (
	a = character(
		".#####." +
			"#.....#" +
			"#.....#" +
			"#######" +
			"#.....#" +
			"#.....#" +
			"#.....#")
	b = character(
		"######." +
			"#.....#" +
			"#.....#" +
			"######." +
			"#.....#" +
			"#.....#" +
			"######.")
	c = character(
		"#######" +
			"#......" +
			"#......" +
			"#......" +
			"#......" +
			"#......" +
			"#######")
	d = character(
		"#####.." +
			"#....#." +
			"#.....#" +
			"#.....#" +
			"#.....#" +
			"#....#." +
			"#####..")
	e = character(
		"#######" +
			"#......" +
			"#......" +
			"#######" +
			"#......" +
			"#......" +
			"#######")
	f = character(
		"#######" +
			"#......" +
			"#......" +
			"#####.." +
			"#......" +
			"#......" +
			"#......")
)

func main() {
	m := mind.New(0.3, 1000000, 6, "sigmoid")
	m.Learn([][][]float64{
		{f, mapletter('F')},
		{e, mapletter('E')},
		{d, mapletter('D')},
		{c, mapletter('C')},
		{b, mapletter('B')},
		{a, mapletter('A')},
	})

	result := m.Predict([][]float64{
		character( // slightly resembles an 'A'
			"######." +
				"#......" +
				"#......" +
				"#.##.#." +
				"#......" +
				"#..#..." +
				"#......")})
	log.Println(result)
	fmt.Println(closestasciiart(result.At(0, 0)))
}
func character(chars string) []float64 {
	flt := make([]float64, len(chars))
	for i := 0; i < len(chars); i++ {
		if chars[i] == '#' {
			flt[i] = 1.0
		} else { // if '.'
			flt[i] = 0.0
		}
	}
	return flt
}

func mapletter(letter byte) []float64 {
	if letter == 'A' {
		return []float64{0.16}
	}
	if letter == 'B' {
		return []float64{0.32}
	}
	if letter == 'C' {
		return []float64{0.48}
	}
	if letter == 'D' {
		return []float64{0.64}
	}
	if letter == 'E' {
		return []float64{0.80}
	}
	if letter == 'F' {
		return []float64{1.0}
	}
	return []float64{0.0}
}

func closestasciiart(result float64) string {
	if result >= 0 && result < 0.24 {
		return "A"
	}
	if result >= 0.24 && result < 0.40 {
		return "B"
	}
	if result >= 0.40 && result < 0.56 {
		return "C"
	}
	if result >= 0.56 && result < 0.72 {
		return "D"
	}
	if result >= 0.72 && result < 0.88 {
		return "E"
	}
	if result >= 0.88 && result <= 1.0 {
		return "F"
	}
	return ""
}
