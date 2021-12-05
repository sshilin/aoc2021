package day01

import (
	"testing"

	"github.com/sshilin/aoc2021/utils"
)

var (
	input = utils.ReadInts("../inputs/day01.txt")
)

func Test_part1(t *testing.T) {
	count := 0
	prev := input[0]
	for _, i := range input[1:] {
		if i > prev {
			count++
		}
		prev = i
	}

	t.Log("Result:", count) // 1616
}

func Test_part2(t *testing.T) {
	count := 0
	prev := sum(input[0:3])
	for i := 1; i < len(input)-2; i++ {
		s := sum(input[i : i+3])
		if s > prev {
			count++
		}
		prev = s
	}

	t.Log("Result:", count) // 1645
}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
