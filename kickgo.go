package kickgo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Kick struct {
	ws          *WS
	ChatroomIds []string
	handlers    map[string][]*eventHandlerInstance
	token       string
	Custom      map[string]any // To store custom gobal data, bcs Kick obj is sent everywhere
}

// Returns a pointer to a new Kick object
func New() *Kick {
	var k *Kick = &Kick{
		ws: newWS(),
	}
	k.defaultHandlers()
	k.Custom = map[string]any{}
	return k
}

// Send a message to our ws client, like to subscribe to a channel
func (k *Kick) send(v interface{}) {
	k.ws.Send(v)
}

func (k *Kick) recv(data []byte) {
	if e, err := bytesToTPtr(data, &ChatEvent{}); err == nil {
		// if there are/is handler(s) for our event
		if ehs, ok := k.handlers[e.Event]; ok {
			// for each, execute it
			for _, eh := range ehs {
				eh.eventHandler.Handle(k, []byte(e.Data))
			}
		}
	}
}

func (k *Kick) Listen(token string) {
	k.token = token
	k.ws.connect(nil)
	defer k.ws.disconnect()
	k.AddHandler(k.subscribeHandler)
	k.ws.listen(k.recv)
}

// Tells our ws client to subscribe to a certain channel to listen for its events
func (k *Kick) subscribe(chatroomid string) {
	k.send(
		Subscribe{
			Event: SubscribeEventName,
			Data: SubscribeData{
				Auth:    "",
				Channel: "chatrooms." + chatroomid,
			},
		},
	)
}

// set our map of handlers to its default value if its value is nil
func (k *Kick) defaultHandlers() {
	if k.handlers == nil {
		k.handlers = map[string][]*eventHandlerInstance{}
	}
}

// Add an handler for the event related to the required struct
// ex func(*Kick, *Connection_established) ->
// execute this function is the event "pusher:connection_established" is received
func (k *Kick) AddHandler(handler interface{}) {
	eh := handlerForInterface(handler)
	if eh == nil {
		log.Println("Invalid handler received, no event can call this handler.")
	}
	k.addEventHandler(eh)
}

// internal function to add handler
func (k *Kick) addEventHandler(eventHandler EventHandler) {
	k.defaultHandlers()
	ehi := &eventHandlerInstance{eventHandler}
	k.handlers[eventHandler.Type()] = append(k.handlers[eventHandler.Type()], ehi)
}

// Subscribe to our chatrooms at the begining of the listening
func (k *Kick) subscribeHandler(*Kick, *Connection_established) {
	for _, chatroomid := range k.ChatroomIds {
		k.subscribe(chatroomid)
	}
}

// subscribe to a chatroom after the lcient has started
func (k *Kick) SubscribeOnTheFly(chatroomid string) {
	if !k.ws.Status() {
		log.Panicln("Cannot register a chatroom on the fly if no connection has been made.")
	}
	k.ChatroomIds = append(k.ChatroomIds, chatroomid)
	k.subscribe(chatroomid)
}

func (k *Kick) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		req.Header.Set("Authorization", "Bearer "+k.token)
		res, err := http.DefaultClient.Do(req)
		if err == nil {
			return io.ReadAll(res.Body)
		}
	}
	return nil, err
}

func (k *Kick) GetChannelByName(cname string) (*Channel, error) {
	var channel *Channel = &Channel{}
	data, err := k.get(fmt.Sprintf(endpointLivestream, cname))

	if err == nil {
		err = json.Unmarshal(data, channel)
		if err == nil {
			return channel, nil
		}
	}
	log.Println(err.Error())
	return nil, err
}

func (k *Kick) AddChatroom(cname string) {
	channel, err := k.GetChannelByName(cname)
	if err == nil {
		k.ChatroomIds = append(k.ChatroomIds, channel.Chatroom.ChatableId)
	}
}

func (k *Kick) AddChatroomId(cid string) {
	k.ChatroomIds = append(k.ChatroomIds, cid)
}
