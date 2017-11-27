// https://www.hackerrank.com/challenges/array-left-rotation/problem

// a solution with a true left rotation in a loop

package main
import "fmt"

func main() {
    var n, d int
    fmt.Scan(&n, &d)

    d = d % n

    var arr = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    for ; d > 0; d-- {
        t := arr[0]
        for i := 0; i < n-1; i++ {
            arr[i] = arr[i+1]
        }
        arr[n-1] = t
    }

    for _, v := range(arr) {
        fmt.Printf("%v ", v)
    }

    fmt.Println()
}
