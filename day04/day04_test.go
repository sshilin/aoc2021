package day04

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sshilin/aoc2021/utils"
)

var (
	input = utils.ReadStrings("../inputs/day04.txt")
)

type Cell struct {
	num  int
	mark bool
}

type Board [5][5]Cell

func (brd *Board) Mark(num int) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if brd[r][c].num == num {
				brd[r][c].mark = true
			}
		}
	}
}

func (brd *Board) SumUnmarked() int {
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !brd[r][c].mark {
				sum += brd[r][c].num
			}
		}
	}

	return sum
}

func (brd *Board) Bingo() bool {
	// check rows
	for r := 0; r < 5; r++ {
		check := true
		for c := 0; c < 5; c++ {
			if !brd[r][c].mark {
				check = false
				break
			}
		}
		if check {
			return true
		}
	}

	// check columns
	for c := 0; c < 5; c++ {
		check := true
		for r := 0; r < 5; r++ {
			if !brd[r][c].mark {
				check = false
				break
			}
		}
		if check {
			return true
		}
	}

	return false
}

func readNums() []int {
	nums := make([]int, 0)
	for _, s := range strings.Split(input[0], ",") {
		i, _ := strconv.Atoi(s)
		nums = append(nums, i)
	}

	return nums
}

func readBoards() []*Board {
	boards := make([]*Board, 0)
	brd := &Board{}
	r := 0
	for i, line := range input[2:] {
		if line != "" {
			for c, v := range strings.Fields(line) {
				num, _ := strconv.Atoi(v)
				brd[r][c].num = num
			}
			r++
		}
		if line == "" || i == len(input[2:])-1 {
			boards = append(boards, brd)
			brd = &Board{}
			r = 0
		}
	}

	return boards
}

func Test_part1(t *testing.T) {
	nums := readNums()
	boards := readBoards()

	for _, num := range nums {
		for _, board := range boards {
			board.Mark(num)
		}
		for _, board := range boards {
			if win := board.Bingo(); win {
				t.Log("Result:", board.SumUnmarked()*num) // 39902
				return
			}
		}
	}
}

func Test_part2(t *testing.T) {
	nums := readNums()
	boards := readBoards()
	winnerBoards := map[int]bool{}

	for _, num := range nums {
		for _, board := range boards {
			board.Mark(num)
		}
		for i, board := range boards {
			if win := board.Bingo(); win {
				if ok := winnerBoards[i]; !ok {
					winnerBoards[i] = true
					t.Log("Result:", board.SumUnmarked()*num) // 26936
				}
			}
		}
	}
}
