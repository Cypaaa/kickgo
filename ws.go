package kickgo

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type WS struct {
	conn   *websocket.Conn
	url    url.URL
	status bool
}

// Makes a new ws client
func newWS() *WS {
	return &WS{
		url: url.URL{
			Scheme:   endpointWSScheme,
			Host:     endpointWSHost,
			Path:     endpointWSPath,
			RawQuery: endpointWSRawQuery,
		},
		conn:   &websocket.Conn{},
		status: false,
	}
}

// Returns if the conn is opened or closed
// true if opened
// false if closed
func (ws *WS) Status() bool {
	return ws.status
}

// Disconnect our client from the ws server
func (ws *WS) disconnect() {
	ws.status = false
	err := ws.conn.Close()
	if err != nil {
		log.Fatal("dial:", err)
	}
}

// returns the client connection to the server
func (ws *WS) connect(header http.Header) {
	c, _, err := websocket.DefaultDialer.Dial(ws.url.String(), header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	ws.conn = c
	ws.status = true
}

// loop used to listen for the server messages
func (ws *WS) listen(msgHandler func([]byte)) {
	for {
		_, m, err := ws.conn.ReadMessage()
		if err != nil {
			// not Fatal bcs Listen must (defer)close the conn
			log.Println("error:", err)
			return
		}
		msgHandler(m)
	}
}

// send json message to the server
func (ws *WS) Send(v interface{}) error {
	return ws.conn.WriteJSON(v)
}
