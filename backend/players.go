package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

func GetPlayersTop(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("frontend/players.gohtml", "frontend/head.gohtml", "frontend/header.gohtml")

		//res, err := db.Query("SELECT DISTINCT kz_uid.id, kz_uid.last_name FROM kz_uid INNER JOIN kz_records ON kz_uid.id = kz_records.user_id")
		res, err := db.Query("SELECT kz_records.id, kz_records.map_id, kz_uid.last_name, kz_records.user_id, kz_records.time FROM kz_uid INNER JOIN kz_records ON kz_uid.id = kz_records.user_id INNER JOIN kz_maps ON kz_maps.id = kz_records.map_id WHERE kz_records.weapon = 6 AND kz_records.is_pro_record = 1 ORDER BY kz_records.user_id")
		if err != nil {
			panic(err)
		}

		var proRecords []ProRecords
		for res.Next() {
			var proRecord ProRecords
			var Dyblicat bool
			err = res.Scan(&proRecord.RecID, &proRecord.MapID, &proRecord.PlayerName, &proRecord.PlayerID, &proRecord.Time)
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

		playersInfo = GetTopPlayers(tops, playersInfo)
		err = t.ExecuteTemplate(w, "players", playersInfo)
		if err != nil {
			panic(err)
		}
	}
}
