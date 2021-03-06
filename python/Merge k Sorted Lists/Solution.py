# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def mergeKLists(self, _lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        if len(_lists) == 0:
            return []

        lists = []
        for l in _lists:
            if l != None:
                lists.append(l)

        self.quickSort(lists, 0, len(lists) - 1)
        result = []
        while len(lists) > 0:
            result.append(lists[0].val)
            lists[0] = lists[0].next
            if lists[0] == None:
                lists = lists[1:]
            elif len(lists) > 0:
                self.sort(lists)
        return result

    def sort(self, array):
        for i in xrange(0,len(array) - 1t):
            if array[i].val <= array[i + 1].val:
                break;
            else:
                temp_node = array[i + 1]
                array[i + 1] = array[i]
                array[i] = temp_node

    def quickSort(self, array,first,last):
        if first < last:
            pos = self.split(array,first,last)
            self.quickSort(array,first,pos-1)
            self.quickSort(array,pos+1,last)
        return array

    def split(self, array,first,last):
        left  = first
        right = last
        while left < right:
            while array[first].val < array[right].val:
                right -= 1
            while array[first].val >= array[left].val and left < right:
                left += 1
            if left < right:
                array[left],array[right] = array[right],array[left]
        array[first],array[left] = array[left],array[first]
        return left

