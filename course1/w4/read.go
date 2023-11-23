package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Name struct {
    fname string
    lname string
}

func main() {
    var fileName string
    var names []Name

    fmt.Print("Enter the name of the text file: ")
    fmt.Scan(&fileName)

    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)

        var firstName, lastName string
        if len(parts) > 0 {
            firstName = parts[0]
            if len(firstName) > 20 {
                firstName = firstName[:20]
            }
        }
        if len(parts) > 1 {
            lastName = parts[1]
            if len(lastName) > 20 {
                lastName = lastName[:20]
            }
        }

        names = append(names, Name{fname: firstName, lname: lastName})
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }

    for _, name := range names {
        fmt.Printf("First Name: %-20s Last Name: %-20s\n", name.fname, name.lname)
    }
}
