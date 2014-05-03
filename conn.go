package main

import (
	"github.com/gorilla/websocket"
	"github.com/mkouhei/golinkdraw/modules"
	"log"
	"net/http"
	"time"
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
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// serverWs handles webocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	go writer(ws)
	reader(ws)
}

func writer(ws *websocket.Conn) {
	pingTicker := time.NewTicker(pingPeriod)
	pollTicker := time.NewTicker(pollPeriod)
	defer func() {
		pingTicker.Stop()
		pollTicker.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-pollTicker.C:
			if err := ws.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Println(err)
			}

			// generate SVG string
			strSVG := StringSVG(modules.RenderingSVG)

			if err := ws.WriteMessage(websocket.TextMessage, []byte(strSVG)); err != nil {
				log.Println(err)
				return
			}

		case <-pingTicker.C:

			if err := ws.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Println(err)
			}

			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println(err)
				return
			}

		}
	}
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	if err := ws.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
	}
	ws.SetPongHandler(func(string) error {
		if err := ws.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.Println(err)
		}
		return nil
	})
	for {
		if mt, _, err := ws.ReadMessage(); err != nil {
			log.Printf("messageType: %d; message: %s", mt, err)
			break
		}
	}
}
