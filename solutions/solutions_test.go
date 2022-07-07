package solutions

import (
	"leetcode/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	array := []int{-5, -1, 2, 7, 11, 15}
	got := TwoSum(array, 9)
	want := []int{2, 3}

	assert.Equal(t, want, got)
}

func TestIsPalindrome(t *testing.T) {
	assert := assert.New(t)
	cut := 121
	assert.True(IsPalindrome(cut))
	cut = 122
	assert.False(IsPalindrome(cut))
	cut = -121
	assert.False(IsPalindrome(cut))
	cut = 123242321
	assert.True(IsPalindrome(cut))
}

func TestRomanToInt(t *testing.T) {
	assert := assert.New(t)
	{
		cut := "III"
		assert.Equal(3, RomanToInt(cut))
	}
	{
		cut := "LVIII"
		assert.Equal(58, RomanToInt(cut))
	}
	{
		cut := "MCMXCIV"
		assert.Equal(1994, RomanToInt(cut))
	}
}

func TestIntToRoman(t *testing.T) {
	assert := assert.New(t)
	{
		cut := 3
		assert.Equal("III", IntToRoman(cut))
	}
	{
		cut := 58
		assert.Equal("LVIII", IntToRoman(cut))
	}
	{
		cut := 1994
		assert.Equal("MCMXCIV", IntToRoman(cut))
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	assert := assert.New(t)
	{
		cut := []string{"flower", "flow", "flight"}
		assert.Equal("fl", LongestCommonPrefix(cut))
	}
	{
		cut := []string{"applea", "applb", "applec", "app"}
		assert.Equal("app", LongestCommonPrefix(cut))
	}
	{
		cut := []string{"dog", "racecar", "car"}
		assert.Equal("", LongestCommonPrefix(cut))
	}
}

func TestIsValidParentheses(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsValidParentheses("()(){}{[()]}"))
	assert.False(IsValidParentheses("()[(])){}{[()]}"))
	assert.False(IsValidParentheses("["))
}

func TestMergeTwoLists(t *testing.T) {
	assert := assert.New(t)
	{
		list1 := utils.MakeLinkedList([]int{1, 3, 5, 34, 55})
		list2 := utils.MakeLinkedList([]int{1, 2, 8, 13, 21})
		cut := MergeTwoLists(list1, list2)
		assert.Equal(1, cut.Val)
		cut = cut.Next
		assert.Equal(1, cut.Val)
		cut = cut.Next
		assert.Equal(2, cut.Val)
		cut = cut.Next
		assert.Equal(3, cut.Val)
		cut = cut.Next
		assert.Equal(5, cut.Val)
		cut = cut.Next
		assert.Equal(8, cut.Val)
		cut = cut.Next
		assert.Equal(13, cut.Val)
		cut = cut.Next
		assert.Equal(21, cut.Val)
		cut = cut.Next
		assert.Equal(34, cut.Val)
		cut = cut.Next
		assert.Equal(55, cut.Val)
		cut = cut.Next
		assert.Nil(cut)
	}
	{
		list1 := utils.MakeLinkedList([]int{})
		list2 := utils.MakeLinkedList([]int{1})
		cut := MergeTwoLists(list1, list2)
		assert.Equal(1, cut.Val)
		assert.Nil(cut.Next)
	}
}
