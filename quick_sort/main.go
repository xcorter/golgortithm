package main

import (
    "fmt"
)

func main() {
    fmt.Println("Quick sort!")
    a := []int{5,8,5,7,8,9,1}
    fmt.Printf("%v\n", a)
    result := sort(a)
    fmt.Printf("%v", result)
}

func sort(arr []int) []int {
    length := len(arr)
    if (length < 2) {
        return arr
    }
    pivotKey := length / 2
    pivot := arr[pivotKey]
    var left []int
    var right []int
    var equals []int

    for _, value := range arr {
        if (value < pivot) {
            left = append(left, value)
        }
        if (value > pivot) {
            right = append(right, value)
        }
        if (value == pivot) {
            equals = append(equals, value)
        }
    }
    left = sort(left)
    right = sort(right)
    left =  append(left, equals...)
    return append(left, right...)
}