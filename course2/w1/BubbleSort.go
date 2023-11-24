package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Swap swaps the elements at index i and i+1.
func Swap(slice []int, i int) {
    slice[i], slice[i+1] = slice[i+1], slice[i]
}

// BubbleSort sorts the slice using bubble sort algorithm.
func BubbleSort(slice []int) {
    n := len(slice)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if slice[j] > slice[j+1] {
                Swap(slice, j)
            }
        }
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter up to 10 integers, separated by space:")

    // Read a line of input
    scanner.Scan()
    input := scanner.Text()

    // Split the input into words
    words := strings.Fields(input)

    // Convert words to integers and add to the slice
    var numbers []int
    for _, word := range words {
        number, err := strconv.Atoi(word)
        if err != nil {
            fmt.Println("Error: Please enter valid integers.")
            return
        }
        numbers = append(numbers, number)
    }

    if len(numbers) > 10 {
        fmt.Println("Error: Please enter no more than 10 integers.")
        return
    }

    // Sort the numbers
    BubbleSort(numbers)

    // Print the sorted numbers
    fmt.Println("Sorted numbers:", numbers)
}
