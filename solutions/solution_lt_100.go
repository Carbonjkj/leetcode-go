package solutions

import (
	"fmt"
	"leetcode/utils"
	"os"
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

// 26
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[last] {
			nums[last+1] = nums[i]
			last++
		}
	}
	return last + 1
}

// 27
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	ptr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ptr] = nums[i]
			ptr++
		}
	}
	return ptr
}

// 28 31.27% 67.73%
func StrStrNormal(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	match := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			match = i
			for j := 0; j < len(needle); j++ {
				if match+j >= len(haystack) {
					return -1
				}
				if haystack[match+j] != needle[j] {
					match++
					break
				}
			}
			if match == i {
				return match
			} else {
				i = match - 1
			}
		}
	}
	return -1
}

// 28 100.00% 7.09%
func StrStrKMP(haystack string, needle string) int {
	fmt.Fprintf(os.Stdout, "=========\n")
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	var lps = make([]int, len(needle))
	pos := 0
	lps[0] = 0
	i := 1
	for i < len(lps) {
		if needle[i] == needle[pos] {
			pos++
			lps[i] = pos
			i++
		} else {
			if pos == 0 {
				lps[i] = 0
				i++
			} else {
				pos = lps[pos-1]
			}
		}
	}
	x := 0
	y := 0
	for x < len(haystack) {
		fmt.Fprintf(os.Stdout, "current %c \n", needle[y])
		if needle[y] == haystack[x] {
			x++
			y++
		}
		if y == len(needle) {
			return x - y
		} else if x < len(haystack) && needle[y] != haystack[x] {
			if y != 0 {
				y = lps[y-1]
				fmt.Fprintf(os.Stdout, "restart from %c \n", needle[y])
			} else {
				x++
			}
		}
	}
	return -1
}
