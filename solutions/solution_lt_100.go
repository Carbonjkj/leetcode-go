package solutions

import (
	"leetcode/utils"
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
	v := []string{"I", "V", "X", "L", "C", "D", "M", " "}
	rc := ""
	cursor := 0
	for num > 0 {
		mod := num % 10
		switch mod {
		case 9:
			rc = v[2*cursor] + v[2*cursor+2] + rc
		case 4:
			rc = v[2*cursor] + v[2*cursor+1] + rc
		default:
			append := ""
			for mod > 0 {
				if mod-5 >= 0 {
					append = v[2*cursor+1]
					mod -= 5
				}
				if mod-1 >= 0 {
					append = append + v[2*cursor]
					mod -= 1
				}
			}
			rc = append + rc
		}
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

// 14
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longest := strs[0]
	for _, value := range strs {
		for i, c := range longest {
			if i >= len(value) || c != rune(value[i]) {
				longest = longest[:i]
				break
			}
		}
	}
	return longest
}

// 20
func IsValidParentheses(s string) bool {
	str := ""
	rc := ' '
	for _, c := range s {
		test := false
		switch c {
		case '(', '[', '{':
			str = string(c) + str
		case ')':
			rc = '('
			test = true
		case ']':
			rc = '['
			test = true
		case '}':
			rc = '{'
			test = true
		}
		if !test {
			continue
		}
		if len(str) == 0 {
			return false
		}
		if rc != rune(str[0]) {
			return false
		}
		if len(str) == 1 {
			str = ""
		} else {
			str = str[1:]
		}
	}
	return (len(str) == 0)
}

// 21
func MergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var origin, append, rc *utils.ListNode
	if list1.Val <= list2.Val {
		origin = list1
		rc = list1
		append = list2
	} else {
		origin = list2
		rc = list2
		append = list1
	}
	for origin.Next != nil && append != nil {
		if append.Val >= origin.Val && append.Val < origin.Next.Val {
			head := append
			append = append.Next
			head.Next = origin.Next
			origin.Next = head
		}
		origin = origin.Next
	}
	if append != nil {
		origin.Next = append
	}
	return rc
}
