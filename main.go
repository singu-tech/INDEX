package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    if os.Getenv("INDEX_DEBUG") == "TRUE" {
        log.Println("[DEBUG] Enable `DEBUG` mod.")
    }
    fmt.Println("Hello World!")
}
