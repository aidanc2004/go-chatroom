package main

import (
    "fmt"
    "net"
    "os"
)

const (
    host = "localhost"
    port = "59091"
)

func main() {
    ln, err := net.Listen("tcp", host+":"+port)
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
    name := make([]byte, 20)

    conn.Write([]byte("Please enter your name: "))
    
    conn.Read(name)

    conn.Write([]byte("Hello, " + string(name)))

    conn.Close()
}
