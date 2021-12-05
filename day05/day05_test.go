package day05

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2021/utils"
)

var (
	input = utils.ReadStrings("../inputs/day05.txt")
)

type Point struct {
	x, y int
}

// DrawVector draws a vector using Bresenham's line algorithm
func DrawVector(from, to Point) []Point {
	points := make([]Point, 0)

	x, y := from.x, from.y

	var dx, sx int
	if from.x < to.x {
		sx = 1
		dx = to.x - from.x
	} else {
		sx = -1
		dx = from.x - to.x
	}

	var dy, sy int
	if from.y < to.y {
		sy = 1
		dy = from.y - to.y
	} else {
		sy = -1
		dy = to.y - from.y
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

func Test_part1(t *testing.T) {
	points := map[Point]int{}

	for _, line := range input {
		fields := strings.Split(line, " -> ")
		p1 := strings.Split(fields[0], ",")
		p2 := strings.Split(fields[1], ",")
		p1x, _ := strconv.Atoi(p1[0])
		p1y, _ := strconv.Atoi(p1[1])
		p2x, _ := strconv.Atoi(p2[0])
		p2y, _ := strconv.Atoi(p2[1])

		point1 := Point{p1x, p1y}
		point2 := Point{p2x, p2y}
		if point1.x == point2.x || point1.y == point2.y { // only v and h lines
			for _, p := range DrawVector(point1, point2) {
				points[p] += 1
			}
		}
	}

	count := 0
	for _, v := range points {
		if v >= 2 {
			count++
		}
	}

	t.Log("Result:", count) // 5608
}

func Test_part2(t *testing.T) {
	points := map[Point]int{}

	for _, line := range input {
		fields := strings.Split(line, " -> ")
		p1 := strings.Split(fields[0], ",")
		p2 := strings.Split(fields[1], ",")
		p1x, _ := strconv.Atoi(p1[0])
		p1y, _ := strconv.Atoi(p1[1])
		p2x, _ := strconv.Atoi(p2[0])
		p2y, _ := strconv.Atoi(p2[1])

		point1 := Point{p1x, p1y}
		point2 := Point{p2x, p2y}

		for _, p := range DrawVector(point1, point2) {
			points[p] += 1
		}
	}

	count := 0
	for _, v := range points {
		if v >= 2 {
			count++
		}
	}

	t.Log("Result:", count) // 20299
}
