package main

import (
	"fmt"
)

func (r *Room) GameEngine() {
	var message []byte
	for {
		select {
		case message = <-r.Messenger.Player_1_From:
			message = append([]byte("message came from the Player_1: "), message...)
			fmt.Println(string(message))
		case message = <-r.Messenger.Player_2_From:
			message = append([]byte("message came from the Player_2: "), message...)
		}
		if message != nil {
			r.Messenger.Player_1_To <- message
			r.Messenger.Player_2_To <- message
			message = nil
		}
	}

	return
}
