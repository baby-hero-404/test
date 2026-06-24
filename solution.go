package main

import "fmt"

// CountOccurrences counts the number of occurrences of an integer x in a slice of integers arr.
func CountOccurrences(arr []int, x int) int {
    count := 0
    for _, val := range arr {
        if val == x {
            count++
        }
    }
    return count
}

func main() {
    arr := []int{1, 2, 3, 2, 4, 2, 5}
    x := 2
    result := CountOccurrences(arr, x)
    fmt.Printf("Occurrences of %d: %d\n", x, result)
}
