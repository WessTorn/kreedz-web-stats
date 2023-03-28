package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"math"
	"net/http"
)

type LastPro struct {
	MapName    string
	MapID      int
	PlayerName string
	PlayerID   int
	Time       uint32
	FormatTime string
	Date       string
	Weapon     int
	TopInMap   int
}

type TopInMap struct {
	PlayerID int
	Time     uint32
}

func GetLastPro(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("frontend/lastpro.gohtml", "frontend/head.gohtml", "frontend/header.gohtml")

		res, err := db.Query("SELECT kz_maps.mapname, kz_records.map_id, kz_uid.last_name, kz_records.user_id, kz_records.time, kz_records.Date, kz_records.weapon FROM kz_uid INNER JOIN kz_records ON kz_uid.id = kz_records.user_id INNER JOIN kz_maps ON kz_maps.id = kz_records.map_id WHERE kz_records.is_pro_record = 1 AND kz_records.weapon = 6 ORDER BY kz_records.Date DESC")
		if err != nil {
			panic(err)
		}

		var arrLastPro []LastPro
		for res.Next() {
			var oneLastPro LastPro
			err = res.Scan(&oneLastPro.MapName, &oneLastPro.MapID, &oneLastPro.PlayerName, &oneLastPro.PlayerID, &oneLastPro.Time, &oneLastPro.Date, &oneLastPro.Weapon)
			if err != nil {
				panic(err)
			}
			oneLastPro.FormatTime = formatTime(math.Float32frombits(oneLastPro.Time))
			arrLastPro = append(arrLastPro, oneLastPro)
		}

		for i := 0; i < len(arrLastPro); i++ {
			arrLastPro[i].TopInMap = GetFirstInMap(arrLastPro, arrLastPro[i].MapID, arrLastPro[i].PlayerID)
		}

		err = t.ExecuteTemplate(w, "lastPro", arrLastPro)
		if err != nil {
			panic(err)
		}
	}
}

func GetFirstInMap(arrRec []LastPro, mapID int, PlayerID int) int {
	var top []TopInMap
	for i := 0; i < len(arrRec); i++ {
		var oneRec TopInMap
		if arrRec[i].MapID == mapID {
			oneRec.PlayerID = arrRec[i].PlayerID
			oneRec.Time = arrRec[i].Time
			top = append(top, oneRec)
		}
	}

	for i := 0; i < len(top)-1; i++ {
		for j := 0; j < (len(top) - 1 - i); j++ {
			if top[j].Time > top[j+1].Time {
				top[j], top[j+1] = top[j+1], top[j]
			}
		}
	}
	var num int
	for i := 0; i < len(arrRec); i++ {
		num++
		if top[i].PlayerID == PlayerID {
			return num
		}
	}
	return 0
}
