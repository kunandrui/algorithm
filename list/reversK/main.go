package reversK

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	protect := &ListNode{0, head}
	lastGroupEnd := protect
	//1. 分组
	currGroupHead := head
	for currGroupHead != nil {
		currGroupEnd := getGoupEnd(currGroupHead, k)
		if currGroupEnd == nil {
			break
		}
		//2. 翻转链表
		nextGoupHead := currGroupEnd.Next
		reverse(currGroupHead, nextGoupHead)

		//3. 连接
		lastGroupEnd.Next = currGroupEnd
		currGroupHead.Next = nextGoupHead

		lastGroupEnd = currGroupHead
		currGroupHead = nextGoupHead
	}
	return protect.Next
}

func getGoupEnd(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

func reverse(head *ListNode, tail *ListNode) {
	last := head
	head = head.Next
	for head != tail {
		next := head.Next

		head.Next = last

		last = head
		head = next
	}
	return
}
