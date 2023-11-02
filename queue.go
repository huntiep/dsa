package main

type Queue struct {
    head *Node
    tail *Node
}

func (q *Queue) enqueue(val int) {
    n := new_node(val)
    q.tail.next = n
    q.tail = n
}

func (q *Queue) dequeue() int {
    if q.head != nil {
        v := q.head.val
        q.head = q.head.next
        return v
    } else {
        return -1
    }
}
