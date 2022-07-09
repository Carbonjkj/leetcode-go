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
	cut = 2147483647
	assert.False(IsPalindrome(cut))
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

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3, 3, 4, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	cut := RemoveDuplicates(nums)
	assert.Equal(len(expect), cut)
	for i := 0; i < cut; i++ {
		assert.Equal(expect[i], nums[i])
	}
}

func TestRemoveElement(t *testing.T) {
	assert := assert.New(t)
	nums := []int{1, 2, 3, 3, 4, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9}
	expect := []int{1, 2, 3, 3, 5, 5, 6, 6, 7, 7, 8, 8, 9}
	cut := RemoveElement(nums, 4)
	assert.Equal(len(expect), cut)
	for i := 0; i < cut; i++ {
		assert.Equal(expect[i], nums[i])
	}
}

func TestStrStr(t *testing.T) {
	assert := assert.New(t)
	{
		haystack := "hakunamatata"
		needle := "una"

		assert.Equal(3, StrStrNormal(haystack, needle))
		assert.Equal(3, StrStrKMP(haystack, needle))
	}
	{
		haystack := "aaa"
		needle := "aaaa"

		assert.Equal(-1, StrStrNormal(haystack, needle))
		assert.Equal(-1, StrStrKMP(haystack, needle))
	}
	{
		haystack := "mississippi"
		needle := "issipi"
		assert.Equal(-1, StrStrNormal(haystack, needle))
		assert.Equal(-1, StrStrKMP(haystack, needle))
	}
	{
		haystack := "mississippi"
		needle := "pi"
		assert.Equal(9, StrStrNormal(haystack, needle))
		assert.Equal(9, StrStrKMP(haystack, needle))
	}
}

func TestSearchInsert(t *testing.T) {
	assert := assert.New(t)
	{
		array := []int{1, 2, 3, 4, 5, 6, 7}
		assert.Equal(7, SearchInsert(array, 8))
		assert.Equal(0, SearchInsert(array, 1))
		assert.Equal(2, SearchInsert(array, 3))
	}
}
