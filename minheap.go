package main

type MinHeap struct {
    heap []int
}

func (h *MinHeap) insert(val int) {
    h.heap = append(h.heap, val)
    i := len(h.heap) - 1
    p := (i-1)/2
    for p >=0 {
        if h.heap[p] > h.heap[i] {
            t := h.heap[p]
            h.heap[p] = h.heap[i]
            h.heap[i] = t
            i = p
            p = (i-1)/2
        } else {
            return
        }
    }
}

func (h *MinHeap) delete() int {
    v := h.heap[len(h.heap)-1]

    return v
}
