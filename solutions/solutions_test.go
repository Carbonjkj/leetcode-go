package solutions

import (
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
