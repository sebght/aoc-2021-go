package days

import (
	"aoc-2021-go/utils"
	"testing"
)

func TestFirstPartDay1OneIncrease(t *testing.T) {
	result := FirstPartDay1([]int{1, 2})
	if result != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1)
	}
}

func TestFirstPartDay1OneDecrease(t *testing.T) {
	result := FirstPartDay1([]int{2, 1})
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestFirstPartDay1MultipleIncreases(t *testing.T) {
	result := FirstPartDay1([]int{1, 2, 4, 3})
	if result != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 2)
	}
}

func TestFirstPartDay1RealData(t *testing.T) {
	result := FirstPartDay1(utils.ReadDay1("../inputs/day1.txt"))
	if result != 1228 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1228)
	}
}

func TestSecondPartDay2TwoWindowsOneIncrease(t *testing.T) {
	result := SecondPartDay1([]int{199, 200, 208, 210})
	if result != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1)
	}
}

func TestSecondPartDay2TwoWindowsOneDecrease(t *testing.T) {
	result := SecondPartDay1([]int{208, 210, 200, 207})
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestSecondPartDay2OfficialExample(t *testing.T) {
	result := SecondPartDay1([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if result != 5 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 5)
	}
}
