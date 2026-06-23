package problem

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	parent := &ListNode{}
	curr := parent
	carry := 0
	digit := 0

	for l1 != nil || l2 != nil {
		digit, carry = getDigitAndCarry(l1, l2, carry)
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		curr.Next = &ListNode{
			Val: digit,
		}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}

	return parent.Next
}

func getDigitAndCarry(
	l1 *ListNode,
	l2 *ListNode,
	existingCarry int,
) (digit int, carry int) {
	addition := existingCarry
	if l1 != nil {
		addition += l1.Val
	}
	if l2 != nil {
		addition += l2.Val
	}

	digit = addition % 10
	carry = addition / 10
	return
}
