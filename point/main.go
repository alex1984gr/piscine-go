package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('_')
		n = -n
	}
	var digits []rune
	for n > 0 {
		d := rune(n%10 + '0')
		digits = append([]rune{d}, digits...)
		n /= 10
	}
	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	PrintStr("x = ")
	PrintInt(points.x)
	PrintStr(", y = ")
	PrintInt(points.y)
	z01.PrintRune('\n')
}
