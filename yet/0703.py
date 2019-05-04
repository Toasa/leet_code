# TLE

class Node:
    def __init__(self, val):
        self.val = val
        self.next = None

class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.k = k

        nums.sort(reverse=True)
        if len(nums) == 0:
            self.head = None
        else:
            self.head = Node(nums[0])
        tmp = self.head
        for item in nums[1:]:
            next = Node(item)
            tmp.next = next
            tmp = next
        
    def add(self, val: int) -> int:
        if self.head == None:
            self.head = Node(val)
        elif self.head.val <= val:
            n = Node(val)
            n.next = self.head
            self.head = n
        elif self.head.next == None:
                n = Node(val)
                self.head.next = n
        else:
            tmp = self.head
            while True:
                if tmp.next == None:
                    n = Node(val)
                    tmp.next = n
                    break
                elif tmp.val >= val and val >= tmp.next.val:
                    n = Node(val)
                    n.next = tmp.next
                    tmp.next = n
                    break
                tmp = tmp.next

        # linked listの表示
        # tmp = self.head
        # while tmp != None:
        #     print(tmp.val, "->", end="")
        #     tmp = tmp.next
        # print()
        
        tmp = self.head
        for i in range(self.k - 1):
            tmp = tmp.next
        return tmp.val

                