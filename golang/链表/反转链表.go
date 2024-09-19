package main

func reverseListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	per := &ListNode{}
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = per
		per = cur
		cur = tmp
	}

	return per
}
