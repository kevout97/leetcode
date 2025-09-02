package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0
	var aux *ListNode = &ListNode{}
	total := aux

	var tmpL1 = l1
	var tmpL2 = l2

	for tmpL1 != nil || tmpL2 != nil || carry > 0 {

		var sum = carry

		if tmpL1 != nil {
			sum += tmpL1.Val
			tmpL1 = tmpL1.Next
		}

		if tmpL2 != nil {
			sum += tmpL2.Val
			tmpL2 = tmpL2.Next
		}

		totalValue := sum % 10
		carry = sum / 10

		aux.Next = &ListNode{
			Val:  totalValue,
			Next: nil,
		}

		aux = aux.Next

	}

	return total.Next

}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	totalList := []int{}
	total := addTwoNumbers(l1, l2)

	for total != nil {
		totalList = append(totalList, total.Val)
		total = total.Next
	}

	fmt.Printf("%v", totalList)
}
