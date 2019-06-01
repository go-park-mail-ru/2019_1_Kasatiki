package lobby

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/connections"
	rm "github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/lobby/room"
)

type GameToConnect struct {
	Room rm.RoomId
	P_id rm.PlayerId
}

type Lobby struct {
	// Все соединения на данный cd ..
	ProcessedPlayers map[string]GameToConnect
	// Все созданные комнаты
	Rooms map[rm.RoomId]*rm.Room
	// Номер последней комнаты (для сохранения уникальности RoomId)
	LastRoom rm.RoomId
	// Пользователь, ждущий подключения другого пользователя.
	WaitingConnection *connections.UserConnection
	// Канал, по которому сообщается, какая комната должна быть удалена
	DeleteRooms chan rm.RoomId
}

// Создаем новое лобби
func NewLobby() *Lobby {
	fmt.Println("Creating new Lobby")
	return &Lobby{
		ProcessedPlayers: make(map[string]GameToConnect),
		Rooms:            make(map[rm.RoomId]*rm.Room),
		DeleteRooms:      make(chan rm.RoomId, 5),
	}
}

// Запускаем лобби
func (lb *Lobby) Run(connectionQueue chan *connections.UserConnection) {
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

func (lb *Lobby) AddPlayer(connection *connections.UserConnection) error {
	// Проверка на то, что пользователь с таким ником в игре
	if connection.Login == "" {
		return errors.New("Bad login")
	}
	game, ok := lb.ProcessedPlayers[connection.Login]
	if ok {
		// Если игрок уже был в игре, но по какой-то причине отлетел - восстанавливаем соединение.
		fmt.Println("Reconnecting player")
		lb.Rooms[game.Room].Reconnect(connection)
		return nil
	}

	// Todo SinglePlayer
	if connection.TypeGame != "Multiplayer" {
		// Меняем id комнаты
		lb.LastRoom++
		return nil
	}

	// Если ждущего игрока нет - назначаем игрока
	if lb.WaitingConnection == nil {
		fmt.Println("New waiter")
		//log.Printf("Set connection user = '%s' as waiting", connection.Token)
		lb.WaitingConnection = connection
		return nil
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

	return nil
}

// Создаем новую комнату и коннектим двух игроков
// Возвращаем саму комнату и коннекты
func CreatingMatch(waiter *connections.UserConnection, current *connections.UserConnection,
	DeleteRooms chan rm.RoomId, LastRoom rm.RoomId) (room *rm.Room, Conn1 GameToConnect, Conn2 GameToConnect) {
	fmt.Println("Creating new Room and connects")
	// Создаем новую комнату
	var players []*connections.UserConnection
	players = append(players, waiter, current)
	room = rm.NewRoom(players, DeleteRooms, LastRoom)
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
func (lb *Lobby) DeletingRoom(roomId rm.RoomId) {
	// Достаем экземпляр комнаты из набора существущих комнат
	room, ok := lb.Rooms[roomId]
	if !ok {
		fmt.Print("Deleting room error: wrong roomId:" + roomId.String())
		return
	}
	for _, p := range room.Players {
		delete(lb.ProcessedPlayers, p.Login)
	}
	delete(lb.Rooms, roomId)
	fmt.Println("Room with roomId" + roomId.String() + "deleted!")
	return
}
