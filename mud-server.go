package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func handleConnection(c net.Conn) {
	c.Write([]byte(string("Welcome to AndyMUD dare you enter...\n")))
        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }

                text := strings.TrimSpace(string(netData))
                if text == "STOP" {
                        break
                }
		c.Write([]byte(string(text + " you are an arse\n")))
        }
        c.Close()
}

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a port number!")
                return
        }

        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp4", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()

        for {
                c, err := l.Accept()
                if err != nil {
                        fmt.Println(err)
                        return
                }
                go handleConnection(c)
        }
}
