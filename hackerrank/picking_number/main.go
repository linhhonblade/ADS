package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
    // Write your code here
    var start_idx, diff, pivot, next_pivot_idx int32
    for i, v := range a {
        for j, k := range a[i+1:] {
            if a[j] != v {
                next_pivot_idx = j
            }
            diff = diff + [a[j] - a[j-1]]
            if math.Abs(diff) > 1 {
                
            }
        }
    }
}