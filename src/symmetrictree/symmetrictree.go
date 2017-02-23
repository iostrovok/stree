package symmetrictree

type Tree struct {
	ID int

	Left, Right *Tree
}

func (this *Tree) Add(left, right *Tree) {
	this.Left = left
	this.Right = right
}

// Check full tree
func (this *Tree) CheckSimmtry() bool {

	// Check empty tree
	if this.Left == nil && this.Right == nil {
		return true
	}
	if this.Left == nil || this.Right == nil {
		return false
	}

	left := this.CheckL([]int{})
	right := this.CheckR([]int{})

	if len(left) != len(right) {
		return false
	}

	for i := range right {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

// Gets left leafs
func (this *Tree) CheckL(stack []int) []int {

	if this.Left != nil {
		stack = append(stack, this.Left.ID)
		stack = this.Left.CheckL(stack)
	}

	if this.Right != nil {
		stack = this.Right.CheckL(stack)
	}

	return stack
}

// Gets right leafs
func (this *Tree) CheckR(stack []int) []int {

	if this.Right != nil {
		stack = append(stack, this.Right.ID)
		stack = this.Right.CheckR(stack)
	}

	if this.Left != nil {
		stack = this.Left.CheckR(stack)
	}

	return stack
}
