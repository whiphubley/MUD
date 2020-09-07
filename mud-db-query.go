package main

import (
    "database/sql"
    "fmt"
    "strconv"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    database, _ := sql.Open("sqlite3", "./mud-database.db")

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
    rows_room, _ := database.Query("SELECT id, desc, n, e, s, w FROM room")
    var id_room int
    var desc_room string
    var n_room int
    var e_room int
    var s_room int
    var w_room int
    for rows_room.Next() {
        rows_room.Scan(&id_room, &desc_room, &n_room, &e_room, &s_room, &w_room)
        fmt.Println(strconv.Itoa(id_room) + ": " + desc_room + " " + strconv.Itoa(n_room) + " " + strconv.Itoa(e_room) + " " + strconv.Itoa(s_room) + " " + strconv.Itoa(w_room))
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
