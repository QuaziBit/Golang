package components

// Player structure
type Player struct {
	id      int
	name    string
	isHuman bool
	points  int
}

// Players is slice of players
var players []Player

// SetPlayerInfo function used to initialize player
func (p Player) SetPlayerInfo(id int, name string, isHuman bool) {
	playerInfo := Player{id, name, isHuman, 0}
	//Players = make([]Player, 0, 2)
	players = append(players, playerInfo)
}

// UpdatePoints update player points
func (p Player) UpdatePoints(playerName string, points int) {
	if players[0].name == playerName {
		players[0].points = points
	} else if players[1].name == playerName {
		players[1].points = points
	}
}

// GetPlayerInfo get all player information
func (p Player) GetPlayerInfo(isUser bool) (id int, name string, isHuman bool, points int) {
	var tmp Player
	if players[0].isHuman == isUser {
		tmp = players[0]
	} else if players[1].isHuman == isUser {
		tmp = players[1]
	}
	return tmp.id, tmp.name, tmp.isHuman, tmp.points
}
