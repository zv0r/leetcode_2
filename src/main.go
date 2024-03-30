package main

import "fmt"

func main() {
	l1 := buildNumber([]int{2, 4, 3})
	l2 := buildNumber([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)
	printNumber(l3)

	l4 := buildNumber([]int{9, 9, 9, 9})
	l5 := buildNumber([]int{4, 4, 5})
	l6 := addTwoNumbers(l4, l5)
	printNumber(l6)

	l7 := buildNumber([]int{1})
	l8 := buildNumber([]int{1})
	l9 := addTwoNumbers(l7, l8)
	printNumber(l9)

	l10 := buildNumber([]int{8})
	l11 := buildNumber([]int{4})
	l12 := addTwoNumbers(l10, l11)
	printNumber(l12)

	l13 := buildNumber([]int{0})
	l14 := buildNumber([]int{0})
	l15 := addTwoNumbers(l13, l14)
	printNumber(l15)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printNumber(num *ListNode) {
	for num != nil {
		fmt.Print(num.Val, " ")
		num = num.Next
	}
	fmt.Println()
}

func buildNumber(nums []int) *ListNode {
	var root *ListNode = &ListNode{Val: 0, Next: nil}
	root.Val = nums[0]
	var node *ListNode = root
	for _, i := range nums[1:] {
		node.Next = &ListNode{Val: i, Next: nil}
		node = node.Next
	}
	return root
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	transfer := 0
	root := &ListNode{Val: 0, Next: nil}
	curr_node := root
	prev_node := root

	for l1 != nil || l2 != nil {
		l1_val := 0
		l2_val := 0
		if l1 != nil {
			l1_val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2_val = l2.Val
			l2 = l2.Next
		}
		curr_node.Val = (l1_val + l2_val + transfer) % 10
		transfer = (l1_val + l2_val + transfer) / 10
		prev_node = curr_node
		curr_node.Next = &ListNode{Val: transfer, Next: nil}
		curr_node = curr_node.Next
	}

	if curr_node.Val == 0 {
		prev_node.Next = nil
	}

	return root
}
