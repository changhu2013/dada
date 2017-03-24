package main

import (
    "log"
    "github.com/goburrow/serial"
)

func main() {
    port, err := serial.Open(&serial.Config{Address: "/dev/ttyUSB0"})
    
    if err != nil {
        log.Fatal(err)
    }
    
    defer port.Close()
    
    _, err = port.Write([]byte("serial"))
    
    if err != nil {
        log.Fatal(err)
    }
}
