def swapPairs(head):
    dummy = ListNode(0)
    dummy.next = head
    cur = dummy
    while cur.next and cur.next.next:
        first = cur.next
        second = cur.next.next
        cur.next = second
        first.next = second.next
        second.next = first
        cur = cur.next.next
    return dummy.next
