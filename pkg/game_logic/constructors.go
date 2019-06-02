package game_logic

import (
	"fmt"
	"reflect"

	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/connections"
)

// Создание игры
// Проинициализировать карту
// Заполнить массив объектов
func GameIni(roomPlayers map[string]*connections.UserConnection, advsData []*Adv) (*Game, StartGame) {
	fmt.Println("Game Iniiii")
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
	game.GameObjects.Advs = AdvsCreate(40, game.Map, game.GameObjects.Players, advsData)
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

func AdvsCreate(count int, gameMap *Map, players map[string]*Player, advsData []*Adv) (advs []*Adv) {
	var id int
	tileSize := gameMap.TileSize
	// Достаем все ключи плееров
	keys := reflect.ValueOf(players).MapKeys()
	counter := 0
	for i := 0; i < count; i++ {
		if counter >= len(advsData) {
			counter = 0
		} else {
			counter++
		}

		id++
		adv := &Adv{
			// Сетим плеера в качестве цели
			// каждому плееру одинаковое количество реклам.
			Player: players[keys[len(keys)*i/count].Interface().(string)],
			Url:    advsData[counter].Url,
			Pict:   "http://0.0.0.0:8080/AdvsImgs/" + advsData[counter].Pict,
		}
		x := 50
		y := 50
		if i%4 == 0 {
			x = (i%3)*4 + 1
			y = 40 + (i%3)*4
		} else if i%4 == 1 {
			x = 49 + (i%3)*4
			y = 98 - (i%3)*4
		} else if i%4 == 2 {
			x = 98 - (i%3)*4
			y = 40 + (i%3)*4
		} else if i%4 == 3 {
			x = 98 - (i%3)*4
			y = 98 - (i%3)*4
		}
		if (gameMap.Field[x-1][y] == 1) &&
			(gameMap.Field[x+1][y] == 1) &&
			(gameMap.Field[x][y-1] == 1) &&
			(gameMap.Field[x][y+1] == 1) {
			// ничего
		} else {
			adv.Spawn(x*tileSize, y*tileSize, tileSize)
			advs = append(advs, adv)
			gameMap.Field[x][y] = 0
		}
		// adv.Spawn(gameMap.SizeX/2 * i+1, gameMap.SizeY/2 * i, gameMap.TileSize)
		// advs = append(advs, adv)
	}

	return
}
