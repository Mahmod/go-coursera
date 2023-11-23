package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a string: ")
    input, _ := reader.ReadString('\n')

    // Convert the input to lowercase for case-insensitive comparison
    inputLower := strings.ToLower(input)

    // Trim the newline character from the input
    inputLower = strings.TrimSpace(inputLower)

    if strings.HasPrefix(inputLower, "i") && strings.Contains(inputLower, "a") && strings.HasSuffix(inputLower, "n") {
        fmt.Println("Found!")
    } else {
        fmt.Println("Not Found!")
    }
}
