package main

import (
    "fmt"
)

func main() {
    fmt.Println("Quick sort!")
    a := []int{5,8,5,7,8,9,1,23,65,2}
    fmt.Printf("%v\n", a)
    result := sort(a)
    fmt.Printf("%v", result)
}

func sort(arr []int) []int {
    for j := 1; j < len(arr); j++ {
        key := arr[j]
        i := j - 1
        for (i >= 0 && arr[i] > key) {
            arr[i +1] = arr[i]
            i = i - 1
        }
        arr[i + 1] = key
    }
    return arr
}