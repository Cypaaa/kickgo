package kickgo

type ChatEvent struct {
	Event   string // e.g "App\\Events\\ChatMessageSentEvent"
	Data    string // e.g ChatMessageSentEvent
	Channel string // "chatroom"
}

type ChatMessageSentEvent struct {
	Message *ChatMessage `json:"message"`
	User    *ChatUser    `json:"user"`
}

type Connection_established struct {
	SocketId        string `json:"socket_id"`
	ActivityTimeout int    `json:"activity_timeout"`
}
type Subscription_succeeded struct{}

type ChatMessageReact struct {
	MessageId  string `json:"message_id"`
	Reaction   string `json:"reaction"`
	ChatroomId string `json:"chatroom_id"`
	UserId     string `json:"user_id"`
}

type ChatMessageDeletedEvent struct {
	Id         string `json:"id"`
	DeletedBy  string `json:"deleted_by"`
	ChatroomId string `json:"chatroom_id"`
}

type Subscribe struct {
	Event string        `json:"event"`
	Data  SubscribeData `json:"data"`
}

type SubscribeData struct {
	Auth    string `json:"auth"`
	Channel string `json:"channel"`
}

type ChatMessage struct {
	ID                string `json:"id"`
	Message           string `json:"message"`
	Type              string `json:"type"`
	RepliedTo         string `json:"replied_to"`
	IsInfo            bool   `json:"is_info"`
	LinkPreview       string `json:"link_preview"`
	ChatroomId        string `json:"chatroom_id"`
	Role              string `json:"role"`
	CreatedAt         int64  `json:"created_at"`
	Action            string `json:"action"`
	OptionalMessage   string `json:"optional_message"`
	MonthsSubscribed  uint16 `json:"months_subscribed"`
	SubscriptionCount uint16 `json:"subscription_count"`
	GiftedUsers       string `json:"giftedUsers"`
}

type ChatUser struct {
	ID               int64    `json:"id"`
	Username         string   `json:"username"`
	Role             string   `json:"role"`
	IsSuperAdmin     bool     `json:"isSuperAdmin"`
	ProfileThumb     string   `json:"profile_thumb"`
	Verified         bool     `json:"verified"`
	FollowerBadges   []string `json:"follower_badges"`
	IsSubscribed     bool     `json:"is_subscribed"`
	IsFounder        bool     `json:"is_founder"`
	MonthsSubscribed uint16   `json:"months_subscribed"`
	QuantityGifted   uint16   `json:"quantity_gifted"`
}
