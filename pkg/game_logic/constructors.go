package game_logic

import (
	"fmt"
	"reflect"

	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/connections"
)

// Создание игры
// Проинициализировать карту
// Заполнить массив объектов
func GameIni(roomPlayers map[string]*connections.UserConnection) (*Game, StartGame) {
	fmt.Println("Game Ini")
	var game Game
	var res StartGame
	game.GameObjects = &GameObjects{}
	fmt.Println("Generation of map")
	game.Map, game.GameObjects.Barrier = MapGeneration()
	//game.GameObjects.Players = make(map[string]*Player)
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
	game.GameObjects.Advs = AdvsCreate(10, game.Map, game.GameObjects.Players)
	for _, p := range game.GameObjects.Advs {
		var info AdvInfo
		info.Object = p.Object
		res.Advs = append(res.Advs, info)
	}
	res.Barrier = game.GameObjects.Barrier

	game.ZonesIni()
	//fmt.Println("Число барьеров", len(game.GameObjects.Barrier))
	//
	//fmt.Println("Zones:", len(game.Zones))
	//
	//for _, z := range game.Zones {
	//	fmt.Printf("Zone numb : %d, StartX : %d, StartY : %d, EndX : %d, EndY : %d \n", z.Number, z.StartX, z.StartY, z.EndX, z.EndY)
	//}
	//for k, z := range game.StaticCollection {
	//	fmt.Println("Zone Numb: ", k, " Numbers: ", len(z))
	//	//for _, b := range z {
	//	//	fmt.Printf("Name : %s, StartX : %d, StartY : %d, EndX : %d, EndY : %d \n", b.Name, b.X, b.Y, b.Xsize, b.Ysize)
	//	//}
	//}
	return &game, res
}

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
		players[p.Login].Spawn(gameMap.SizeX*gameMap.TileSize/2+id*5*gameMap.TileSize, gameMap.SizeX*gameMap.TileSize/2, gameMap.TileSize, gameMap.TileSize)
		fmt.Printf("Player was spawned in X: %d,    Y : %d \n", players[p.Login].Object.X, players[p.Login].Object.Y)
		players[p.Login].CreateDefaultWeapon()
	}

	return
}

func AdvsCreate(count int, gameMap *Map, players map[string]*Player) (advs []*Adv) {
	var id int
	// Достаем все ключи плееров
	keys := reflect.ValueOf(players).MapKeys()
	for i := 0; i < count; i++ {
		id++
		adv := &Adv{
			// Сетим плеера в качестве цели
			// каждому плееру одинаковое количество реклам.
			Player: players[keys[len(keys)*i/count].Interface().(string)],
		}
		adv.Spawn(gameMap.SizeX/2, gameMap.SizeY/2, gameMap.TileSize)
		advs = append(advs, adv)
	}
	return
}
