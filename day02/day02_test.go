package day02

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2021/utils"
)

var (
	input = utils.ReadStrings("../inputs/day02.txt")
)

func Test_part1(t *testing.T) {
	pos := 0
	depth := 0

	for _, line := range input {
		fields := strings.Fields(line)
		direction := fields[0]
		amount, _ := strconv.Atoi(fields[1])

		switch direction {
		case "forward":
			pos += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	t.Log("Result:", pos*depth) // 1451208
}

func Test_part2(t *testing.T) {
	pos := 0
	depth := 0
	aim := 0

	for _, line := range input {
		parts := strings.Fields(line)
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			pos += amount
			depth += amount * aim
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	t.Log("Result:", pos*depth) // 1620141160
}
