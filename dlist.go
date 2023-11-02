package main

import "fmt"

type DNode struct {
    val int
    next *DNode
    prev *DNode
}

type DList struct {
    head *DNode
}

func new_dnode(val int) *DNode {
    n := DNode {
        val: val,
        next: nil,
        prev: nil,
    }
    return &n
}

func (l *DList) length() int {
    len := 0
    n := l.head
    for ; n != nil; len++ {
        n = n.next
    }
    return len
}

func (l *DList) insertAt(item int, idx int) {
    h := l.head
    for ; h != nil && idx > 0; idx-- {
        h = h.next
    }
    if h != nil {
        n := new_dnode(item)
        n.next = h.next
        h.next.prev = n
        n.prev = h
        h.next = n
    } else {
        // TODO
    }
}

func (l *DList) remove(item int) {
    h := l.head
    for ; h != nil; h = h.next {
        if h.val == item {
            h.prev.next = h.next
            h.next.prev = h.prev
            return
        }
    }
}

func (l *DList) removeAt(idx int) {
    n := l.head
    for ; n != nil && idx > 0; idx-- {
        n = n.next
    }
    if n != nil {
        n.prev.next = n.next
        n.next.prev = n.prev
    }
}

func (l *DList) append(item int) {
    n := new_dnode(item)
    h := l.head
    if h == nil {
        l.head = n
        return
    }

    for ; h.next != nil; h = h.next {
    }
    n.prev = h
    h.next = n
}

func (l *DList) prepend(item int) {
    n := new_dnode(item)
    if l.head == nil {
        l.head = n
        return
    }
    n.next = l.head
    l.head.prev = n
    l.head = n
}

func (l *DList) get(idx int) int {
    n := l.head
    for ; n != nil && idx > 0; idx-- {
        n = n.next
    }
    if n == nil {
        // err
        return -1
    } else {
        return n.val
    }
}

func (l *DList) print() {
    n := l.head
    for ; n != nil; n = n.next {
        fmt.Print(n.val, " -> ")
    }
    fmt.Println("nil")

}

func testDList() {
    h := new_dnode(1)
    l := DList { head: h }
    l.append(2)
    l.append(3)
    l.append(4)
    l.append(5)
    l.append(6)
    l.append(7)
    l.append(8)
    l.append(9)
    l.append(10)
    l.print()
    l.prepend(0)
    l.print()
    fmt.Println(l.get(3))
    fmt.Println(l.length())
    l.removeAt(3)
    l.print()
    l.insertAt(3, 2)
    l.print()
    l.remove(3)
    l.print()
}
