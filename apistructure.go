package kickgo

type Channel struct {
	ID                  string             `json:"id"`
	UserId              string             `json:"user_id"`
	Slug                string             `json:"slug"`
	PlaybackURL         string             `json:"playback_url"`
	NameUpdatedAt       string             `json:"name_updated_at"`
	FollowerCount       int64              `json:"follower_count"`
	VODEnabled          bool               `json:"vod_enabled"`
	SubscriptionEnabled bool               `json:"subscription_enabled"`
	CanHost             bool               `json:"can_host"`
	SubscriberBadges    []SubscriberBadge  `json:"subscriber_badges"`
	BannerImage         Banner             `json:"banner_image"`
	RecentCategories    []Category         `json:"recent_categories"`
	Role                string             `json:"role"` // "owner" or "moderator"
	Muted               bool               `json:"muted"`
	FollowerBadges      []FollowerBadge    `json:"follower_badges"`
	Livestream          Livestream         `json:"livestream"`
	User                User               `json:"user"`
	Chatroom            Chatroom           `json:"chatroom"`
	AscendingLinks      []AscendingLink    `json:"ascending_links"`
	Plan                Plan               `json:"plan"`
	PreviousLivestream  PreviousLivestream `json:"previous_livestream"`
	Verified            Verified           `json:"verified"`
}

type SubscriberBadge struct {
	ID         string `json:"id"`
	ChannelId  string `json:"channel_id"`
	Months     int64  `json:"months"`
	BadgeImage Badge  `json:"badge_image"`
}

type Badge struct {
	Srcset string `json:"srcset"`
	Src    string `json:"src"`
}

type Banner struct {
	Responsive string `json:"responsive"`
	URL        string `json:"url"`
}

type Category struct {
	ID          string       `json:"id"`
	CategoryId  string       `json:"category_id"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	Tags        []string     `json:"tags"`
	Description string       `json:"description"`
	DeletedAt   string       `json:"deleted_at"`
	Banner      Banner       `json:"banner"`
	Category    CategoryInfo `json:"category"`
}

type CategoryInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Icon string `json:"icon"`
}

type FollowerBadge struct {
	ID int64 `json:"id"`
}

type Livestream struct {
	ID            string     `json:"id"`
	Slug          string     `json:"slug"`
	ChannelId     string     `json:"channel_id"`
	CreatedAt     string     `json:"created_at"`
	SessionTitle  string     `json:"session_title"`
	IsLive        bool       `json:"is_live"`
	RiskLevelId   string     `json:"risk_level_id"`
	Source        string     `json:"source"`
	TwitchChannel string     `json:"twitch_channel"`
	Language      string     `json:"language"`
	Duration      int64      `json:"duration"`
	IsMature      bool       `json:"is_mature"`
	ViewerCount   int64      `json:"viewer_count"` // Deprecated I guess
	Viewers       int64      `json:"viewers"`
	Thumbnail     Thumbnail  `json:"thumbnail"`
	Categories    []Category `json:"categories"`
	Tags          []string   `json:"tags"`
}

type Thumbnail struct {
	Responsive string `json:"responsive"`
	URL        string `json:"url"`
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	AgreedToTerms bool   `json:"agreed_to_terms"`
	Bio           string `json:"bio"`
	Contry        string `json:"country"`
	State         string `json:"state"`
	City          string `json:"city"`
	Instagram     string `json:"instagram"`
	Twitter       string `json:"twitter"`
	Youtube       string `json:"youtube"`
	Discord       string `json:"discord"`
	Tiktok        string `json:"tiktok"`
	Facebook      string `json:"facebook"`
	ProfilePic    string `json:"profile_pic"`
}

type Chatroom struct {
	ID           string `json:"id"`
	ChannelId    string `json:"channel_id"`
	ChatableId   string `json:"chatable_id"`
	SlowMode     bool   `json:"slow_mode"`
	ChatableType string `json:"chatable_type"`
	ChatMode     string `json:"chat_mode"`
	ChatModeOld  string `json:"chat_mode_old"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type AscendingLink struct {
	ID          string `json:"id"`
	ChannelId   string `json:"channel_id"`
	Description string `json:"description"`
	Link        string `json:"link"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Order       int64  `json:"order"`
	Title       string `json:"title"`
}

type Plan struct {
	ID           string `json:"id"`
	ChannelId    string `json:"channel_id"`
	StripePlanId string `json:"stripe_plan_id"`
	Amount       string `json:"amount"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type PreviousLivestream struct {
	ID            string     `json:"id"`
	Slug          string     `json:"slug"`
	ChannelId     string     `json:"channel_id"`
	CreatedAt     string     `json:"created_at"`
	SessionTitle  string     `json:"session_title"`
	IsLive        bool       `json:"is_live"`
	RiskLevelId   string     `json:"risk_level_id"`
	Source        string     `json:"source"`
	TwitchChannel string     `json:"twitch_channel"`
	Language      string     `json:"language"`
	Duration      int64      `json:"duration"`
	IsMature      bool       `json:"is_mature"`
	ViewerCount   int64      `json:"viewer_count"` // Deprecated I guess
	Thumbnail     Thumbnail  `json:"thumbnail"`
	Views         int64      `json:"views"`
	Tags          []string   `json:"tags"`
	Categories    []Category `json:"categories"`
	Video         Video      `json:"video"`
}

type Video struct {
	ID                 string `json:"id"`
	LivestreamId       string `json:"live_stream_id"`
	Slug               string `json:"slug"`
	Thumb              string `json:"thumb"`
	S3                 string `json:"s3"`
	TradingPlateformId string `json:"trading_plateform_id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	UUID               string `json:"uuid"`
	Views              int64  `json:"views"`
	DeletedAt          string `json:"deleted_at"`
}

type Verified struct {
	ID        string `json:"id"`
	ChannelId string `json:"channel_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
