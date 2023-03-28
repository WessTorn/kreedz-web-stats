package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"math"
	"net/http"
)

type PlayerRec struct {
	MapID      int
	MapName    string
	Time       uint32
	FormatTime string
	Date       string
	Weapon     int
}

type TemplateDataPlayer struct {
	Players       Player
	PlayerRecords []PlayerRec
}

func GetStatusPlayer(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("frontend/player.gohtml", "frontend/head.gohtml", "frontend/header.gohtml")
		vars := mux.Vars(r)

		res, err := db.Query(fmt.Sprintf("SELECT * FROM `kz_uid` WHERE `id` = %s", vars["id"]))
		if err != nil {
			panic(err)
		}

		var player Player
		for res.Next() {
			err = res.Scan(&player.ID, &player.SteamID, &player.Name, &player.PlayTime)
			if err != nil {
				panic(err)
			}
		}

		res2, err := db.Query("SELECT kz_records.id, kz_records.map_id, kz_uid.last_name, kz_records.user_id, kz_records.time FROM kz_uid INNER JOIN kz_records ON kz_uid.id = kz_records.user_id INNER JOIN kz_maps ON kz_maps.id = kz_records.map_id WHERE kz_records.weapon = 6 AND kz_records.is_pro_record = 1 ORDER BY kz_records.user_id")
		if err != nil {
			panic(err)
		}

		var proRecords []ProRecords
		for res2.Next() {
			var proRecord ProRecords
			var Dyblicat bool
			err = res2.Scan(&proRecord.RecID, &proRecord.MapID, &proRecord.PlayerName, &proRecord.PlayerID, &proRecord.Time)
			if err != nil {
				panic(err)
			}

			for i := 0; i < len(proRecords); i++ {
				if proRecord.MapID == proRecords[i].MapID && proRecord.PlayerID == proRecords[i].PlayerID {
					if proRecord.RecID > proRecords[i].RecID {
						proRecords[i].Time = proRecord.Time
						Dyblicat = true
						continue
					}
				}
			}
			if !Dyblicat {
				proRecords = append(proRecords, proRecord)
			}
		}

		var playersInfo []PlayerInfo
		playersInfo = GetPlayerInfo(proRecords, playersInfo)

		var tops []TopsInMap
		tops = GetMapTopRecs(proRecords, tops)

		for i := 0; i < len(playersInfo); i++ {
			if playersInfo[i].ID == player.ID {
				player.Rec = playersInfo[i].Rec
				player.Records = playersInfo[i].Records
			}
		}

		player = GetTopPlayer(tops, player)

		res3, err := db.Query(fmt.Sprintf("SELECT kz_records.time, kz_records.Date, kz_records.weapon, kz_maps.mapname, kz_maps.id FROM kz_records INNER JOIN kz_maps ON kz_maps.id = kz_records.map_id WHERE kz_records.user_id = '%s' ORDER BY kz_records.Date DESC LIMIT 15 ", vars["id"]))
		if err != nil {
			panic(err)
		}

		var playerRec []PlayerRec
		for res3.Next() {
			var record PlayerRec
			err = res3.Scan(&record.Time, &record.Date, &record.Weapon, &record.MapName, &record.MapID)
			if err != nil {
				panic(err)
			}
			record.FormatTime = formatTime(math.Float32frombits(record.Time))
			playerRec = append(playerRec, record)
		}
		data := TemplateDataPlayer{
			Players:       player,
			PlayerRecords: playerRec,
		}
		err = t.ExecuteTemplate(w, "player", data)
		if err != nil {
			panic(err)
		}
	}
}
