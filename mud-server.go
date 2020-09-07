package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	//"strconv"
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

func createUser(c net.Conn, q string) string {
	c.Write([]byte(string(q)))
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return netData
		}

		username := strings.TrimSpace(string(netData))

		if username == "" {
			c.Write([]byte(string("You need to enter a username: ")))
		} else {
			database, _ := sql.Open("sqlite3", "./mud-database.db")
			_, err := database.Exec("INSERT OR FAIL INTO users (username, score, room, weapon) VALUES (?, ?, ?, ?)", username, 0, 1, 1)
			if err != nil {
				c.Write([]byte(string("Welcome back " + username + "\n")))
				return username
			}
			c.Write([]byte(string("Welcome " + username + "\n")))
			return username
		}
	}
	c.Close()
	return string("testuser")
}

func enterRoom(c net.Conn, username string, room int) {
	database, _ := sql.Open("sqlite3", "./mud-database.db")
	_, err := database.Exec("UPDATE users SET room = ? WHERE username = ?", room, username)
	if err !=nil {
		fmt.Println(err)
		return
	}
	var desc string
	database.QueryRow("SELECT desc FROM room WHERE id = ?", room).Scan(&desc)
	c.Write([]byte(string(desc + "\n" + "# ")))
	handleCommands(c, username, room)
}

func handleCommands(c net.Conn, username string, room int) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text := strings.TrimSpace(string(netData))
		database, _ := sql.Open("sqlite3", "./mud-database.db")
		var result int
		if text == "" {
			c.Write([]byte(string("You need to enter a command: ")))
		} else if text == "n" {
			err = database.QueryRow("SELECT n FROM room WHERE id = ?", room).Scan(&result)
			if err != nil {
				log.Fatal(err)
			}
			enterRoom(c, username, result)
		} else if text == "e" {
			err = database.QueryRow("SELECT e FROM room WHERE id = ?", room).Scan(&result)
			if err != nil {
				log.Fatal(err)
			}
			enterRoom(c, username, result)
		} else if text == "s" {
			err = database.QueryRow("SELECT s FROM room WHERE id = ?", room).Scan(&result)
			if err != nil {
				log.Fatal(err)
			}
			enterRoom(c, username, result)
		} else if text == "w" {
			err = database.QueryRow("SELECT w FROM room WHERE id = ?", room).Scan(&result)
			if err != nil {
				log.Fatal(err)
			}
			enterRoom(c, username, result)
		}
	}
}

func handleConnection(c net.Conn) {
	// choose to enter the game
	askQuestion(c, "Welcome to FlexMUD dare you enter...(y/n): ", "Goodbye weakling.\n", "Good luck !!\n")
	// get user details
	username := createUser(c, "Please enter you username (new users will be created / existing users will be loaded): ")
	// enter the map
	enterRoom(c, username, 1)
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
