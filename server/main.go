package main

import (
    "fmt"
    "net"
    "os"
)

const (
    HOST = "localhost"
    PORT = "59090"
)

func main() {
    ln, err := net.Listen("tcp", HOST+":"+PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer ln.Close()

    fmt.Println("Listening...")

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err.Error())
            os.Exit(1)
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    _, err := conn.Write([]byte("Hello! :)\n"))
    if err != nil {
        fmt.Println("Error sending:", err.Error())
        os.Exit(1)
    }
    
    conn.Close()
}
