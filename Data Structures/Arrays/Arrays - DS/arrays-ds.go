package main
import "fmt"

func main() {
    N := 0
    fmt.Scan(&N)

    arr := make([]int, N)
    for i,_ := range(arr) {
        fmt.Scan(&arr[i])
    }

    for i := len(arr)-1; i >= 0; i-- {
        fmt.Printf("%v ", arr[i]);
    }
}
