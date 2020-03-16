package ok

type requestGraph interface {
	MethodHTTP() string
	Method() string
	Node() string
	FieldsUrlEncoded() string
}

type commonReq struct {
}

func (r commonReq) MethodHTTP() string {
	return "POST"
}
func (r commonReq) Node() string {
	return "me"
}
func (r commonReq) FieldsUrlEncoded() string {
	return ""
}

//setWebHookReq - set web hook request from
type setWebHookReq struct {
	commonReq
	Url     string   `json:"url,omitempty"`
	Types   []string `json:"types,omitempty"`
	Polling bool     `json:"longPolling,omitempty"`
}

func (r setWebHookReq) Method() string {
	return "subscribe"
}

//setWebHookReq - set web hook request from
type dropWebHookReq struct {
	commonReq
	Url string `json:"url"`
}

func (r dropWebHookReq) Method() string {
	return "unsubscribe"
}

//getWebHooksReq - get all subscribed web hooks
type getWebHooksReq struct {
	commonReq
}

func (r getWebHooksReq) Method() string {
	return "subscriptions"
}

//getWebHooksReq - get all subscribed web hooks
type getUpdatesReq struct {
	commonReq
}

func (r getUpdatesReq) Method() string {
	return "updates"
}

func (r getUpdatesReq) MethodHTTP() string {
	return "GET"
}

//getWebHooksReq - get all subscribed web hooks
type getChatsReq struct {
	commonReq
	Marker string `json:"marker,omitempty"` //Маркер для постраничной выгрузки данных. При первом запросе указывать не требуется, для последующих запросов маркер берется в ответе метода
	Count  int    `json:"count,omitempty"`  //Количество запрашиваемых чатов. От 1 до 100
}

func (r getChatsReq) Method() string {
	return "chats"
}

func (r getChatsReq) MethodHTTP() string {
	return "GET"
}
