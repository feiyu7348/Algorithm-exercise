package listNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m, n := 0, 0
	nodeA, nodeB := headA, headB
	for nodeA != nil {
		nodeA = nodeA.Next
		m++
	}

	for nodeB != nodeA {
		nodeB = nodeB.Next
		n++
	}

	var step int
	var slow, fast *ListNode
	if m > n {
		step = m - n
		fast, slow = headA, headB
	} else {
		step = n - m
		fast, slow = headB, headA
	}

	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}
