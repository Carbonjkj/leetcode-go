package utils

func Add(x, y int) int {
	return x + y
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkedList(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	ln := new(ListNode)
	ln.Val = array[0]
	ln.Next = nil
	if len(array) == 1 {
		return ln
	}
	head := ln
	for _, v := range array[1:] {
		head.Next = new(ListNode)
		head = head.Next
		head.Val = v
		head.Next = nil
	}
	return ln
}
