package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
)

const (
    host = "localhost"
    port = "59090"
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
    conn.Write([]byte("Please enter your name: "))
    
    name, _ := bufio.NewReader(conn).ReadString('\n') 

    name = strings.TrimSuffix(name, "\r\n")
    
    conn.Write([]byte("Hello, " + name + ", how are you?\n"))

    conn.Close()
}
