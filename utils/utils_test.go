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

func TestMakeLinkedList(t *testing.T) {
	assert := assert.New(t)
	{
		cut := MakeLinkedList([]int{1, 2, 3, 4, 5})
		assert.Equal(1, cut.Val)
		cut = cut.Next
		assert.Equal(2, cut.Val)
		cut = cut.Next
		assert.Equal(3, cut.Val)
		cut = cut.Next
		assert.Equal(4, cut.Val)
		cut = cut.Next
		assert.Equal(5, cut.Val)
		cut = cut.Next
		assert.Nil(cut)
	}
	{
		cut := MakeLinkedList([]int{})
		assert.Nil(cut)
	}
	{
		cut := MakeLinkedList([]int{1})
		assert.Equal(1, cut.Val)
		assert.Nil(cut.Next)
	}
}
