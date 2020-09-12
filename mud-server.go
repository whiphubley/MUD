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

var database, _ = sql.Open("sqlite3", "./mud-database.db")
var m = make(map[string]net.Conn)

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
	// tell other players in the room you have entered
	rows_users, err := database.Query("SELECT username FROM users WHERE room = ? AND username != ?", room, username)
	// handle an empty set
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("EMPTY SET")
		} else {
			fmt.Println(err)
			return
		}
		// if the user from the db is connected then send them a message
	} else {
		var user string
		for rows_users.Next() {
			rows_users.Scan(&user)
			if conn, ok := m[user]; ok {
				conn.Write([]byte(string(username + " has entered the room\n# ")))
			}
		}
	}
	// update the room you are in
	_, err = database.Exec("UPDATE users SET room = ? WHERE username = ?", room, username)
	if err != nil {
		fmt.Println(err)
		return
	}
	// show the room description
	var desc string
	database.QueryRow("SELECT desc FROM room WHERE id = ?", room).Scan(&desc)
	c.Write([]byte(string(desc + "\n" + "# ")))
	// handle user input
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
		if text == "" {
			c.Write([]byte(string("You need to enter a command...\n" + "# ")))
		} else if text == "n" || text == "e" || text == "s" || text == "w" {
			query := fmt.Sprintf("SELECT %v FROM room WHERE id = ?", text)
			var result int
			err = database.QueryRow(query, room).Scan(&result)
			if err != nil {
				log.Fatal(err)
			}
			enterRoom(c, username, result)
		} else if text == "QUIT" {
			c.Write([]byte(string("Thanks for playing " + username + " !!\nEscape character is '^]'\n")))
			break
		} else {
			c.Write([]byte(string("COMMANDS:\n\n" + "n: MOVE NORTH\n" + "e: MOVE EAST\n" + "s: MOVE SOUTH\n" + "w: MOVE WEST\n" + "\n# ")))
		}
	}
}

func handleConnection(c net.Conn) {
	// choose to enter the game
	askQuestion(c, "Welcome to FlexMUD dare you enter...(y/n): ", "Goodbye weakling.\n", "Good luck !!\n")
	// get user details
	username := createUser(c, "Please enter you username (new users will be created / existing users will be loaded): ")
	// map username to connection
	m[username] = c
	// enter the map at last location
	var room int
	database.QueryRow("SELECT room FROM users WHERE username = ?", username).Scan(&room)
	enterRoom(c, username, room)
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
