package main

import (
	"fmt"
)

type InputMessage struct {
	Up      bool    `json:"up"`
	Down    bool    `json:"down"`
	Left    bool    `json:"reft"`
	Right   bool    `json:"right"`
	Angular float32 `json:"ang"`
	Shot    bool    `json:"shot"`
}

type DynamycObject struct {
	Name string

	Hp    float32 `json:"hp"`
	HpCap float32 `json:"hpcap"`

	X     float32 `json:"x"`
	Y     float32 `json:"y"`
	Xsize float32 `json:"xsize"`
	Ysize float32 `json:"ysize"`

	Velocity       float32 `json:"velocity"`
	VelocityBarior float32
}

type Bullet struct {
	Object   *DynamycObject `json:"object"`
	Damage   float32
	PlayerId int8 `json:"playerid"`
}

type Weapon struct {
	Id       int
	Name     string
	FireRate float32
	Magazine int
	Bullet   *Bullet
	Coast    int
}

//type Item struct {}

type Player struct {
	Object     *DynamycObject `json:"object"`
	CashPoints float32        `json:"cash"`
	Nickname   string         `json:"nickname"`
	Id         int            `json:"id"`
	Angular    int            `json:"ang"`
	Weapon     *Weapon
	//Inventory 		[]DynamycObject
}

type Adv struct {
	Object  *DynamycObject `json:"object"`
	Url     string
	Pict    string `json:"pict"`
	XTarget float32
	YTarget float32
	Angular float32 `json:"ang"`
}

type Barior struct {
	Id     int            `json:"id"`
	Object *DynamycObject `json:"object"`
}

type Map struct {
	TileSize int   `json:"tailsize"`
	SizeX    int   `json:"sizex"`
	SizeY    int   `json:"sizey"`
	Field    []int `json:"field"`
}

type GameObjects struct {
	Players []Player `json:"palyers"`
	Advs    []Adv    `json:"advs"`
	Bullets []Bullet `json:"bullets"`
	Bariors []Barior `json:"bariors"`
}

type Game struct {
	GameObjects GameObjects `json:"gameobjects"`
	Map         Map         `json:"map"`
	Wave        int         `json:"wave"`
	Url         string      `json:"url"`
	Stage       string      `json:"stage"`
}

// Создание игры
// Проинициализировать карту
// Заполнить массив объектов
func GameIni(roomPlayers map[string]*UserConnection) (game *Game) {
	// MapGeneration
	// Pla
	return
}

// Создание карты
func MapGeneration() (newMap *Map) {
	return
}

// Создание Игроков
func PlayersCreate(roomPlayers map[string]*UserConnection) (players []Player) {
	return
}

// Создание рекламы
func AdvsCreate() (advs []Adv) {
	return
}

// Создание пули
func (p *Player) BulletsCreate() (bs []Bullet) {
	return
}

// аппендит пули в массив пулей
func (p *Player) Shot() {

}

func BariorsCreate() (bariors []Barior) {
	return
}

// Хардкод дефолного оружия
func (p *Player) CreateDefaultWeapon() (w *Weapon) {
	return
}

// Обращение к бд для смены оружия(покупка)
func (p *Player) ChangeWeapon() (w *Weapon) {
	return
}

// Меняет состоние координат внутри поля Объект
func (p *Player) Move() {
	//wasd
}

//	Собирает все объекты для проверки на колизии
func (g *Game) DynamicObjectsCollector() (objs []DynamycObject) {
	return
}

// Управляет исходом колизий
func CollisionMeneger(objects []DynamycObject) {

}

// Проверяет, уничтожены ли все объекты
// Если да - заканчивает волну и игра переходит в паузу
func (g *Game) IsEndOfWave() {

}

// Входная точка для изменения состояния игры
// Принимает в себя структуру, которая получилась после разкодирования из json
func (g *Game) EventListener(mes InputMessage, nickname string) {

}

func (r *Room) GameEngine() {
	// GameIni
	var message InputMessage

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
