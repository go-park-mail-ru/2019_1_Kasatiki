package room

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/game_logic"
)

func (r *Room) GameEngine() {
	// GameIni
	game, re := game_logic.GameIni(r.Players)
	var message game_logic.InputMessage

	var keys []string
	for k, _ := range r.Players {
		keys = append(keys, k)
	}
	re.Id = 1
	r.Players[keys[0]].Connection.WriteJSON(&re)
	re.Id = 2
	r.Players[keys[1]].Connection.WriteJSON(&re)
	for {

		// TODO ХАРДКОД НО ЛЕТАЮЩИЙ
		if r.Players[keys[0]].TypeGame == "Multiplayer" {
			select {
			// Если есть сигнал от 1го игрока - оправляем его 2му игроку
			case message = <-r.Messenger.Player_From[keys[0]]:
				//start := time.Now()
				res := game.EventListener(message, r.Players[keys[0]].Login)
				//Возвращаем структуру Game
				//end := time.Now()
				//potracheno := start.Nanosecond() - end.Nanosecond()
				//fmt.Println("Time : ", potracheno , " Size: ",unsafe.Sizeof(res))
				//fmt.Println(res.Players[0].Object.X , " ", res.Players[0].Object.Y)
				r.Players[keys[1]].Connection.WriteJSON(&res)
				//r.Messenger.Player_To[keys[1]] <- message
			// Если есть сигнал от 2го игрока -  оправляем его 1му игроку
			case message = <-r.Messenger.Player_From[keys[1]]:
				res := game.EventListener(message, r.Players[keys[1]].Login)
				r.Players[keys[0]].Connection.WriteJSON(&res)
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

		res := &game_logic.BulletStatus{}
		var bs []*game_logic.Bullet
		// fmt.Println(len(game.GameObjects.Bullets))
		for i, _ := range game.GameObjects.Bullets {
			// fmt.Println("bullet ", i)
			game.GameObjects.Bullets[i].Run()
			bs = append(bs, game.GameObjects.Bullets[i])
		}
		res.Bullets = bs
		r.Players[keys[0]].Connection.WriteJSON(&res)
		r.Players[keys[1]].Connection.WriteJSON(&res)
	}
	fmt.Println("salaaam")

	return
}
