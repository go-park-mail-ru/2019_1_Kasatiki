package room

import (
	gl "github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/game_logic"
)

func (r *Room) GameEngine() {
	game, re := gl.GameIni(r.Players)
	var message gl.InputMessage

	var keys []string
	for k, _ := range r.Players {
		keys = append(keys, k)
	}
	re.Id = 1
	r.Players[keys[0]].Connection.WriteJSON(&re)
	re.Id = 2
	if len(keys) > 1 {
		r.Players[keys[1]].Connection.WriteJSON(&re)
	}

	for {
		if len(keys) > 1 {
			if r.Players[keys[0]].TypeGame == "Multiplayer" {
				select {
				// Если есть сигнал от 1го игрока - оправляем его 2му игроку
				case message = <-r.Messenger.Player_From[keys[0]]:
					res := game.EventListener(message, r.Players[keys[0]].Login)
					r.Players[keys[1]].Connection.WriteJSON(&res)
				// Если есть сигнал от 2го игрока -  оправляем его 1му игроку
				case message = <-r.Messenger.Player_From[keys[1]]:
					res := game.EventListener(message, r.Players[keys[1]].Login)
					r.Players[keys[0]].Connection.WriteJSON(&res)
				}
			}

		}

	}
}
