package kickgo

import "fmt"

const (
	SubscribeEventName              = "pusher:subscribe"
	Connection_establishedEventName = "pusher:connection_established"
	Subscription_succeededEventName = "pusher_internal:subscription_succeeded"
	ChatMessageSentEventName        = "App\\Events\\ChatMessageSentEvent"
	ChatMessageDeletedEventName     = "App\\Events\\ChatMessageDeletedEvent"
	ChatMessageReactEventName       = "App\\Events\\ChatMessageReact"
)

type EventHandler interface {
	Type() string
	Handle(*Kick, interface{})
}

type eventHandlerInstance struct {
	eventHandler EventHandler
}

type eventHandlerInterface interface {
	Type() string
	New() interface{}
}

// ChatMessageSentEvent
type chatMessageSentEventHandler func(*Kick, *ChatMessageSentEvent)

func (eh chatMessageSentEventHandler) Type() string {
	return ChatMessageSentEventName
}

func (eh chatMessageSentEventHandler) New() interface{} {
	return &ChatMessageSentEvent{}
}

func (eh chatMessageSentEventHandler) Handle(k *Kick, i interface{}) {
	x, err := bytesToTPtr(i.([]byte), &ChatMessageSentEvent{})

	if err == nil {
		eh(k, x)
	} else {
		fmt.Println(err)
	}
}

// ChatMessageDeletedEvent
type chatMessageDeletedEventHandler func(*Kick, *ChatMessageDeletedEvent)

func (eh chatMessageDeletedEventHandler) Type() string {
	return ChatMessageDeletedEventName
}

func (eh chatMessageDeletedEventHandler) New() interface{} {
	return &ChatMessageDeletedEvent{}
}

func (eh chatMessageDeletedEventHandler) Handle(k *Kick, i interface{}) {
	x, err := bytesToTPtr(i.([]byte), &ChatMessageDeletedEvent{})
	if err == nil {
		eh(k, x)
	}
}

// ChatMessageReact
type chatMessageReactEventHandler func(*Kick, *ChatMessageReact)

func (eh chatMessageReactEventHandler) Type() string {
	return ChatMessageReactEventName
}

func (eh chatMessageReactEventHandler) New() interface{} {
	return &ChatMessageReact{}
}

func (eh chatMessageReactEventHandler) Handle(k *Kick, i interface{}) {
	x, err := bytesToTPtr(i.([]byte), &ChatMessageReact{})
	if err == nil {
		eh(k, x)
	}
}

// Subscription_succeeded
type subscription_succeededEventHandler func(*Kick, *Subscription_succeeded)

func (eh subscription_succeededEventHandler) Type() string {
	return Subscription_succeededEventName
}

func (eh subscription_succeededEventHandler) New() interface{} {
	return &Subscription_succeeded{}
}

func (eh subscription_succeededEventHandler) Handle(k *Kick, i interface{}) {
	x, err := bytesToTPtr(i.([]byte), &Subscription_succeeded{})
	if err == nil {
		eh(k, x)
	}
}

// Connection_established
type connection_establishedEventHandler func(*Kick, *Connection_established)

func (eh connection_establishedEventHandler) Type() string {
	return Connection_establishedEventName
}

func (eh connection_establishedEventHandler) New() interface{} {
	return &Connection_established{}
}

func (eh connection_establishedEventHandler) Handle(k *Kick, i interface{}) {
	x, err := bytesToTPtr(i.([]byte), &Connection_established{})
	if err == nil {
		eh(k, x)
	}
}

// Subscribe
type subscribeEventHandler func(*Kick, *Subscribe)

func (eh subscribeEventHandler) Type() string {
	return SubscribeEventName
}

func (eh subscribeEventHandler) New() interface{} {
	return &Subscribe{}
}

func (eh subscribeEventHandler) Handle(k *Kick, i interface{}) {
	x, err := bytesToTPtr(i.([]byte), &Subscribe{})
	if err == nil {
		eh(k, x)
	}
}

func handlerForInterface(handler interface{}) EventHandler {
	switch v := handler.(type) {
	case func(*Kick, *Subscribe):
		return subscribeEventHandler(v)
	case func(*Kick, *Connection_established):
		return connection_establishedEventHandler(v)
	case func(*Kick, *Subscription_succeeded):
		return subscription_succeededEventHandler(v)
	case func(*Kick, *ChatMessageSentEvent):
		return chatMessageSentEventHandler(v)
	case func(*Kick, *ChatMessageDeletedEvent):
		return chatMessageDeletedEventHandler(v)
	case func(*Kick, *ChatMessageReact):
		return chatMessageReactEventHandler(v)
	}
	return nil
}
