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
	var quotient int = 0
	var index int = 0
	total := &ListNode{}
	var previousNode *ListNode

	var tmpL1 = l1
	var tmpL2 = l2

	for tmpL1 != nil || tmpL2 != nil {

		var val1 = 0
		var val2 = 0

		if tmpL1 != nil {
			val1 = tmpL1.Val
		}

		if tmpL2 != nil {
			val2 = tmpL2.Val
		}

		sum := val1 + val2 + quotient
		totalValue := sum % 10
		quotient = sum / 10

		if tmpL1 != nil && tmpL1.Next != nil {
			tmpL1 = tmpL1.Next
		} else {
			tmpL1 = nil
		}

		if tmpL2 != nil && tmpL2.Next != nil {
			tmpL2 = tmpL2.Next
		} else {
			tmpL2 = nil
		}

		if index == 0 {
			total.Val = totalValue
			total.Next = nil
			previousNode = total
			index++
			continue
		}

		currentNode := &ListNode{
			Val:  totalValue,
			Next: nil,
		}

		previousNode.Next = currentNode
		previousNode = currentNode

		index++
	}

	if quotient > 0 {
		previousNode.Next = &ListNode{
			Val:  quotient,
			Next: nil,
		}
	}

	return total

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
