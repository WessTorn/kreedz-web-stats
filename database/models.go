package database

import "time"

type User struct {
	ID      int64  `db:"id"`
	SteamID string `db:"steam_id"`
	Name    string `db:"name"`
}

type Map struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type Record struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	MapID       int64     `db:"map_id"`
	Time        int       `db:"time"`
	Date        time.Time `db:"date"`
	CP          int       `db:"cp"`
	TP          int       `db:"tp"`
	Weapon      int       `db:"weapon"`
	AA          int       `db:"aa"`
	IsProRecord bool      `db:"is_pro_record"`
}
