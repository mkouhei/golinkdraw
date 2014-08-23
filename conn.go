/*
 github.com/mkouhei/golinkdraw/conn.go

 Copyright (c) 2014 Kouhei Maeda <mkouhei@palmtb.net>

 This software is release under the Expat License.

 This source code is a modification of the source code of
 https://github.com/gorilla/websocket/examples/filewatch/main.go

 The original source code is licensed under the 2-Clause BSD License,
 and copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
*/
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll response for chages with this period.
	pollPeriod = 1 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	defaultWidth = 400

	defaultHeight = 400
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

// serverWs handles webocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	go c.writePump()
	c.readPump()
}

func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

func (c *connection) writePump() {
	pingTicker := time.NewTicker(pingPeriod)
	pollTicker := time.NewTicker(pollPeriod)
	defer func() {
		pingTicker.Stop()
		pollTicker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			<-pollTicker.C
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			type Response struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			}
			data := &Response{}
			if err := json.Unmarshal(message, &data); err != nil {
				//log.Println(err)
				continue
			}
			strSVG := StringSVG(data.Width, data.Height)
			if err := c.write(websocket.TextMessage, []byte(strSVG)); err != nil {
				return
			}
		case <-pollTicker.C:
			strSVG := StringSVG(defaultWidth, defaultHeight)
			if err := c.write(websocket.TextMessage, []byte(strSVG)); err != nil {
				return
			}
		case <-pingTicker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *connection) readPump() {
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		h.broadcast <- message
	}
}
