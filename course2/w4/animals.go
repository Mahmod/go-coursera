package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Animal interface
type Animal interface {
    Eat()
    Move()
    Speak()
}

// Cow type
type Cow struct{}

func (c Cow) Eat() {
    fmt.Println("grass")
}

func (c Cow) Move() {
    fmt.Println("walk")
}

func (c Cow) Speak() {
    fmt.Println("moo")
}

// Bird type
type Bird struct{}

func (b Bird) Eat() {
    fmt.Println("worms")
}

func (b Bird) Move() {
    fmt.Println("fly")
}

func (b Bird) Speak() {
    fmt.Println("peep")
}

// Snake type
type Snake struct{}

func (s Snake) Eat() {
    fmt.Println("mice")
}

func (s Snake) Move() {
    fmt.Println("slither")
}

func (s Snake) Speak() {
    fmt.Println("hsss")
}

func main() {
    animals := make(map[string]Animal)
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("> ")
        scanner.Scan()
        userInput := scanner.Text()
        commands := strings.Fields(userInput)

        if len(commands) != 3 {
            fmt.Println("Invalid command. Please enter a command with exactly three arguments.")
            continue
        }

        switch commands[0] {
        case "newanimal":
            animalName, animalType := commands[1], commands[2]
            switch animalType {
            case "cow":
                animals[animalName] = Cow{}
            case "bird":
                animals[animalName] = Bird{}
            case "snake":
                animals[animalName] = Snake{}
            default:
                fmt.Println("Invalid animal type. Please choose from 'cow', 'bird', or 'snake'.")
                continue
            }
            fmt.Println("Created it!")

        case "query":
            animalName, infoType := commands[1], commands[2]
            animal, ok := animals[animalName]
            if !ok {
                fmt.Printf("No animal with name '%s' found.\n", animalName)
                continue
            }

            switch infoType {
            case "eat":
                animal.Eat()
            case "move":
                animal.Move()
            case "speak":
                animal.Speak()
            default:
                fmt.Println("Invalid query type. Please choose from 'eat', 'move', or 'speak'.")
            }

        default:
            fmt.Println("Invalid command. Please start with 'newanimal' or 'query'.")
        }
    }
}
