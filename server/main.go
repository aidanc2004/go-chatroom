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
    name string
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
        
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    // get name from client
    conn.Write([]byte("\nPlease enter your name: "))
    
    name, _ := bufio.NewReader(conn).ReadString('\n')
    name = strings.TrimSuffix(name, "\r\n")
    
    cl := client {conn, name} // create new client struct for connection
    
    clients = append(clients, cl) // add client to slice of all clients
    
    recieveMessages(cl)

    // Close connection and remove client from slice
    conn.Close()
    removeClient(cl)
}

// recieve messages from client
func recieveMessages(cl client) {
    reader := bufio.NewReader(cl.conn)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            // TODO: handle disconnection
            break
        }

        broadcast(cl, msg)
    }
}

// send message to all clients
func broadcast(cl client, msg string) {
    for _, v := range clients {
        if v != cl {
            v.conn.Write( []byte(cl.name + ": " + msg) )
        }
    }
}

// remove a client from clients slice
func removeClient(cl client) {
    for i, v := range clients {
        if v == cl {
	        clients = append(clients[:i], clients[i+1:]...)
        }
    }
}
