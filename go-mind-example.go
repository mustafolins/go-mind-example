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
	m := mind.New(0.3, 100000, 6, "sigmoid")
	m.Learn([][][]float64{
		{f, mapletter('f')},
		{e, mapletter('e')},
		{d, mapletter('d')},
		{c, mapletter('c')},
		{b, mapletter('b')},
		{a, mapletter('a')},
	})

	result := m.Predict([][]float64{
		character( // slightly resembles an 'A'
			"..####." +
				"#.....#" +
				"#.....#" +
				"#.#..##" +
				"#.....#" +
				"#.....#" +
				"#.....#")})
	log.Println(result)
	fmt.Println(closestletter(result.At(0, 0)))
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
	if letter == 'a' {
		return []float64{0.16}
	}
	if letter == 'b' {
		return []float64{0.32}
	}
	if letter == 'c' {
		return []float64{0.48}
	}
	if letter == 'd' {
		return []float64{0.64}
	}
	if letter == 'e' {
		return []float64{0.80}
	}
	if letter == 'f' {
		return []float64{1.0}
	}
	return []float64{0.0}
}

func closestletter(result float64) string {
	if result >= 0 && result < 0.24 {
		return "a"
	}
	if result >= 0.24 && result < 0.40 {
		return "b"
	}
	if result >= 0.40 && result < 0.56 {
		return "c"
	}
	if result >= 0.56 && result < 0.72 {
		return "d"
	}
	if result >= 0.72 && result < 0.88 {
		return "e"
	}
	if result >= 0.88 && result <= 1.0 {
		return "f"
	}
	return ""
}
