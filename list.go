package main

type Node struct {
    val int
    next *Node
}

type List struct {
    head *Node
}

func new_node(val int) *Node {
    n := Node {
        val: val,
        next: nil,
    }
    return &n
}

func (l *List) delete(val int) {
    h := l.head
    if h == nil {
        return
    } else if h.val == val {
        l.head = l.head.next
        return
    }

    for h.next != nil {
        if h.next.val == val {
            h.next = h.next.next
            return
        }
    }
}

func (l *List) len() int {
    i := 0
    h := l.head
    for h != nil {
        i += 1
        h = h.next
    }
    return i
}

func (l *List) insert_at(val int, index int) {
    h := l.head
    for ; index > 0; index-- {
        if h == nil {
            // err
            return
        }
        h = h.next
    }
    n := new_node(val)
    n.next = h.next
    h.next = n
}

func (l *List) remove_at(index int) int {
    h := l.head
    if h == nil {
        // err
        return -1
    } else if index == 0 {
        v := h.val
        l.head = l.head.next
        return v
    }

    for ; index > 1; index-- {
        if h.next == nil {
            // err
            return -1
        }
        h = h.next
    }
    v := h.next.val
    h.next = h.next.next
    return v
}

