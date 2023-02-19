package impl

type Node[T interface{}] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T interface{}] interface {
	Append(val T) (T, bool)
	Insert(val T, position int) (T, bool)
	Get(position int) (T, bool)
	Delete(position int) (T, bool)
	Traverse(callback func (val T, index int)) 
}

type linked_list[T interface{}] struct {
	head *Node[T]
}

func New[T interface{}]() LinkedList[T] {
	return &linked_list[T]{
		head: nil,
	}
}

func (ll *linked_list[T]) Traverse(callback func (T, int)) {
	index := 0
	
	for node := ll.head; node.Next != nil; node = node.Next {
		callback(node.Data, index)
		
		index++
	}
}

func (ll *linked_list[T]) Append(val T) (T, bool) {
	node := ll.head
	
	newNode := &Node[T]{Data: val, Next: nil}
	
	if ll.head == nil {
		ll.head = newNode

		return ll.head.Data, true
	}
	
	for ; node.Next != nil; {
		node = node.Next
	}

	node.Next = newNode

	return newNode.Data, true
}

func (ll *linked_list[T]) getNode(position int) (*Node[T], bool) {
	default_node := &Node[T]{}

	if ll.head == nil {
		return default_node, false
	}
	
	pos := 0

	node := ll.head

	for pos < position && node.Next != nil {
		node = node.Next

		pos++
	}

	if pos != position {
		return default_node, false	
	}
	
	return node, true
}

func (ll *linked_list[T]) Get(position int) (T, bool)  {
	var value T
	
	node, exists := ll.getNode(position)

	if !exists {
		return value, false 
	}

	return node.Data, true
}

func (ll *linked_list[T]) Insert(val T, position int) (T, bool) {
	var value T
	
	// ? Create new node
	new_node := &Node[T]{Data: val, Next: nil}
	
	// ? Insert as head if ll is empty   
	if ll.head == nil {
		ll.head = new_node

		return ll.head.Data, true
	}
	
	// ? Insert at 0
	if position == 0 {
		new_node.Next = ll.head
		ll.head = new_node

		return ll.head.Data, true
	}

	// ? Get before and current prev_node at "position" to preserve link
	prev_node, prev_exists := ll.getNode(position - 1)

	if !prev_exists {
		return value, false
	}

	next_node := prev_node.Next

	prev_node.Next = new_node

	if next_node == nil  {
		return new_node.Data, true
	}

	new_node.Next = next_node

	return new_node.Data, true
}

func (ll *linked_list[T]) Delete(position int) (T, bool) {
	var value T

	if position == 0 {
		next := ll.head.Next
		
		// ? Reset the element to be deketed
		ll.head.Data = value
		ll.head.Next = nil

		ll.head = next		

		if ll.head == nil {
			return value, false	
		}

		return ll.head.Data, true
	}
	
	prev_node, prev_exists := ll.getNode(position - 1)

	if !prev_exists {
		return value, false
	}

	next_node := prev_node.Next

	if next_node == nil {
		prev_node.Next = nil
	}

	new_next := next_node.Next

	deleted := next_node.Data

	next_node.Data = value
	next_node.Next = nil

	prev_node.Next = new_next
	
	return deleted, true
}