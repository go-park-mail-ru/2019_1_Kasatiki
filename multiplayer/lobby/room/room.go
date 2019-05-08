package room

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/connections"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/multiplayer/game_logic"
	"log"
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

// Пока не учитывал игровую логику
// Здесь чисто игроки и каналы общения

type Room struct {
	Players map[string]*connections.UserConnection

	// Каналы для устранения дисконекта
	Recovery struct {
		Player_IsAvailableRead  map[string]chan struct{}
		Player_IsAvailableWrite map[string]chan struct{}
	}

	// Каналы, c помощью которых go room.GameMaster() общается с внешним миром
	Messenger struct {
		Player_From map[string]chan game_logic.InputMessage
		Player_To   map[string]chan game_logic.Game
	}

	// Созданную комнату необходимо будет отправить в канал комнат которые мы собираемся уничтожить
	Suicide chan RoomId

	// Номер текущей комнаты
	CurrentRoomId RoomId
}

// Создание новой комнаты
func NewRoom(players []*connections.UserConnection, completedRooms chan RoomId, ownNumber RoomId) (room *Room) {
	// Инициализируем структуру комнаты
	room = &Room{
		Suicide:       completedRooms,
		CurrentRoomId: ownNumber,
	}

	// Создаем мапы
	room.Players = make(map[string]*connections.UserConnection)
	room.Messenger.Player_From = make(map[string]chan game_logic.InputMessage)
	room.Messenger.Player_To = make(map[string]chan game_logic.Game)
	room.Recovery.Player_IsAvailableRead = make(map[string]chan struct{})
	room.Recovery.Player_IsAvailableWrite = make(map[string]chan struct{})

	for _, p := range players {
		room.Players[p.Login] = p
		room.Messenger.Player_From[p.Login] = make(chan game_logic.InputMessage, 5)
		room.Messenger.Player_To[p.Login] = make(chan game_logic.Game, 5)
		room.Recovery.Player_IsAvailableRead[p.Login] = make(chan struct{}, 1)
		room.Recovery.Player_IsAvailableWrite[p.Login] = make(chan struct{}, 1)
		go room.WebSocketReader(p.Login)
		go room.WebSocketWriter(p.Login)
	}

	go room.GameEngine()

	//	Какова работа логики комнаты?
	// Внутри каждая комната управляется одним GameEngine - горутиной.

	// 2 обслуживающие горутины нужны для изоляции соединения от игровой логики.(Синглплеер)
	//    	╭─Player_1_From─>>─╮
	// Player_1            GameEngine
	//    	╰─Player_1_To───<<─╯

	// 4 обслуживающие горутины нужны для изоляции соединения от игровой логики.(Мультплеер)
	//    	╭─Player_1_From─>>─╮      ╭─<<─Player_2_From─╮
	// Player_1         	  GameEngine            	Player_2
	//    	╰─Player_1_To───<<─╯      ╰─>>─Player_2_To───╯

	fmt.Println("Room created with")
	for i, p := range players {
		fmt.Println("Romm created with Player" + strconv.Itoa(i) + ": " + p.Login)
	}
	return
}

// Закрываем лавочку и горутины
func (r *Room) StopRoom() {
	for _, p := range r.Players {
		close(r.Recovery.Player_IsAvailableRead[p.Login])
		close(r.Messenger.Player_To[p.Login])
		close(r.Recovery.Player_IsAvailableWrite[p.Login])
	}
	return
}

// Отпраление текущего id комнаты в канал "для удаления"
func (r *Room) RemoveRoom() {
	r.Suicide <- r.CurrentRoomId
	return
}

// Восстановление соединения
func (r *Room) Reconnect(user *connections.UserConnection) {
	fmt.Println("Reconnect sessioni = '%s' as role %d", user.Token, user.Login)

	for _, p := range r.Players {
		if p != nil {
			_ = r.Players[p.Login].Connection.Close()
		}
		r.Players[p.Login] = user
		select {
		case r.Recovery.Player_IsAvailableRead[p.Login] <- struct{}{}:
		default:
		}
		select {
		case r.Recovery.Player_IsAvailableWrite[p.Login] <- struct{}{}:
		default:
		}
	}

	return
}

// Обслуживающая горутина From
func (r *Room) WebSocketReader(Nickname string) {
	for {
		_, message, err := r.Players[Nickname].Connection.ReadMessage()
		// fmt.Println(Nickname)
		if err != nil {
			fmt.Println("Erro sdfsdr from user role 0 with Token '" + r.Players[Nickname].Token + "': '" + err.Error() + "'.")
			_, stillOpen := <-r.Recovery.Player_IsAvailableRead[Nickname]
			if !stillOpen {
				close(r.Messenger.Player_From[Nickname])
				break
			}
		} else {
			var m game_logic.InputMessage
			json.Unmarshal(message, &m)
			r.Messenger.Player_From[Nickname] <- m
		}
	}
	log.Print("WebSocketReader room = " + r.CurrentRoomId.String() + ", Nickname = " + Nickname + " correctly completed.")
	return
}

// Обслуживающая горутина To
func (r *Room) WebSocketWriter(Nickname string) {
	for _, p := range r.Players {
		if p.Login == Nickname {
		BreakDance:
			for {
				select {
				case message := <-r.Messenger.Player_To[p.Login]:
					for _, pl := range r.Players {
						if pl.Login != Nickname {
							err := pl.Connection.WriteJSON(&message)
							if err != nil {
								_, stillOpen := <-r.Recovery.Player_IsAvailableWrite[p.Login]
								if !stillOpen {
									break BreakDance
								}
							}
						}
					}

				}
			}
			_ = r.Players[p.Login].Connection.Close()
		}
	}
	fmt.Println("WebSocketWriter room = " + r.CurrentRoomId.String() + ", role = " + Nickname + " correctly completed.")
	return
}
