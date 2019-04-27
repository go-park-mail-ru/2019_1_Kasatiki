// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"2019_1_Kasatiki/pkg/middleware"
	"2019_1_Kasatiki/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	//"2019_1_Kasatiki/pkg/dbhandler"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send     chan models.Message
	Nickname string
	ImgUrl   string
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// Todo:
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var msg models.Message
		msg.Body = string(message)
		msg.Nickname = c.Nickname
		msg.Edited = false
		t := time.Now()
		msg.Timestamp = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
		msg.Imgurl = c.ImgUrl
		err = c.hub.DB.InsertMessage(msg)
		if err != nil {
			fmt.Println(err)
		}
		c.hub.broadcast <- msg
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			//w.Write()
			_ = json.NewEncoder(w).Encode(&message)

			//w.Write()
			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				//w.Write(<-c.conn.ReadJSON(c.send))
				json.NewEncoder(w).Encode(c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	nickname := "anonymous"
	imgurl := ""
	cookie, err := r.Cookie("session_id")
	if err != nil {
		client := &Client{hub: hub, conn: conn, send: make(chan models.Message, 256), ImgUrl: imgurl, Nickname: nickname}
		client.hub.register <- client
		go client.writePump()
		go client.readPump()
		return
	}
	claims, err := middleware.CheckAuth(cookie)
	if err != nil {
		client := &Client{hub: hub, conn: conn, send: make(chan models.Message, 256), ImgUrl: imgurl, Nickname: nickname}
		client.hub.register <- client
		go client.writePump()
		go client.readPump()
		return
	}
	id := int(claims["id"].(float64))
	user, err := hub.DB.GetUser(id)
	client := &Client{hub: hub, conn: conn, send: make(chan models.Message, 256), ImgUrl: user.ImgUrl, Nickname: user.Nickname}
	client.hub.register <- client
	go client.writePump()
	go client.readPump()
}
