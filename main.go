package main

import (
    "fmt"
    "github.com/eiannone/keyboard"
)

func main() {
    fmt.Println("Hello, welcome to the Maze")
    keysEvents, err := keyboard.GetKeys(10)
    if err!=nil {
        panic(err)
    }

    defer func() {
        _ = keyboard.Close();
    }()

    fmt.Println("Press ESC key to quit")
    for {
        event := <- keysEvents
        if event.Err!=nil{
            panic(event.Err)
        }
        fmt.Printf("You pressed rune: %q, key %X\n", event.Rune, event.Key)

        if event.Key == keyboard.KeyEsc{
            break
        }

    }
}
