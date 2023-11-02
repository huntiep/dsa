package main

import "fmt"

type BT struct {
    val int
    left *BT
    right *BT
}

func cmpBT(a, b *BT) bool {
    if a == nil || b == nil {
        if a == b {
            return true
        } else {
            return false
        }
    }
    if a.val != b.val {
        return false
    }
    return cmpBT(a.left, b.left) && cmpBT(a.right, b.right)
}

func testBT() {
    a := &BT { val: 1, left: nil, right: nil }
    a.left = &BT { val: 2, left: nil, right: nil }
    a.left.left = &BT { val: 3, left: nil, right: nil }
    a.left.right = &BT { val: 4, left: nil, right: nil }
    a.right = &BT { val: 5, left: nil, right: nil }
    a.right.left = &BT { val: 6, left: nil, right: nil }
    a.right.right = &BT { val: 7, left: nil, right: nil }

    b := &BT { val: 1, left: nil, right: nil }
    b.left = &BT { val: 2, left: nil, right: nil }
    b.left.left = &BT { val: 3, left: nil, right: nil }
    b.left.right = &BT { val: 4, left: nil, right: nil }
    b.right = &BT { val: 5, left: nil, right: nil }
    b.right.left = &BT { val: 6, left: nil, right: nil }
    b.right.right = &BT { val: 7, left: nil, right: nil }
    fmt.Println(cmpBT(a, b))

    b.right.left.val = 10
    fmt.Println(cmpBT(a, b))

    a = &BT { val: 1, left: nil, right: nil }
    a.left = &BT { val: 2, left: nil, right: nil }
    a.left.left = &BT { val: 3, left: nil, right: nil }
    b = &BT { val: 1, left: nil, right: nil }
    b.left = &BT { val: 2, left: nil, right: nil }
    b.right = &BT { val: 3, left: nil, right: nil }
    fmt.Println(cmpBT(a, b))
}
