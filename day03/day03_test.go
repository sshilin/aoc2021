package day03

import (
	"strconv"
	"testing"

	"github.com/sshilin/aoc2021/utils"
)

var (
	input = utils.ReadStrings("../inputs/day03.txt")
)

func readUints() []uint {
	ints := make([]uint, 0)
	for _, s := range input {
		i, _ := strconv.ParseUint(s, 2, len(s))
		ints = append(ints, uint(i))
	}

	return ints
}

func Test_part1(t *testing.T) {
	bitSize := len(input[0])
	uints := readUints()
	ones := make([]int, bitSize)

	for _, i := range uints {
		for b := 0; b < bitSize; b++ {
			if (i>>b)&0x01 == 1 {
				ones[b]++
			}
		}
	}

	var gamma uint
	var epsilon uint

	for b := 0; b < bitSize; b++ {
		if ones[b] > len(input)-ones[b] {
			gamma |= 1 << b
		} else {
			epsilon |= 1 << b
		}
	}

	t.Log("Result:", gamma*epsilon) // 4174964
}

func countOnes(uints []uint, bitSize int) []int {
	ones := make([]int, bitSize)
	for _, i := range uints {
		for b := 0; b < bitSize; b++ {
			if (i>>b)&0x01 == 1 {
				ones[bitSize-1-b]++
			}
		}
	}
	return ones
}

func Test_part2(t *testing.T) {
	bitSize := len(input[0])
	uints := readUints()
	oxygen := make([]uint, len(uints))
	co2 := make([]uint, len(uints))
	copy(oxygen, uints)
	copy(co2, uints)

	for b := 0; b < bitSize; b++ {
		ones := countOnes(oxygen, bitSize)
		oxygenNew := make([]uint, 0)
		if ones[b] > len(oxygen)-ones[b] { // most common 1
			for i := 0; i < len(oxygen); i++ {
				if (oxygen[i]>>(bitSize-1-b))&0x01 == 1 { // keep 1
					oxygenNew = append(oxygenNew, oxygen[i])
				}
			}
		} else if len(oxygen)-ones[b] > ones[b] { // most common 0
			for i := 0; i < len(oxygen); i++ {
				if (oxygen[i]>>(bitSize-1-b))&0x01 == 0 { // keep 0
					oxygenNew = append(oxygenNew, oxygen[i])
				}
			}
		} else if ones[b] == len(oxygen)-ones[b] { //  1 == 0
			for i := 0; i < len(oxygen); i++ { // keep 1
				if (oxygen[i]>>(bitSize-1-b))&0x01 == 1 {
					oxygenNew = append(oxygenNew, oxygen[i])
				}
			}
		}
		oxygen = oxygenNew

		if len(oxygen) == 1 {
			break
		}
	}

	for b := 0; b < bitSize; b++ {
		ones := countOnes(co2, bitSize)
		co2New := make([]uint, 0)
		if ones[b] < len(co2)-ones[b] { // less common 1
			for i := 0; i < len(co2); i++ {
				if (co2[i]>>(bitSize-1-b))&0x01 == 1 { // keep 1
					co2New = append(co2New, co2[i])
				}
			}
		} else if len(co2)-ones[b] < ones[b] { // less common 0
			for i := 0; i < len(co2); i++ {
				if (co2[i]>>(bitSize-1-b))&0x01 == 0 { // keep 0
					co2New = append(co2New, co2[i])
				}
			}
		} else if ones[b] == len(co2)-ones[b] { //  1 == 0
			for i := 0; i < len(co2); i++ { // keep 0
				if (co2[i]>>(bitSize-1-b))&0x01 == 0 {
					co2New = append(co2New, co2[i])
				}
			}
		}
		co2 = co2New

		if len(co2) == 1 {
			break
		}
	}

	t.Log("Result:", oxygen[0]*co2[0]) // 4474944
}
