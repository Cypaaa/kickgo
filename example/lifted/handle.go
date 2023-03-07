package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Cypaaa/kickgo"
)

func subscribed(k *kickgo.Kick, _ *kickgo.Subscription_succeeded) {
	log.Println("Listening for events in the chatroom(s) ", k.ChatroomIds)
}

func messageSent(k *kickgo.Kick, e *kickgo.ChatMessageSentEvent) {
	if e.Message.Action == "" {
		SUBCOUNT++
		fmt.Println(SUBCOUNT)
	}
}

func SendLoop() {
	old := SUBCOUNT
	for {
		time.Sleep(50 * time.Millisecond)
		for i, conn := range WSCONN {
			if SUBCOUNT > old {
				err := conn.WriteMessage(1, []byte(fmt.Sprint(SUBCOUNT+SUBGAP)))
				if err != nil {
					log.Println(err)
					conn.Close()
					WSCONN = append(WSCONN[:i], WSCONN[i+1:]...)
				}
				old = SUBCOUNT
			}
		}
	}
}

func wsHandler(k *kickgo.Kick, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	WSCONN = append(WSCONN, ws)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form["submit"] != nil {
		uintVal, err := strconv.ParseUint(r.Form["gap"][0], 10, 64)
		if err == nil {
			SUBGAP = uintVal
		}
		http.Redirect(w, r, "/", 302)
	} else if SUBGAP == 99999999999 {
		w.Write([]byte("<html><body><form method=\"post\"><label for=\"gap\">Current sub count</label><input name=\"gap\" id=\"gap\" type=\"number\" value=\"0\"><br><input name=\"submit\" type=\"submit\" value=\"Valid\"></form></body></html>"))
	} else {
		w.Write([]byte("<h1 id=\"subcount\">" + strconv.FormatUint(SUBGAP, 10) + "</h1>"))
		w.Write([]byte("<script defer>const subcount = document.getElementById(\"subcount\");\nlet socket = new WebSocket(\"ws://localhost:9898/ws\");\nsocket.onmessage = (e) => { subcount.innerHTML = e.data; console.log(e.data); };</script>"))
	}
}
