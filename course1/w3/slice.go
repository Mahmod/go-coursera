package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    slice := make([]int, 0, 3) // Slice with length 0 and capacity 3

    for {
        fmt.Print("Enter an integer or 'X' to exit: ")
        scanner.Scan()
        input := scanner.Text()

        // Check if the user wants to exit
        if strings.ToUpper(input) == "X" {
            break
        }

        // Convert the input to an integer
        num, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid input. Please enter an integer.")
            continue
        }

        // Add the integer to the slice
        slice = append(slice, num)

        // Sort the slice
        sort.Ints(slice)

        // Print the sorted slice
        fmt.Println("Sorted slice:", slice)
    }
}
