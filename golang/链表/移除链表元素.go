package main

func removeElement(head *ListNode, val int) {
	dummy := &ListNode{Next: head}

	for cur := dummy; cur != nil && cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
}

func removeElement1(head *ListNode, target int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == target {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
