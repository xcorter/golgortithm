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

func merge(left []int, right []int) []int {
    res := []int{}
    maxLeft := len(left)
    maxRight := len(right)
    currLeft := 0
    currRight := 0
    for (maxLeft > currLeft || maxRight > currRight) {
        if (maxLeft == currLeft) {
            res = append(res, right[currRight])
            currRight++
            continue
        }
        if (maxRight == currRight) {
            res = append(res, left[currLeft])
            currLeft++
            continue
        }

        if (left[currLeft] < right[currRight]) {
            res = append(res, left[currLeft])
            currLeft++
        } else {
            res = append(res, right[currRight])
            currRight++
        }
    }
    return res
}

func sort(arr []int) []int {
    length := len(arr)
    if (length < 2) {
        return arr
    }

    middle := length / 2
    left := arr[:middle]
    right := arr[middle:]

    left = sort(left)
    right = sort(right)
    return merge(left, right)
}
