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

type MapRecord struct {
	Name       string
	Time       uint32
	FormatTime string
	Date       string
	PlayerID   int
	Weapon     int
	Tp         int
	Cp         int
	Top        int
	MapID      string
}

type TemplateData struct {
	Records []MapRecord
	MapName string
	MapId   string
}

func GetMapRecordNub(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("frontend/map.gohtml", "frontend/head.gohtml", "frontend/header.gohtml")
		vars := mux.Vars(r)

		res, err := db.Query(fmt.Sprintf("SELECT `last_name`, `user_id`, `time`, `date`, `weapon`, `cp`, `tp` FROM `kz_uid` as user INNER JOIN (SELECT * FROM `kz_records` WHERE `map_id` = '%s' AND `aa` = 0 AND `weapon` = 6 AND `is_pro_record` = 0 ORDER BY `time` LIMIT 15) as record ON user.id = record.user_id ORDER BY `time`;", vars["id"]))
		if err != nil {
			panic(err)
		}

		var records []MapRecord
		var top int
		for res.Next() {
			var Record MapRecord
			top++
			err = res.Scan(&Record.Name, &Record.PlayerID, &Record.Time, &Record.Date, &Record.Weapon, &Record.Cp, &Record.Tp)
			if err != nil {
				panic(err)
			}
			Record.FormatTime = formatTime(math.Float32frombits(Record.Time))
			Record.Top = top
			Record.MapID = fmt.Sprintf("%s", vars["id"])
			records = append(records, Record)
		}

		res2, err := db.Query(fmt.Sprintf("SELECT `mapname` FROM `kz_maps` WHERE `id` = '%s'", vars["id"]))
		if err != nil {
			panic(err)
		}
		var mapName string
		for res2.Next() {
			err = res2.Scan(&mapName)
			if err != nil {
				panic(err)
			}
		}

		data := TemplateData{
			Records: records,
			MapName: mapName,
			MapId:   vars["id"],
		}
		err = t.ExecuteTemplate(w, "mapRecordNub", data)
		if err != nil {
			panic(err)
		}
	}
}

func GetMapRecordPro(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("frontend/map.gohtml", "frontend/head.gohtml", "frontend/header.gohtml")
		vars := mux.Vars(r)

		res, err := db.Query(fmt.Sprintf("SELECT `last_name`, `user_id`, `time`, `date`, `weapon` FROM `kz_uid` as user INNER JOIN (SELECT * FROM `kz_records` WHERE `map_id` = '%s' AND `aa` = 0 AND `weapon` = 6 AND `is_pro_record` = 1 ORDER BY `time` LIMIT 15) as record ON user.id = record.user_id ORDER BY `time`;", vars["id"]))
		if err != nil {
			panic(err)
		}

		var records []MapRecord
		var top int
		for res.Next() {
			var Record MapRecord
			top++
			err = res.Scan(&Record.Name, &Record.PlayerID, &Record.Time, &Record.Date, &Record.Weapon)
			if err != nil {
				panic(err)
			}
			Record.FormatTime = formatTime(math.Float32frombits(Record.Time))
			Record.Top = top
			Record.MapID = fmt.Sprintf("%s", vars["id"])
			records = append(records, Record)
		}

		res2, err := db.Query(fmt.Sprintf("SELECT `mapname` FROM `kz_maps` WHERE `id` = '%s'", vars["id"]))
		if err != nil {
			panic(err)
		}
		var mapName string
		for res2.Next() {
			err = res2.Scan(&mapName)
			if err != nil {
				panic(err)
			}
		}

		data := TemplateData{
			Records: records,
			MapName: mapName,
			MapId:   vars["id"],
		}
		err = t.ExecuteTemplate(w, "mapRecordPro", data)
		if err != nil {
			panic(err)
		}
	}
}

func formatTime(time float32) (res string) {
	var iMin, iSec, iMS float64
	iMin = math.Floor(float64(time / 60.0))
	iSec = math.Floor(float64(time) - iMin*60.0)
	iMS = math.Floor((float64(time) - (iMin*60.0 + iSec)) * 100)
	res = fmt.Sprintf("%02d:%02d.%02d", int(iMin), int(iSec), int(iMS))
	return
}
