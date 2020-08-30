package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func askQuestion(c net.Conn, q string, n string, y string) {
        c.Write([]byte(string(q)))
        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }

                text := strings.TrimSpace(string(netData))
                if text == "n" {
			c.Write([]byte(string(n)))
                        break
                }
                c.Write([]byte(string(y)))
		return
        }
        c.Close()
}

func handleConnection(c net.Conn) {
	askQuestion(c, "Welcome to FlexMUD dare you enter...\n", "Goodbye weakling.\n", "Good luck peasant !!\n")
        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }

                text := strings.TrimSpace(string(netData))
                if text == "STOP" {
			c.Write([]byte(string("Thanks for playing !!\n")))
                        break
                }
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
