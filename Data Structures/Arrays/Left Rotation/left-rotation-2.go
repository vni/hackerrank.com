// https://www.hackerrank.com/challenges/array-left-rotation/problem

package main
import "fmt"

func main() {
    var n, d, i int
    fmt.Scan(&n, &d)

    var arr = make([]int, d)
    for i = 0; i < d && i < n; i++ {
        fmt.Scan(&arr[i])
    }

    for ; i < n; i++ {
        var v int
        fmt.Scan(&v)
        fmt.Printf("%v ", v)
    }

    for i = 0; i < d && i < n; i++ {
        fmt.Printf("%v ", arr[i])
    }

    fmt.Println()
}
