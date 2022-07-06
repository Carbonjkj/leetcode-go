package solutions

import (
	"leetcode/utils"
	"testing"
)

func TestTwoSum(t *testing.T) {
	array := []int{-5, -1, 2, 7, 11, 15}
	got := TwoSum(array, 9)
	want := []int{2, 3}

	if !utils.Equal(want, got) {
		t.Errorf("got %d, want %d", got, want)
	}
}
