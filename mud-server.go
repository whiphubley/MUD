package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"net"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
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

func createUser(c net.Conn, q string) {
	c.Write([]byte(string(q)))
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		username := strings.TrimSpace(string(netData))

		if username == "" {
			c.Write([]byte(string("You need to enter a username: ")))
		} else {
			database, _ := sql.Open("sqlite3", "./mud-database.db")
			_, err := database.Exec("INSERT OR FAIL INTO users (username, score, room, weapon) VALUES (?, ?, ?, ?)", username, 0, 1, 1)
			if err != nil {
				c.Write([]byte(string("Welcome back " + username + "\n")))
				return
			}
			c.Write([]byte(string("Welcome " + username + "\n")))
			return
		}
	}
	c.Close()
}

func handleConnection(c net.Conn) {
	// choose to enter the game
	askQuestion(c, "Welcome to FlexMUD dare you enter...(y/n): ", "Goodbye weakling.\n", "Good luck !!\n")
	// get user details
	createUser(c, "Please enter you username (new users will be created / existing users will be loaded): ")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text := strings.TrimSpace(string(netData))
		if text == "QUIT" {
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
