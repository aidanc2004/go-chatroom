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

type client struct {
    conn net.Conn
    nick string
}

var clients = make([]client, 0) // slice of all connected clients

func main() {
    // start listening for connections
    ln, err := net.Listen("tcp", host+":"+port)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer ln.Close()

    fmt.Println("Listening...")

    for {
        // accept a connection
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err.Error())
            os.Exit(1)
        }
        
        cl := client {conn, "nick"} // create new client struct for connection

        clients = append(clients, cl) // add client to slice of all clients
        
        go handleConnection(cl)
    }
}

func handleConnection(cl client) {
    // get name from client
    cl.conn.Write([]byte("Please enter your name: "))
    
    name, _ := bufio.NewReader(cl.conn).ReadString('\n')

    name = strings.TrimSuffix(name, "\r\n")
    
    cl.conn.Write([]byte("Hello, " + name + ", how are you?\n"))
    
    // TODO: update name in array

    cl.conn.Close()

    removeClient(cl)
}

func removeClient(cl client) {
    for i, v := range clients {
        if v == cl {
	        clients = append(clients[:i], clients[i+1:]...)
        }
    }
}
