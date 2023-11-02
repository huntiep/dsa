package main

type Stack struct {
    head *Node
}

func (s *Stack) push(val int) {
    n := new_node(val)
    n.next = s.head
    s.head = n
}

func (s *Stack) pop() int {
    v := s.head.val
    s.head = s.head.next
    return v
}
