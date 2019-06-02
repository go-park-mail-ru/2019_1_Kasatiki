package room

import (
	gl "github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/game_logic"
	"time"
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
	ticker := time.NewTicker(time.Second / 30)
	gs := gl.GameStatus{}
	res := &gl.BulletStatus{}
	for {

		if len(keys) > 1 {
			if r.Players[keys[0]].TypeGame == "Multiplayer" {
				select {
				// Если есть сигнал от 1го игрока - оправляем его 2му игроку
				case message = <-r.Messenger.Player_From[keys[0]]:
					gs = game.EventListener(message, r.Players[keys[0]].Login)
				// Если есть сигнал от 2го игрока -  оправляем его 1му игроку
				case message = <-r.Messenger.Player_From[keys[1]]:
					gs = game.EventListener(message, r.Players[keys[1]].Login)
				case <-ticker.C:
					r.Players[keys[0]].Connection.WriteJSON(&gs)
					r.Players[keys[1]].Connection.WriteJSON(&gs)
				}
				var bs []*gl.Bullet
				objs := gl.GetGameObjs()
				for i := 0; i < len(game.GameObjects.Bullets); i++ {
					for j := 0; j < len(game.GameObjects.Advs); j++ {
						if i == len(game.GameObjects.Bullets) {
							break
						}
						if game.GameObjects.Bullets[i].IsCollisionInWay(game.GameObjects.Advs[j].Object) {
							game.GameObjects.Advs[j].Object.Hp -= game.GameObjects.Bullets[i].Damage
							game.GameObjects.Bullets = append(game.GameObjects.Bullets[:i], game.GameObjects.Bullets[i+1:]...)
							if game.GameObjects.Advs[j].Object.Hp == 0 {
								game.GameObjects.Advs = append(game.GameObjects.Advs[:j], game.GameObjects.Advs[j+1:]...)
							}
							// break чтобы он не декрементил hp у всех реклам
							break
						}
					}
				}
				for i := 0; i < len(game.GameObjects.Bullets); i++ {
					game.GameObjects.Bullets[i].Run()
					// проверяем, вышла ли пуля за пределы карты
					if !game.GameObjects.Bullets[i].IsOnMap(game.Map) {
						game.GameObjects.Bullets = append(game.GameObjects.Bullets[:i], game.GameObjects.Bullets[i+1:]...)
						continue
					}
					for j := 0; j < len(objs); j++ {
						if i == len(game.GameObjects.Bullets) {
							break
						}
						if gl.IsCollision(game.GameObjects.Bullets[i].Object, objs[j]) {
							if objs[j].Name != "Player" {
								game.GameObjects.Bullets = append(game.GameObjects.Bullets[:i], game.GameObjects.Bullets[i+1:]...)
							}
						}
					}
					if i == len(game.GameObjects.Bullets) {
						break
					}
					bs = append(bs, game.GameObjects.Bullets[i])
				}
				res.Bullets = bs
			}
			r.Players[keys[0]].Connection.WriteJSON(&res)
			r.Players[keys[1]].Connection.WriteJSON(&res)
		} else {
			select {
			// Если есть сигнал от 1го игрока - оправляем его 2му игроку
			case message = <-r.Messenger.Player_From[keys[0]]:
				gs = game.EventListener(message, r.Players[keys[0]].Login)
			// Если есть сигнал от 2го игрока -  оправляем его 1му игроку
			case <-ticker.C:
				r.Players[keys[0]].Connection.WriteJSON(&gs)
			}
			var bs []*gl.Bullet
			objs := gl.GetGameObjs()
			for i := 0; i < len(game.GameObjects.Bullets); i++ {
				for j := 0; j < len(game.GameObjects.Advs); j++ {
					if i == len(game.GameObjects.Bullets) {
						break
					}
					if game.GameObjects.Bullets[i].IsCollisionInWay(game.GameObjects.Advs[j].Object) {
						game.GameObjects.Advs[j].Object.Hp -= game.GameObjects.Bullets[i].Damage
						game.GameObjects.Bullets = append(game.GameObjects.Bullets[:i], game.GameObjects.Bullets[i+1:]...)
						if game.GameObjects.Advs[j].Object.Hp == 0 {
							game.GameObjects.Advs = append(game.GameObjects.Advs[:j], game.GameObjects.Advs[j+1:]...)
						}
						// break чтобы он не декрементил hp у всех реклам
						break
					}
				}
			}
			for i := 0; i < len(game.GameObjects.Bullets); i++ {
				game.GameObjects.Bullets[i].Run()
				// проверяем, вышла ли пуля за пределы карты
				if !game.GameObjects.Bullets[i].IsOnMap(game.Map) {
					game.GameObjects.Bullets = append(game.GameObjects.Bullets[:i], game.GameObjects.Bullets[i+1:]...)
					continue
				}
				for j := 0; j < len(objs); j++ {
					if i == len(game.GameObjects.Bullets) {
						break
					}
					if gl.IsCollision(game.GameObjects.Bullets[i].Object, objs[j]) {
						if objs[j].Name != "Player" {
							game.GameObjects.Bullets = append(game.GameObjects.Bullets[:i], game.GameObjects.Bullets[i+1:]...)
						}
					}
				}
				if i == len(game.GameObjects.Bullets) {
					break
				}
				bs = append(bs, game.GameObjects.Bullets[i])
			}
			res.Bullets = bs
		}
		r.Players[keys[0]].Connection.WriteJSON(&res)
	}
}
