// https://www.hackerrank.com/challenges/camelcase/problem

package main
import (
    "fmt"
    "unicode"
)

func main() {
    var s string
    fmt.Scan(&s)

    var words = 0

    for _, v := range(s) {
        if unicode.IsUpper(v) {
            words++
        }
    }

    if len(s) > 0 {
        words++
    }
    fmt.Println(words)
}
