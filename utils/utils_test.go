package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.Equal(t, arr1, arr2)
}
