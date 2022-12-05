package types

type element[T any] struct {
	data T
	next *element[T]
}

type Stack[T any] struct {
	head *element[T]
	Size int
}

func (stk *Stack[T]) Push(data T) {
	element := new(element[T])
	element.data = data
	temp := stk.head
	element.next = temp
	stk.head = element
	stk.Size++
}

// get data of the element on the top of the stack without popping
func (stk *Stack[T]) Peek() T {
	return stk.head.data
}

func (stk *Stack[T]) Pop() T {
	if stk.head == nil {
		return *new(T) // return zero value
	}
	r := stk.head.data
	stk.head = stk.head.next
	stk.Size--

	return r
}

// get values of stack as list
func (stk *Stack[T]) Values() []T {
	if stk.head == nil {
		return []T{}
	}

	return stk.head.getUpTo()
}

// Get all element data below this one, including the element's data itself
func (elem *element[T]) getUpTo() []T {
	if elem.next != nil {
		return append([]T{elem.data}, elem.next.getUpTo()...) // add to front
	} else {
		return []T{elem.data}
	}
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Size: 0, head: nil}
}
