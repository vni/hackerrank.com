// https://www.hackerrank.com/challenges/2d-array/problem

package main
import "fmt"

func main() {
    var ar [6][6]int8

    for r := 0; r < 6; r++ {
        for c := 0; c < 6; c++ {
            fmt.Scan(&ar[r][c])
        }
    }

    var max_sum int8 = -128 // -9 * 7 = -63 -> the minimum, which we can get, so -100 will be less in any case
    for r := 0; r < 4; r++ {
        for c := 0; c < 4; c++ {
            local_sum := ar[r][c] + ar[r][c+1] + ar[r][c+2] +
                         ar[r+1][c+1] +
                         ar[r+2][c] + ar[r+2][c+1] + ar[r+2][c+2]

            if local_sum > max_sum {
                max_sum = local_sum
            }
        }
    }

    fmt.Println(max_sum)
}
