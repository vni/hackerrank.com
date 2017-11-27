// https://www.hackerrank.com/challenges/array-left-rotation/problem

// a solution with 3 reverses. Do not require additional memory
// (only O(1) for temp var)
// It is O(2*N) => O(N) solution
// (need to reverse 2 subarrays and 1 full array)

package main
import "fmt"

func reverse(arr []int) {
    n := len(arr)
    for i := 0; i < n/2; i++ {
        t := arr[i]
        arr[i] = arr[n-1-i]
        arr[n-1-i] = t
    }
}

func main() {
    var n, d int
    fmt.Scan(&n, &d)

    d = d % n

    var arr = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    reverse(arr[:d])
    reverse(arr[d:])
    reverse(arr[:])

    for _, v := range(arr) {
        fmt.Printf("%v ", v)
    }

    fmt.Println()
}
