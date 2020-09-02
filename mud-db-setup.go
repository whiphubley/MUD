package main

import (
    "database/sql"
    "fmt"
    "strconv"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    database, _ := sql.Open("sqlite3", "./mud-database.db")

    create_users, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, score INTEGER, room INTEGER, weapon INTEGER)")
    create_users.Exec()
    create_rank, _ := database.Prepare("CREATE TABLE IF NOT EXISTS rank (id INTEGER PRIMARY KEY, score INTEGER, title TEXT)")
    create_rank.Exec()
    create_room, _ := database.Prepare("CREATE TABLE IF NOT EXISTS room (id INTEGER PRIMARY KEY, desc TEXT)")
    create_room.Exec()
    create_weapon, _ := database.Prepare("CREATE TABLE IF NOT EXISTS weapon (id INTEGER PRIMARY KEY, desc TEXT)")
    create_weapon.Exec()

    insert_users, _ := database.Prepare("INSERT INTO users (username, score, room, weapon) VALUES (?, ?, ?, ?)")
    insert_users.Exec("testuser", 0, 1, 1)

    insert_rank, _ := database.Prepare("INSERT INTO rank (score, title) VALUES (?, ?)")
    insert_rank.Exec(0, "SERF")
    insert_rank.Exec(100, "PEASANT")

    insert_room, _ := database.Prepare("INSERT INTO room (desc) VALUES (?)")
    insert_room.Exec("Welcome to room 01")
    insert_room.Exec("Welcome to room 02")

    insert_weapon, _ := database.Prepare("INSERT INTO weapon (desc) VALUES (?)")
    insert_weapon.Exec("Feather Duster")
    insert_weapon.Exec("Pen")
}
