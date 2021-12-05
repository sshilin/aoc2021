package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInts(name string) []int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	ints := make([]int, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			ints = append(ints, i)
		} else {
			log.Fatalln(err)
		}
	}
	return ints
}

func ReadStrings(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	lines := make([]string, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type Point struct {
	x, y int
}

func Abs(val int) int {
	if val < 0 {
		return -(val)
	}
	return val
}

// DrawVector draws a vector using Bresenham's line algorithm
func DrawVector(from, to Point) []Point {
	points := make([]Point, 0)

	x, y := from.x, from.y

	dx := Abs(to.x - from.x)
	sx := -1
	if from.x < to.x {
		sx = 1
	}

	dy := -Abs(to.y - from.y)
	sy := -1
	if from.y < to.y {
		sy = 1
	}

	err := dx + dy

	for {
		points = append(points, Point{x, y})
		if x == to.x && y == to.y {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x += sx
		}
		if e2 <= dx {
			err += dx
			y += sy
		}
	}

	return points
}
