class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def removeElements(head, val):
    dummy = ListNode(0)
    dummy.next = head
    cur = dummy
    while cur.next:
        if cur.next.val == val:
            cur.next = cur.next.next
        else:
            cur = cur.next
    return dummy.next
