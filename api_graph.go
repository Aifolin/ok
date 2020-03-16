package ok

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	maxGetChatsCount = 100
)

type GroupGraphApi struct {
	token  string
	client *http.Client
}

func NewGroupGraphApi(token string) *GroupGraphApi {
	api := new(GroupGraphApi)
	api.token = token
	api.client = &http.Client{Timeout: time.Second * 15}
	return api
}

func (g *GroupGraphApi) SetWebHook(url string, polling bool) error {
	var types []string
	if polling {
		url = ""
		types = []string{
			string(MESSAGE_CREATED),
			string(MESSAGE_CALLBACK),
			string(CHAT_SYSTEM)}
	}
	var resp CommonGraphResp
	err := g.call(
		setWebHookReq{
			Url:     url,
			Polling: polling,
			Types:   types,
		},
		&resp)
	if err != nil {
		return err
	}
	if resp.ErrorCode != 0 {
		return fmt.Errorf("error: %+v", resp)
	}
	return nil
}

func (g *GroupGraphApi) DropWebHook(url string) error {
	var resp CommonGraphResp
	err := g.call(dropWebHookReq{Url: url}, &resp)
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("response success=false: %+v", resp)
	}
	return nil
}

func (g *GroupGraphApi) GetWebHooks() error {
	var resp CommonGraphResp
	err := g.call(getWebHooksReq{}, &resp)
	if err != nil {
		return err
	}
	if resp.ErrorCode != 0 {
		return fmt.Errorf("error: %+v", resp)
	}
	return nil
}

func (g *GroupGraphApi) GetUpdates() error {
	var resp CommonGraphResp
	err := g.call(getUpdatesReq{}, &resp)
	if err != nil {
		return err
	}
	if resp.ErrorCode != 0 {
		return fmt.Errorf("error: %+v", resp)
	}
	return nil
}

func (g *GroupGraphApi) GetChats() (resp GetChatsResp, err error) {
	more := true
	marker := ""
	for more {
		err = g.call(getChatsReq{
			Marker: marker,
			Count:  maxGetChatsCount,
		}, &resp)
		if err != nil {
			return
		}
		if resp.ErrorCode != 0 {
			return resp, fmt.Errorf("error: %+v", resp.CommonGraphResp)
		}
		if resp.Marker != "" {
			marker = resp.Marker
		} else {
			more = false
		}
	}

	return
}

func (g *GroupGraphApi) GetMembers() ([]GetCurrentUserResp, error) {
	return nil, nil
}

func (g *GroupGraphApi) SendMsg(chatId string, msg string) error {
	return fmt.Errorf("not implemented")
}

func (g *GroupGraphApi) GetMsgs(chatId string) error {
	return nil
}

//call - signs req and calls http API endpoint
func (c *GroupGraphApi) call(req requestGraph, resp interface{}) (err error) {
	reqData, err := json.MarshalIndent(req, " ", "   ")
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(reqData)
	r, err := http.NewRequest(
		req.MethodHTTP(),
		fmt.Sprintf("%s/%s/%s?access_token=%s", baseUrlGRAPH, req.Node(), req.Method(), c.token),
		body)
	r.Header.Set("Content-Type", "application/json;charset=utf-8")
	logs.Debug("%+v", r)
	logs.Debug("%s", reqData)
	httpResp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	respData, err := ioutil.ReadAll(httpResp.Body)
	logs.Debug("%s", respData)
	if err != nil {
		return err
	}
	return json.Unmarshal(respData, resp)
}
