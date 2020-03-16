package ok

type CommonGraphResp struct {
	Success   bool `json:"success"`
	ErrorCode int
	ErrorMsg  string
}

type GetChatsResp struct {
	CommonGraphResp
	Chats  []Chat `json:"chats"`
	Marker string `json:"marker"`
}

type Chat struct {
	Type          string           `json:"type"`
	Status        string           `json:"status"`
	Title         string           `json:"title"`
	Icon          ChatIcon         `json:"icon"`
	Participants  map[string]int64 `json:"participants"`
	LastEventTime int64            `json:"lastEventTime"`
	ChatID        string           `json:"chat_id"`
	OwnerID       string           `json:"owner_id"`
	GroupID       string           `json:"group_id"`
}
type ChatIcon struct {
	URL string `json:"url"`
}
