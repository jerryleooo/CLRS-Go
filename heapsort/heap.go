package main

import (
    "fmt"
)

func parent(i int) int {
    return (i - 1) / 2
}

func left(i int) int {
    return 2 * i + 1
}

func right(i int) int {
    return 2 * i + 2
}

func maxHeapify(A []int, n int, i int) {
    l := left(i)
    r := right(i)
    largest := i

    if l < n && A[l] > A[i] {
        largest = l
    }
    if r < n && A[r] > A[largest] {
        largest = r
    }

    if largest != i {
        A[i], A[largest] = A[largest], A[i]
        maxHeapify(A, n, largest)
    }
}

func buildMaxHeap(A []int, n int) {
    for i := n / 2 - 1; i >= 0; i -- {
        maxHeapify(A, n, i)
    }
}

func heapSort(A []int) {
    n := len(A)
    buildMaxHeap(A, n)
    for i := n - 1; i > 0; i -- {
        A[i], A[0] = A[0], A[i]
        n -= 1
        maxHeapify(A, n, 0)
    }
}

func main() {
    A := []int{1, 7, 5, 6, 2, 3, 0}
    fmt.Println("before heapify", A)
    maxHeapify(A, len(A), 0)
    fmt.Println("after heapify", A)

    B := []int{1, 7, 5, 6, 2, 3, 0}

    fmt.Println("before buildMaxHeap", B)
    buildMaxHeap(B, len(B))
    fmt.Println("after buildMaxHeap", B)

    C := []int{1, 7, 5, 6, 2, 3, 0}
    fmt.Println("before heapSort", C)
    heapSort(C)
    fmt.Println("after heapSort", C)
}
