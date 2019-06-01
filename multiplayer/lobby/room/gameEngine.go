package room

import (
	"fmt"

	gl "github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/game_logic"
)

func (r *Room) GameEngine() {
	// GameIni
	game, re := gl.GameIni(r.Players)
	var message gl.InputMessage

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

		// Стрельба

		res := &gl.BulletStatus{}
		var bs []*gl.Bullet
		objs := gl.GetGameObjs()
		// fmt.Println(objs[1])
		for i := 0; i < len(game.GameObjects.Bullets); i++ {
			game.GameObjects.Bullets[i].Run()
			for j := 0; j < len(objs); j++ {
				if i == len(game.GameObjects.Bullets) {
					break
				}
				if gl.IsCollision(game.GameObjects.Bullets[i].Object, objs[j]) {
					// fmt.Println("object ", j, game.GameObjects.Bullets[i].Object.X, game.GameObjects.Bullets[i].Object.Y)
					if objs[j].Name != "Player" {
						game.GameObjects.Bullets = append(game.GameObjects.Bullets[:i], game.GameObjects.Bullets[i+1:]...)
						// game.GameObjects.Bullets[i] = game.GameObjects.Bullets[len(game.GameObjects.Bullets) - 1]
						// fmt.Println("get 1", len(game.GameObjects.Bullets), i)
						// game.GameObjects.Bullets = game.GameObjects.Bullets[:len(game.GameObjects.Bullets) - 1]
						// fmt.Println("get 2")
					}
				}
			}
			if i == len(game.GameObjects.Bullets) {
				break
			}
			bs = append(bs, game.GameObjects.Bullets[i])
		}
		res.Bullets = bs
		r.Players[keys[0]].Connection.WriteJSON(&res)
		r.Players[keys[1]].Connection.WriteJSON(&res)
	}
	fmt.Println("salaaam")

	return
}
