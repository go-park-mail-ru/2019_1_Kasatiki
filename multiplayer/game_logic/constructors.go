package game_logic

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/connections"
)

// Создание игры
// Проинициализировать карту
// Заполнить массив объектов
func GameIni(roomPlayers map[string]*connections.UserConnection) (*Game, StartGame) {
	var game Game
	var res StartGame
	game.GameObjects = &GameObjects{}
	game.Map, game.GameObjects.Barrier = MapGeneration()
	game.GameObjects.Players = make(map[string]*Player)
	game.GameObjects.Players = PlayersCreate(roomPlayers, game.Map)
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
	res.Barrier = game.GameObjects.Barrier
	//for _, b := range res.Barrier {
	//	fmt.Println(b.Object.X)
	//}
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
		players[p.Login].Spawn(gameMap.SizeX * gameMap.TileSize / 2 + 5 * gameMap.TileSize / 2, gameMap.SizeX * gameMap.TileSize / 2 + 5 * gameMap.TileSize / 2, gameMap.TileSize, gameMap.TileSize)
		players[p.Login].CreateDefaultWeapon()
	}
	return
}

// Создание рекламы
func AdvsCreate() (advs []Adv) {
	return
}
