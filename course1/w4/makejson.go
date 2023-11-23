package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    userInfo := make(map[string]string)

    fmt.Print("Enter name: ")
    scanner.Scan()
    userInfo["name"] = scanner.Text()

    fmt.Print("Enter address: ")
    scanner.Scan()
    userInfo["address"] = scanner.Text()

    // Marshal the map into a JSON object
    jsonInfo, err := json.Marshal(userInfo)
    if err != nil {
        fmt.Println("Error creating JSON object:", err)
        return
    }

    // Print the JSON object
    fmt.Println("JSON object:", string(jsonInfo))
}
