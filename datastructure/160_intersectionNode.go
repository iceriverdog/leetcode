package datastructure

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	for curB != curA {   // 少判断
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}

	}
	return curA

}
