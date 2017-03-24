package main

import (
    "log"
    "github.com/goburrow/serial"
    "fmt"
)

func main() {
    
    fmt.Println("start....")
    port, err := serial.Open(&serial.Config{Address: "/dev/ttyUSB0"})
    
    fmt.Println(port)
    fmt.Println(err)
    
    if err != nil {
        log.Fatal(err)
    }
    
    defer port.Close()
    
    _, err = port.Write([]byte("serial"))
    
    if err != nil {
        log.Fatal(err)
    }
}
