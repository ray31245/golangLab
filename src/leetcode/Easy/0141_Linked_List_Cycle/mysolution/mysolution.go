package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	return hasCycle(head)
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowPoint := head.Next
	if slowPoint == nil {
		return false
	}
	fastPoint := head.Next.Next
	if fastPoint == nil {
		return false
	}
	if fastPoint == slowPoint {
		return true
	}
	for fastPoint != nil && fastPoint.Next != nil && slowPoint != head {
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
		if slowPoint == fastPoint {
			return true
		}
	}
	return false
}
