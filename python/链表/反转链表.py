def reverseListNOde(head):
    # 双指针法
    if head is None or head.next is None:
        return head

    cur = head
    pre = None
    while cur:
        temp = cur.next
        cur.next = pre
        pre = cur
        cur = temp
    return pre
