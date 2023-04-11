package main

import (
    "fmt"
)
func main() {


    msg := "1234567"
    data := []rune(msg)

    for i := 0; i < len(data); i++ {
        if i % 2 == 1 {
            continue
        }
    
        if i > 8 {
            break
        }
        fmt.Printf("Char %c Unicode: %U, starts at byte position: %d\n", data[i], data[i], i)
    }

    fmt.Println()
}