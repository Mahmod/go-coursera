package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Animal struct {
    food       string
    locomotion string
    noise      string
}

func (a Animal) Eat() {
    fmt.Println(a.food)
}

func (a Animal) Move() {
    fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
    fmt.Println(a.noise)
}

func main() {
    animals := map[string]Animal{
        "cow":   {food: "grass", locomotion: "walk", noise: "moo"},
        "bird":  {food: "worms", locomotion: "fly", noise: "peep"},
        "snake": {food: "mice", locomotion: "slither", noise: "hsss"},
    }

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("> ")
        scanner.Scan()
        userInput := scanner.Text()
        inputParts := strings.Split(userInput, " ")

        if len(inputParts) != 2 {
            fmt.Println("Invalid input. Please enter two strings: an animal name and an action.")
            continue
        }

        animalName, action := inputParts[0], inputParts[1]
        animal, ok := animals[animalName]

        if !ok {
            fmt.Printf("No information available for '%s'\n", animalName)
            continue
        }

        switch action {
        case "eat":
            animal.Eat()
        case "move":
            animal.Move()
        case "speak":
            animal.Speak()
        default:
            fmt.Printf("Invalid action '%s'. Valid actions are 'eat', 'move', 'speak'.\n", action)
        }
    }
}
