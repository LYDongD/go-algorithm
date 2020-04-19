package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	dummy := &ListNode{0, nil}
	p := dummy
	carry := 0
	p1, p2 := l1, l2
	for p1 != nil || p2 != nil {

		//cal sum & update p1,p2
		sum := 0
		if p1 == nil {
			sum += p2.Val
			p2 = p2.Next
		} else if p2 == nil {
			sum += p1.Val
			p1 = p1.Next
		} else {
			sum += p1.Val + p2.Val
			p1, p2 = p1.Next, p2.Next
		}

		//add cary & update carry
		sum += carry
		if sum > 9 {
			sum = sum % 10
			carry = 1
		} else {
			carry = 0
		}

		//add sum node
		p.Next = &ListNode{sum, nil}
		p = p.Next
	}

	//handle carry at last Node
	if carry > 0 {
		p.Next = &ListNode{carry, nil}
	}

	return dummy.Next
}

func array2Link(arr []int) *ListNode {
	var head, cursor *ListNode
	for _, num := range arr {
		if head == nil {
			head = &ListNode{num, nil}
			cursor = head
		} else {
			cursor.Next = &ListNode{num, nil}
			cursor = cursor.Next
		}
	}

	return head
}

func link2Array(head *ListNode) []int {
	result := []int{}
	cursor := head
	for cursor != nil {
		result = append(result, cursor.Val)
		cursor = cursor.Next
	}

	return result
}

func main() {
	//testArr := []int{2, 4, 3}
	//testLink := array2Link(testArr)
	//fmt.Println(link2Array(testLink))
	sumLink := addTwoNumbers(array2Link([]int{2, 4, 3}), array2Link([]int{5, 6, 4}))
	fmt.Println(link2Array(sumLink))
}
