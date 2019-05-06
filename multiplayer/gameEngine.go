package main

type mes struct {
	message string `json: message`
}

func (r *Room) GameEngine() {
	var message mes
	for {
		select {
		case message = <-r.Messenger.Player_1_From:
			message.message = "By player 1" + message.message
		case message = <-r.Messenger.Player_2_From:
			message.message = "By player 2" + message.message
		}
		if message.message != "" {
			r.Messenger.Player_1_To <- message
			r.Messenger.Player_2_To <- message
			message.message = ""
		}
	}

	return
}
