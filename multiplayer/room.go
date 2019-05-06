package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

// Пока не учитывал игровую логику
// Здесь чисто игроки и каналы общения

type Room struct {
	Player_1 *UserConnection
	Player_2 *UserConnection

	// Состояния игры.
	Player_1_UploadedCharacters bool
	Player_2_UploadedCharacters bool
	UserTurnNumber              PlayerId

	// Каналы для устранения дисконекта
	Recovery struct {
		Player_1_IsAvailableRead  chan struct{}
		Player_1_IsAvailableWrite chan struct{}
		Player_2_IsAvailableRead  chan struct{}
		Player_2_IsAvailableWrite chan struct{}
	}

	// Каналы, c помощью которых go room.GameMaster() общается с внешним миром
	Messenger struct {
		Player_1_From chan mes
		Player_1_To   chan mes
		Player_2_From chan mes
		Player_2_To   chan mes
	}

	// Созданную комнату необходимо будет отправить в канал комнат которые мы собираемся уничтожить
	Suicide chan RoomId

	// Номер текущей комнаты
	CurrentRoomId RoomId
}

// Создание новой комнаты
func NewRoom(player1, player2 *UserConnection, completedRooms chan RoomId, ownNumber RoomId) (room *Room) {
	// Инициализируем структуру комнаты
	room = &Room{
		Player_1:      player1,
		Player_2:      player2,
		Suicide:       completedRooms,
		CurrentRoomId: ownNumber,
	}

	// Инициализируем каналы
	room.Messenger.Player_1_From = make(chan mes, 5)
	room.Messenger.Player_1_To = make(chan mes, 5)
	room.Messenger.Player_2_From = make(chan mes, 5)
	room.Messenger.Player_2_To = make(chan mes, 5)

	room.Recovery.Player_1_IsAvailableRead = make(chan struct{}, 1)
	room.Recovery.Player_1_IsAvailableWrite = make(chan struct{}, 1)
	room.Recovery.Player_2_IsAvailableRead = make(chan struct{}, 1)
	room.Recovery.Player_2_IsAvailableWrite = make(chan struct{}, 1)

	//	Какова работа логики комнаты?
	// Внутри каждая комната управляется одним GameEngine - горутиной.

	// 4 обслуживающие горутины нужны для изоляции соединения от игровой логики.
	//    	╭─Player_1_From─>>─╮      ╭─<<─Player_2_From─╮
	// Player_1         	  GameEngine            	Player_2
	//    	╰─Player_1_To───<<─╯      ╰─>>─Player_2_To───╯

	go room.WebSocketReader(0)
	go room.WebSocketWriter(0)
	go room.WebSocketReader(1)
	go room.WebSocketWriter(1)
	//go room.GameEngine()

	fmt.Println("Room created with Player_1 = '%s', Player_2 = '%s'", room.Player_1.Token, room.Player_2.Token)
	return
}

// Закрываем лавочку и горутины
func (r *Room) StopRoom() {
	close(r.Recovery.Player_1_IsAvailableRead)
	close(r.Messenger.Player_1_To)
	close(r.Recovery.Player_1_IsAvailableWrite)

	close(r.Recovery.Player_2_IsAvailableRead)
	close(r.Messenger.Player_2_To)
	close(r.Recovery.Player_2_IsAvailableWrite)

	fmt.Println("room with Player_1.Token='" + r.Player_1.Token + "', r.Player_2.Token='" + r.Player_2.Token + "' closed")
	return
}

// Отпраление текущего id комнаты в канал "для удаления"
func (r *Room) RemoveRoom() {
	r.Suicide <- r.CurrentRoomId
	return
}

// Восстановление соединения
func (r *Room) Reconnect(user *UserConnection, role PlayerId) {
	fmt.Println("Reconnect sessioni = '%s' as role %d", user.Token, role)

	if role == 0 {
		if r.Player_1 != nil {
			_ = r.Player_1.Connection.Close()
		}
		r.Player_1 = user
		select {
		case r.Recovery.Player_1_IsAvailableRead <- struct{}{}:
		default:
		}
		select {
		case r.Recovery.Player_1_IsAvailableWrite <- struct{}{}:
		default:
		}
	} else {
		if r.Player_2 != nil {
			_ = r.Player_2.Connection.Close()
		}
		r.Player_2 = user
		select {
		case r.Recovery.Player_2_IsAvailableRead <- struct{}{}:
		default:
		}
		select {
		case r.Recovery.Player_2_IsAvailableWrite <- struct{}{}:
		default:
		}
	}
	return
}

// Обслуживающая горутина From
func (r *Room) WebSocketReader(role PlayerId) {
	if role == 0 {
		for {
			_, message, err := r.Player_1.Connection.ReadMessage()
			if err != nil {
				fmt.Println("Error from user role 0 with Token '" + r.Player_1.Token + "': '" + err.Error() + "'.")
				_, stillOpen := <-r.Recovery.Player_1_IsAvailableRead
				if !stillOpen {
					close(r.Messenger.Player_1_From)
					break
				}
			} else {
				log.Print("message from user role 0 with Token '" + r.Player_1.Token + "': '" + string(message) + "'.")
				var m mes
				m.message = string(message)
				r.Messenger.Player_1_From <- m
			}
		}
	} else {
		for {
			_, message, err := r.Player_2.Connection.ReadMessage()
			if err != nil {
				fmt.Println("Error from user role 1 with Token '" + r.Player_1.Token + "': '" + err.Error() + "'.")
				_, stillOpen := <-r.Recovery.Player_2_IsAvailableRead
				if !stillOpen {
					close(r.Messenger.Player_2_From)
					break
				}
			} else {
				log.Print("message from user role 1 with Token '" + r.Player_1.Token + "': '" + string(message) + "'.")
				var m mes
				m.message = string(message)
				r.Messenger.Player_2_From <- m
			}
		}
	}
	log.Print("WebSocketReader room = " + r.CurrentRoomId.String() + ", role = " + role.String() + " correctly completed.")
	return
}

// Обслуживающая горутина To
func (r *Room) WebSocketWriter(role PlayerId) {
	fmt.Println("z nen")
	if role == 0 {
	MessageSending1:
		for message := range r.Messenger.Player_1_To {
			for {
				w, err := r.Player_1.Connection.NextWriter(websocket.TextMessage)
				if err != nil {
					fmt.Println(err)
					_, stillOpen := <-r.Recovery.Player_1_IsAvailableWrite
					if !stillOpen {
						break MessageSending1
					}
					_ = json.NewEncoder(w).Encode(&message)
				} else {
					break
				}
			}
		}
		_ = r.Player_1.Connection.Close()
	} else {
	MessageSending2:
		for message := range r.Messenger.Player_2_To {
			for {
				w, err := r.Player_2.Connection.NextWriter(websocket.TextMessage)
				if err != nil {
					_, stillOpen := <-r.Recovery.Player_2_IsAvailableWrite
					if !stillOpen {
						break MessageSending2
					}
					_ = json.NewEncoder(w).Encode(&message)
				} else {
					break
				}
			}
		}
		_ = r.Player_2.Connection.Close()
	}
	fmt.Println("WebSocketWriter room = " + r.CurrentRoomId.String() + ", role = " + role.String() + " correctly completed.")
	return
}
