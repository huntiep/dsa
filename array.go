package main

import "fmt"

func linear_search(arr []int, val int) int {
    for i := 0; i < len(arr); i++ {
        if arr[i] == val {
            return i
        }
    }
    return -1
}

func binary_search(arr []int, val int) int {
    low := 0
    high := len(arr)
    for low < high {
        mid := low + (high-low) / 2
        if arr[mid] == val {
            return mid
        } else if arr[mid] > val {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return -1
}

func bubblesort(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := 1; j < len(arr) - i; j++ {
            if arr[j-1] > arr[j] {
                t := arr[j]
                arr[j] = arr[j-1]
                arr[j-1] = t
            }
        }
    }
}

func quicksort(arr []int) {
    qs(arr, 0, len(arr) - 1)
}

func qs(arr []int, low, high int) {
    if low >= high {
        return
    }

    pivot := partition(arr, low, high)
    qs(arr, low, pivot-1)
    qs(arr, pivot+1, high)
}

func partition(arr []int, low, high int) int {
    pivot := arr[high]
    idx := low - 1

    for i := low; i < high; i++ {
        if arr[i] <= pivot {
            idx++
            tmp := arr[i]
            arr[i] = arr[idx]
            arr[idx] = tmp
        }
    }
    idx++
    arr[high] = arr[idx]
    arr[idx] = pivot
    return idx
}

func testArray() {
    var arr [10]int
    arr = [10]int{5, 3, 7, 6, 1, 8, 4, 10, 2, 9}
    fmt.Println(linear_search(arr[:], 5))
    fmt.Println(linear_search(arr[:], 6))
    fmt.Println(linear_search(arr[:], 11))
    quicksort(arr[:])
    //bubblesort(arr[:])
    fmt.Println(arr)
    fmt.Println(binary_search(arr[:], 3))
    fmt.Println(binary_search(arr[:], 2))
    fmt.Println(binary_search(arr[:], 11))
}
