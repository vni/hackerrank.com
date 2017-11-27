// https://www.hackerrank.com/challenges/array-left-rotation/problem

// a solution with a true array and a true rotation

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

    var temp = make([]int, d)
    for i := 0; i < d; i++ {
        temp[i] = arr[i]
    }

    for i := 0; i < n-d; i++ {
        arr[i] = arr[i+d]
    }

    for i, j := n-d, 0; i < n; i++ {
        arr[i] = temp[j]
        j++
    }

    for _, v := range(arr) {
        fmt.Printf("%v ", v)
    }

    fmt.Println()
}
