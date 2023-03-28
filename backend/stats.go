package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type ProRecords struct {
	RecID      int
	MapID      int
	PlayerName string
	PlayerID   int
	Time       uint32
}

type PlayerInfo struct {
	ID       int
	Name     string
	Rec      []MapRec
	FirstRec int
	Records  int
	Points   int
	Top      int
}

type MapRec struct {
	MapID int
}

type MapRecTop struct {
	PlayerID int
	Time     uint32
}

type TopsInMap struct {
	MapID int
	Rec   []MapRecTop
}

type Player struct {
	ID       int
	Name     string
	SteamID  string
	PlayTime int
	Rec      []MapRec
	FirstRec int
	Records  int
	Points   int
}

func GetPlayerInfo(records []ProRecords, players []PlayerInfo) []PlayerInfo {
	var playerID int
	for i := 0; i < len(records); i++ {
		var playerInfo PlayerInfo
		var maprec MapRec
		playerInfo.Name = records[i].PlayerName
		playerInfo.ID = records[i].PlayerID
		maprec.MapID = records[i].MapID

		if playerID == 0 {
			playerInfo.Records++
			playerInfo.Rec = append(playerInfo.Rec, maprec)
			players = append(players, playerInfo)
			playerID++
		} else if playerInfo.ID == players[playerID-1].ID {
			players[playerID-1].Records++
			players[playerID-1].Rec = append(players[playerID-1].Rec, maprec)
		} else {
			playerInfo.Records++
			playerInfo.Rec = append(playerInfo.Rec, maprec)
			players = append(players, playerInfo)
			playerID++
		}
	}
	return players
}

func GetMapTopRecs(records []ProRecords, tops []TopsInMap) []TopsInMap {
	for i := 0; i < (len(records) - 1); i++ {
		for j := 0; j < ((len(records) - 1) - i); j++ {
			if records[j].MapID > records[j+1].MapID {
				records[j], records[j+1] = records[j+1], records[j]
			}
		}
	}

	var topID int
	for i := 0; i < len(records); i++ {
		var top TopsInMap
		var maprec MapRecTop

		maprec.PlayerID = records[i].PlayerID
		maprec.Time = records[i].Time
		top.MapID = records[i].MapID

		if topID == 0 {
			topID++
			top.Rec = append(top.Rec, maprec)
			tops = append(tops, top)
		} else if top.MapID == tops[topID-1].MapID {
			tops[topID-1].Rec = append(tops[topID-1].Rec, maprec)
		} else {
			topID++
			top.Rec = append(top.Rec, maprec)
			tops = append(tops, top)
		}
	}
	for i := 0; i < len(tops); i++ {
		for u := 0; u < len(tops[i].Rec)-1; u++ {
			for j := 0; j < (len(tops[i].Rec) - 1 - u); j++ {
				if tops[i].Rec[j].Time > tops[i].Rec[j+1].Time {
					tops[i].Rec[j], tops[i].Rec[j+1] = tops[i].Rec[j+1], tops[i].Rec[j]
				}
			}
		}
	}
	return tops
}

func GetTopPlayers(tops []TopsInMap, players []PlayerInfo) []PlayerInfo {
	for p := 0; p < len(players); p++ {
		for r := 0; r < len(players[p].Rec); r++ {
			for t := 0; t < len(tops); t++ {
				if players[p].Rec[r].MapID != tops[t].MapID {
					continue
				}
				var rankInMap int
				for g := 0; g < len(tops[t].Rec); g++ {
					rankInMap++
					if players[p].ID != tops[t].Rec[g].PlayerID {
						continue
					}
					if rankInMap == 1 {
						players[p].FirstRec++
					}
					players[p].Points += getPoints(rankInMap)
					break
				}
			}
		}
	}

	for i := 0; i < (len(players) - 1); i++ {
		for j := 0; j < ((len(players) - 1) - i); j++ {
			if players[j].Points < players[j+1].Points {
				players[j], players[j+1] = players[j+1], players[j]
			}
		}
	}

	for i := 0; i < len(players); i++ {
		players[i].Top = i + 1
	}
	return players
}

func getPoints(rankInMap int) (res int) {
	switch rankInMap {
	case 1:
		res = 30
	case 2:
		res = 29
	case 3:
		res = 28
	case 4:
		res = 27
	case 5:
		res = 26
	case 6:
		res = 25
	case 7:
		res = 24
	case 8:
		res = 23
	case 9:
		res = 22
	case 10:
		res = 21
	case 11:
		res = 20
	case 12:
		res = 19
	case 13:
		res = 18
	case 14:
		res = 17
	default:
		res = 16
	}
	return
}

func GetTopPlayer(tops []TopsInMap, player Player) Player {
	var rankInMap int
	for r := 0; r < len(player.Rec); r++ {
		for t := 0; t < len(tops); t++ {
			if player.Rec[r].MapID == tops[t].MapID {
				for g := 0; g < len(tops[t].Rec); g++ {
					rankInMap++
					if player.ID == tops[t].Rec[g].PlayerID {
						switch rankInMap {
						case 1:
							player.FirstRec++
							player.Points += 30
						case 2:
							player.Points += 29
						case 3:
							player.Points += 28
						case 4:
							player.Points += 27
						case 5:
							player.Points += 26
						case 6:
							player.Points += 25
						case 7:
							player.Points += 24
						case 8:
							player.Points += 23
						case 9:
							player.Points += 22
						case 10:
							player.Points += 21
						case 11:
							player.Points += 20
						case 12:
							player.Points += 19
						case 13:
							player.Points += 18
						case 14:
							player.Points += 17
						default:
							player.Points += 16
						}
						rankInMap = 0
						continue
					}
				}
			}
		}
		rankInMap = 0
	}
	return player
}
