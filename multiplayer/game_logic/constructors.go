package game_logic

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/connections"
	"reflect"
)

// Создание игры
// Проинициализировать карту
// Заполнить массив объектов
func GameIni(roomPlayers map[string]*connections.UserConnection) (*Game, StartGame) {
	var game Game
	var res StartGame
	game.Map = MapGeneration()
	game.GameObjects = &GameObjects{}
	game.GameObjects.Players = PlayersCreate(roomPlayers, game.Map)
	game.GameObjects.Advs = AdvsCreate(10, game.Map, game.GameObjects.Players)
	res.Map = *game.Map
	for _, p := range game.GameObjects.Players {
		var info PlayerInfo
		info.Object = p.Object
		info.Id = p.Id
		info.CashPoints = p.CashPoints
		info.Nickname = p.Nickname
		info.Id = p.Id
		res.Players = append(res.Players, info)
	}
	for _, p := range game.GameObjects.Advs {
		var info AdvInfo
		info.Object = p.Object
		res.Advs = append(res.Advs, info)
	}
	return &game, res
}

// Создание карты
//func MapGeneration() (newMap *Map) {
//	return
//}

//type DynamycObject struct {
//	Name string
//
//	Hp    float32 `json:"hp"`
//	HpCap float32 `json:"hpcap"`
//
//	X     float32 `json:"x"`
//	Y     float32 `json:"y"`
//	Xsize float32 `json:"xsize"`
//	Ysize float32 `json:"ysize"`
//
//	Velocity       float32 `json:"velocity"`
//	VelocityBarior float32
//}

// Создание Игроков
func PlayersCreate(roomPlayers map[string]*connections.UserConnection, gameMap *Map) (players map[string]*Player) {
	players = make(map[string]*Player)
	var id int
	for _, p := range roomPlayers {
		id++
		players[p.Login] = &Player{
			Nickname: p.Login,
			Id:       id,
		}
		players[p.Login].Spawn(gameMap.SizeX/2, gameMap.SizeY/2)
		players[p.Login].CreateDefaultWeapon()
	}
	return
}

// Создание рекламы
func AdvsCreate(count int, gameMap *Map, players map[string]*Player) (advs map[int]*Adv) {
	advs = make(map[int]*Adv, count)
	var id int
	// Достаем все ключи плееров
	keys := reflect.ValueOf(players).MapKeys()
	for i := 0; i < count; i++ {
		id++
		advs[i] = &Adv{
			// Сетим плеера в качестве цели
			// каждому плееру одинаковое количество реклам.
			Player: players[keys[len(keys) * i / count].Interface().(string)],
		}
		advs[i].Spawn(gameMap.SizeX/2, gameMap.SizeY/2)
	}
	return
}

func BariorsCreate() (bariors []Barior) {
	return
}
