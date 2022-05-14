package src

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return true }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 1 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }

type NestedInteger struct {
}

type NestedIterator struct {
	stack []*NestedInteger
}

// problem:flatten-nested-list-iterator
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	size := len(nestedList)
	stack := make([]*NestedInteger, size)
	for ptr := 0; ptr < size; ptr++ {
		stack[ptr] = nestedList[size-ptr-1]
	}
	return &NestedIterator{
		stack: stack,
	}
}

func (this *NestedIterator) Next() int {
	size := len(this.stack)
	cur := this.stack[size-1]
	this.stack = this.stack[:size-1]
	return cur.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	this.findCloseInt()
	return len(this.stack) > 0
}

func (this *NestedIterator) findCloseInt() {
	for {
		size := len(this.stack)
		if size == 0 {
			return
		}
		cur := this.stack[size-1]
		if cur.IsInteger() {
			return
		}
		this.stack = this.stack[:size-1]
		list := cur.GetList()
		length := len(list)
		for i := length - 1; i >= 0; i-- {
			this.stack = append(this.stack, list[i])
		}
	}
}
