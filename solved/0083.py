# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        tmp = head
        if head == None:
            return head

        if tmp.next == None:
            return head

        # tmp -> next
        next = head.next
        while tmp != None:
            if next == None:
                tmp.next = None
                return head

            if next.next == None:
                if tmp.val == next.val:
                    tmp.next = None
                return head

            while tmp.val == next.val:
                if next.next == None:
                    tmp.next = None
                    return head
                tmp.next = next.next
                next = tmp.next

            tmp = tmp.next
            next = next.next