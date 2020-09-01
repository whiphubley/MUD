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

    fmt.Println("USERS:")
    rows_users, _ := database.Query("SELECT id, username, score, room, weapon FROM users")
    var id_users int
    var username_users string
    var score_users int
    var room_users int
    var weapon_users int
    for rows_users.Next() {
        rows_users.Scan(&id_users, &username_users, &score_users, &room_users, &weapon_users)
        fmt.Println(strconv.Itoa(id_users) + ": " + username_users + " " + strconv.Itoa(score_users) + " " + strconv.Itoa(room_users) + " " + strconv.Itoa(weapon_users))
    }

    fmt.Println("RANK:")
    rows_rank, _ := database.Query("SELECT id, score, title FROM rank")
    var id_rank int
    var score_rank int
    var title_rank string
    for rows_rank.Next() {
        rows_rank.Scan(&id_rank, &score_rank, &title_rank)
        fmt.Println(strconv.Itoa(id_rank) + ": " + strconv.Itoa(score_rank) + " " + title_rank)
    }

    fmt.Println("ROOM:")
    rows_room, _ := database.Query("SELECT id, desc FROM room")
    var id_room int
    var desc_room string
    for rows_room.Next() {
        rows_room.Scan(&id_room, &desc_room)
        fmt.Println(strconv.Itoa(id_room) + ": " + desc_room)
    }

    fmt.Println("WEAPON:")
    rows_weapon, _ := database.Query("SELECT id, desc FROM weapon")
    var id_weapon int
    var desc_weapon string
    for rows_weapon.Next() {
        rows_weapon.Scan(&id_weapon, &desc_weapon)
        fmt.Println(strconv.Itoa(id_weapon) + ": " + desc_weapon)
    }
}
