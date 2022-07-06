package utils

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestEqual(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5}

	if !Equal(arr1, arr2) {
		t.Errorf("got false, want true")
	}
}
