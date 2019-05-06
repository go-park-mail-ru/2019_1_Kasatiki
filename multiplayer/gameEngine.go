package main

import (
	// "fmt"
)

type mes struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (r *Room) GameEngine() {
	// var buf mes
	var message1 mes
	var message2 mes
	for {
		select {
		case message1 = <-r.Messenger.Player_1_From:
			// message1.X = buf.X
			// message1.Y = buf.Y
		case message2 = <-r.Messenger.Player_2_From:
			// message2.X = buf.X
			// message2.Y = buf.Y
		}
		// if message2.X != 0 && message2.Y != 0 {
		// 	r.Messenger.Player_1_To <- message2
		// }

		// if message1.X != 0 && message1.Y != 0 {
		// 	r.Messenger.Player_2_To <- message1
		// }
	}

	return
}
