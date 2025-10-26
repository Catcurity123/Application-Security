package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("test.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()  // Ensure file closes when function exits

    n, err := file.Write([]byte("Go is great!"))
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
    
    fmt.Println(n, "bytes written")
}

