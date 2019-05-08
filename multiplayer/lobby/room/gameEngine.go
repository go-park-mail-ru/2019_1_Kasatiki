package room

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/game_logic"
)

func (r *Room) GameEngine() {
	// GameIni
	//game := game_logic.GameIni(r.Players)
	var message game_logic.InputMessage

	var keys []string
	for k, _ := range r.Players {
		keys = append(keys, k)
	}

	for {

		// TODO ХАРДКОД НО ЛЕТАЮЩИЙ
		if r.Players[keys[0]].TypeGame == "Multiplayer" {
			select {
			// Если есть сигнал от 1го игрока - оправляем его 2му игроку
			case message = <-r.Messenger.Player_From[keys[0]]:

				// EventListener()

				//Возвращаем структуру Game
				r.Players[keys[1]].Connection.WriteJSON(&message)
				//r.Messenger.Player_To[keys[1]] <- message
			// Если есть сигнал от 2го игрока -  оправляем его 1му игроку

			case message = <-r.Messenger.Player_From[keys[1]]:
				r.Players[keys[0]].Connection.WriteJSON(&message)
				//r.Messenger.Player_To[keys[0]] <- message
			}
		}

		//TODO ЛАГАЕТ НО ГИБКО

		//for k, from := range r.Messenger.Player_From {
		//	select {
		//	// Если есть сигнал от игрока - оправляем сопернику(Todo всем)
		//	case message, _ := <- from:
		//		//for _, _ := range keys {
		//		r.Messenger.Player_To[k] <- message
		//		//}
		//
		//		continue
		//		//for _, all := range r.Players {
		//		//	if all.Login != key {
		//		//		//r.Players[all.Login].Connection.WriteJSON(&message)
		//		//		r.Messenger.Player_To[all.Login]} <- message
		//		//	}
		//		//}
		//	}
		//}

	}
	fmt.Println(message)

	return
}
