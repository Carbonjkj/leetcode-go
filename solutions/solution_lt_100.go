package solutions

import (
	"strings"
)

// 1
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		sub := target - v
		if m[v] != 0 {
			return []int{m[v] - 1, i}
		} else {
			m[sub] = i + 1
		}
	}
	return []int{}
}

// 9
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	num := x
	for num != 0 {
		rev = rev*10 + num%10
		num /= 10
	}
	return rev == x
}

// 12
func IntToRoman(num int) string {
	v := []rune{'I', 'V', 'X', 'L', 'C', 'D', 'M', ' '}
	rc := ""
	cursor := 0
	for num > 0 {
		var sb strings.Builder
		mod := num % 10
		switch mod {
		case 9:
			sb.WriteRune(v[2*cursor])
			sb.WriteRune(v[2*cursor+2])
			sb.WriteString(rc)
			mod = 0
		case 4:
			sb.WriteRune(v[2*cursor])
			sb.WriteRune(v[2*cursor+1])
			sb.WriteString(rc)
			mod = 0
		default:
			for mod > 0 {
				if mod-5 >= 0 {
					sb.WriteRune(v[2*cursor+1])
					mod -= 5
				}
				if mod-1 >= 0 {
					sb.WriteRune(v[2*cursor])
					mod -= 1
				}
			}
			sb.WriteString(rc)
		}
		rc = sb.String()
		num /= 10
		cursor++
	}
	return rc
}

// 13
func RomanToInt(s string) int {
	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	rc := 0
	last := 0
	for _, c := range s {
		rc += m[c]
		if m[c] > last {
			rc -= last * 2
		}
		last = m[c]
	}
	return rc
}
