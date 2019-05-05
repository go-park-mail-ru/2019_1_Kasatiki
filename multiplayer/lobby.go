package main

import (
	"fmt"
	"strconv"
)

// адрес игровой комнаты, уникальный для данного сервера.
type RoomId int

func (id RoomId) String() string {
	return strconv.Itoa(int(id))
}

type PlayerId int

func (id PlayerId) String() string {
	return strconv.Itoa(int(id))
}

type GameToConnect struct {
	Room RoomId
	P_id PlayerId
}

// Персонаж в представлении сервера.
//type Сharacter struct {
//	Role         PlayerId
//	//Weapon       Weapon
//	ShowedWeapon bool
//}

type Lobby struct {
	// Все соединения на данный момент
	ProcessedPlayers map[string]GameToConnect
	// Все созданные комнаты
	Rooms map[RoomId]*Room
	// Номер последней комнаты (для сохранения уникальности RoomId)
	LastRoom RoomId
	// Пользователь, ждущий подключения другого пользователя.
	WaitingConnection *UserConnection
	// Канал, по которому сообщается, какая комната должна быть удалена
	DeleteRooms chan RoomId
}

// Создаем новое лобби
func NewLobby() *Lobby {
	fmt.Println("Creating new Lobby")
	return &Lobby{
		ProcessedPlayers: make(map[string]GameToConnect),
		Rooms:            make(map[RoomId]*Room),
		DeleteRooms:      make(chan RoomId, 5),
	}
}

// Запускаем лобби
func (lb *Lobby) Run(connectionQueue chan *UserConnection) {
	// Пока есть комнаты и народ в очереди
	for connectionQueue != nil && lb.DeleteRooms != nil {
		select {
		// Если пришел сигнал на удаление комнаты
		case RoomId, ok := <-lb.DeleteRooms:
			if ok {
				lb.DeletingRoom(RoomId)
			} else {
				lb.DeleteRooms = nil
			}
		// Если пришел сигнал на коннект нового пользователя
		case connection, ok := <-connectionQueue:
			if ok {
				lb.AddPlayer(connection)
			} else {
				connectionQueue = nil
			}
		}
	}
	return
}

// Логика работы функции следующая:
//	1) 	Если пользователь уже играл - возвращаем в игру
// 	2) 	Если пользователь новый, но он первый - он становится ждущим
// 	3) 	Если пользователь новый, но он не первый - он становится текущим,
//		соответственно создаем новую комнату

func (lb *Lobby) AddPlayer(connection *UserConnection) {
	// Проверка на то, что пользователь с таким ником в игре
	game, ok := lb.ProcessedPlayers[connection.Login]
	if ok {
		// Если игрок уже был в игре, но по какой-то причине отлетел - восстанавливаем соединение.
		fmt.Println("Reconnecting player")
		lb.Rooms[game.Room].Reconnect(connection, game.P_id)
		return
	}
	// Если ждущего игрока нет - назначаем игрока
	if lb.WaitingConnection == nil {
		fmt.Println("New waiter")
		//log.Printf("Set connection user = '%s' as waiting", connection.Token)
		lb.WaitingConnection = connection
		return
	}

	// Если дошли до сюда - значит у нас есть 2 свободных игрока - ждущий и текущий
	// Создадим комнату и законектим игроков

	//комната				ждущий игрок									текущий игрок
	lb.Rooms[lb.LastRoom], lb.ProcessedPlayers[lb.WaitingConnection.Token], lb.ProcessedPlayers[connection.Token] = CreatingMatch(lb.WaitingConnection,
		connection, lb.DeleteRooms, lb.LastRoom)

	// Ждущего больше нет
	lb.WaitingConnection = nil

	// Меняем id комнаты
	lb.LastRoom++

	return
}

// Создаем новую комнату и коннектим двух игроков
// Возвращаем саму комнату и коннекты
func CreatingMatch(waiter *UserConnection, current *UserConnection,
	DeleteRooms chan RoomId, LastRoom RoomId) (room *Room, Conn1 GameToConnect, Conn2 GameToConnect) {
	fmt.Println("Creating new Room and connects")
	// Создаем новую комнату
	room = NewRoom(waiter, current, DeleteRooms, LastRoom)
	// Создаем коннект для первого игрока(Ждущий)
	Conn1 = GameToConnect{
		Room: LastRoom,
		P_id: 0,
	}
	// Создаем коннект для первого игрока(Текущий)
	Conn2 = GameToConnect{
		Room: LastRoom,
		P_id: 1,
	}
	return
}

// Функция удаления комнаты
func (lb *Lobby) DeletingRoom(roomId RoomId) {
	// Достаем экземпляр комнаты из набора существущих комнат
	room, ok := lb.Rooms[roomId]
	if !ok {
		fmt.Print("Deleting room error: wrong roomId:" + roomId.String())
		return
	}
	delete(lb.ProcessedPlayers, room.Player_1.Token)
	delete(lb.ProcessedPlayers, room.Player_2.Token)
	delete(lb.Rooms, roomId)

	fmt.Println("Room with roomId" + roomId.String() + "deleted!")
	return
}
