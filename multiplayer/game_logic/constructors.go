package game_logic

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/connections"
	"fmt"
)

// Создание игры
// Проинициализировать карту
// Заполнить массив объектов
func GameIni(roomPlayers map[string]*connections.UserConnection) ( *Game) {
	var game Game
	
	game.Map = *MapGeneration()
	fmt.Println("game: ", game)
	return &game
}

// Создание карты
//func MapGeneration() (newMap *Map) {
//	return
//}

// Создание Игроков
func PlayersCreate(roomPlayers map[string]*connections.UserConnection) (players []Player) {
	return
}

// Создание рекламы
func AdvsCreate() (advs []Adv) {
	return
}

func BariorsCreate() (bariors []Barior) {
	return
}
