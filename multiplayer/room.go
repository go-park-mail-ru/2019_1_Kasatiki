package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// Пока не учитывал игровую логику
// Здесь чисто игроки и каналы общения

type Room struct {
	Players map[string]*UserConnection
	//Player_1 *UserConnection
	//Player_2 *UserConnection

	// Состояния игры.
	//Player_1_UploadedCharacters bool
	//Player_2_UploadedCharacters bool

	// Каналы для устранения дисконекта
	Recovery struct {
		Player_IsAvailableRead  map[string]chan struct{}
		Player_IsAvailableWrite map[string]chan struct{}
		//Player_1_IsAvailableRead  chan struct{}
		//Player_1_IsAvailableWrite chan struct{}
		//Player_2_IsAvailableRead  chan struct{}
		//Player_2_IsAvailableWrite chan struct{}
	}

	// Каналы, c помощью которых go room.GameMaster() общается с внешним миром
	Messenger struct {
		Player_From map[string]chan mes
		Player_To   map[string]chan mes
		//Player_1_From chan mes
		//Player_1_To   chan mes
		//Player_2_From chan mes
		//Player_2_To   chan mes
	}

	// Созданную комнату необходимо будет отправить в канал комнат которые мы собираемся уничтожить
	Suicide chan RoomId

	// Номер текущей комнаты
	CurrentRoomId RoomId
}

// Создание новой комнаты
func NewRoom(players []*UserConnection, completedRooms chan RoomId, ownNumber RoomId) (room *Room) {
	// Инициализируем структуру комнаты
	room = &Room{
		Suicide:       completedRooms,
		CurrentRoomId: ownNumber,
	}

	// Создаем мапы
	room.Players = make(map[string]*UserConnection)
	room.Messenger.Player_From = make(map[string]chan mes)                 //make(chan mes, 5)
	room.Messenger.Player_To = make(map[string]chan mes)                   //make(chan mes, 5)
	room.Recovery.Player_IsAvailableRead = make(map[string]chan struct{})  //make(chan struct{}, 1)
	room.Recovery.Player_IsAvailableWrite = make(map[string]chan struct{}) //make(chan struct{}, 1)

	for _, p := range players {
		room.Players[p.Login] = p
		room.Messenger.Player_From[p.Login] = make(chan mes, 5)
		room.Messenger.Player_To[p.Login] = make(chan mes, 5)
		room.Recovery.Player_IsAvailableRead[p.Login] = make(chan struct{}, 1)
		room.Recovery.Player_IsAvailableWrite[p.Login] = make(chan struct{}, 1)
		go room.WebSocketReader(p.Login)
		go room.WebSocketWriter(p.Login)
	}

	go room.GameEngine()

	// Инициализируем каналы
	//room.Messenger.Player_1_From = make(chan mes, 5)
	//room.Messenger.Player_1_To = make(chan mes, 5)
	//room.Messenger.Player_2_From = make(chan mes, 5)
	//room.Messenger.Player_2_To = make(chan mes, 5)

	//room.Recovery.Player_1_IsAvailableRead = make(chan struct{}, 1)
	//room.Recovery.Player_1_IsAvailableWrite = make(chan struct{}, 1)
	//room.Recovery.Player_2_IsAvailableRead = make(chan struct{}, 1)
	//room.Recovery.Player_2_IsAvailableWrite = make(chan struct{}, 1)

	//	Какова работа логики комнаты?
	// Внутри каждая комната управляется одним GameEngine - горутиной.

	// 4 обслуживающие горутины нужны для изоляции соединения от игровой логики.
	//    	╭─Player_1_From─>>─╮      ╭─<<─Player_2_From─╮
	// Player_1         	  GameEngine            	Player_2
	//    	╰─Player_1_To───<<─╯      ╰─>>─Player_2_To───╯

	//go room.WebSocketReader(0)
	//go room.WebSocketWriter(0)
	//go room.WebSocketReader(1)
	//go room.WebSocketWriter(1)
	//go room.GameEngine()
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
	//fmt.Println("room with Player_1.Token='" + r.Player_1.Token + "', r.Player_2.Token='" + r.Player_2.Token + "' closed")
	return
}

// Отпраление текущего id комнаты в канал "для удаления"
func (r *Room) RemoveRoom() {
	r.Suicide <- r.CurrentRoomId
	return
}

// Восстановление соединения
func (r *Room) Reconnect(user *UserConnection) {
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
	//if role == 0 {
	//
	//} else {
	//	if r.Player_2 != nil {
	//		_ = r.Player_2.Connection.Close()
	//	}
	//	r.Player_2 = user
	//	select {
	//	case r.Recovery.Player_2_IsAvailableRead <- struct{}{}:
	//	default:
	//	}
	//	select {
	//	case r.Recovery.Player_2_IsAvailableWrite <- struct{}{}:
	//	default:
	//	}
	//}
	return
}

// Обслуживающая горутина From
func (r *Room) WebSocketReader(Nickname string) {
	//for _, p := range r.Players {
	//	if p.Login == Nickname {
	for {
		_, message, err := r.Players[Nickname].Connection.ReadMessage()
		fmt.Println(Nickname)
		if err != nil {
			fmt.Println("Erro sdfsdr from user role 0 with Token '" + r.Players[Nickname].Token + "': '" + err.Error() + "'.")
			_, stillOpen := <-r.Recovery.Player_IsAvailableRead[Nickname]
			if !stillOpen {
				close(r.Messenger.Player_From[Nickname])
				break
			}
		} else {
			var m mes
			json.Unmarshal(message, &m)
			r.Messenger.Player_From[Nickname] <- m
		}
	}
	//}
	//}

	//if role == 0 {
	//	for {
	//		_, message, err := r.Player_1.Connection.ReadMessage()
	//
	//		if err != nil {
	//			fmt.Println("Erro sdfsdr from user role 0 with Token '" + r.Player_1.Token + "': '" + err.Error() + "'.")
	//			_, stillOpen := <-r.Recovery.Player_1_IsAvailableRead
	//			if !stillOpen {
	//				close(r.Messenger.Player_1_From)
	//				break
	//			}
	//		} else {
	//			var m mes
	//			json.Unmarshal(message, &m)
	//			r.Messenger.Player_1_From <- m
	//		}
	//	}
	//} else {
	//	for {
	//		_, message, err := r.Player_2.Connection.ReadMessage()
	//
	//		if err != nil {
	//			fmt.Println("Error from user role 1 with Token '" + r.Player_1.Token + "': '" + err.Error() + "'.")
	//			_, stillOpen := <-r.Recovery.Player_2_IsAvailableRead
	//			if !stillOpen {
	//				close(r.Messenger.Player_2_From)
	//				break
	//			}
	//		} else {
	//			var m mes
	//			json.Unmarshal(message, &m)
	//			r.Messenger.Player_2_From <- m
	//		}
	//	}
	//}
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
	//if Nickname == 0 {
	//
	//	_ = r.Player_1.Connection.Close()
	//} else {
	//	for {
	//		select {
	//		case message := <-r.Messenger.Player_1_To:
	//			err := r.Player_1.Connection.WriteJSON(&message)
	//			if err != nil {
	//				_, stillOpen := <-r.Recovery.Player_2_IsAvailableWrite
	//				if !stillOpen {
	//					break
	//				}
	//			}
	//		}
	//	}
	//	_ = r.Player_2.Connection.Close()
	//}
	fmt.Println("WebSocketWriter room = " + r.CurrentRoomId.String() + ", role = " + Nickname + " correctly completed.")
	return
}
