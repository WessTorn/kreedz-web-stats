package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"math"
	"net/http"
)

type Maps struct {
	Num     int
	MapID   int
	MapName string
}

type MapsRecFirst struct {
	Num        int
	MapID      int
	PlayerID   int
	MapName    string
	PlayerName string
	Time       uint32
	FormatTime string
	Date       string
	Weapon     int
	Tp         int
}

func GetMapsStatic(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("frontend/maps.gohtml", "frontend/head.gohtml", "frontend/header.gohtml")

		res, err := db.Query(fmt.Sprintf("SELECT DISTINCT kz_maps.id, kz_maps.mapname FROM kz_maps INNER JOIN kz_records ON kz_maps.id = kz_records.map_id ORDER BY kz_maps.mapname LIMIT 25"))
		if err != nil {
			panic(err)
		}

		var arrMaps []Maps
		for res.Next() {
			var maps Maps
			err = res.Scan(&maps.MapID, &maps.MapName)
			arrMaps = append(arrMaps, maps)
		}

		var arrMapsRecFirst []MapsRecFirst
		var num int
		for i := 0; i < len(arrMaps); i++ {
			res2, err := db.Query(fmt.Sprintf("SELECT `user_id`, `last_name`, `time`, `date`, `weapon`, `tp` FROM `kz_uid` as user INNER JOIN (SELECT * FROM `kz_records` WHERE `map_id` = '%d' AND `aa` = 0 AND `weapon` = 6 ORDER BY `time` LIMIT 1) as record ON user.id = record.user_id ORDER BY `time` LIMIT 1;", arrMaps[i].MapID))
			if err != nil {
				panic(err)
			}
			for res2.Next() {
				var mapsRecFirst MapsRecFirst
				err = res2.Scan(&mapsRecFirst.PlayerID, &mapsRecFirst.PlayerName, &mapsRecFirst.Time, &mapsRecFirst.Date, &mapsRecFirst.Weapon, &mapsRecFirst.Tp)
				if mapsRecFirst.PlayerID < 1 {
					continue
				}
				mapsRecFirst.MapID = arrMaps[i].MapID
				mapsRecFirst.MapName = arrMaps[i].MapName
				mapsRecFirst.FormatTime = formatTime(math.Float32frombits(mapsRecFirst.Time))
				num++
				mapsRecFirst.Num = num
				arrMapsRecFirst = append(arrMapsRecFirst, mapsRecFirst)
			}
		}

		err = t.ExecuteTemplate(w, "maps", arrMapsRecFirst)
		if err != nil {
			panic(err)
		}
	}
}
