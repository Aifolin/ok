package ok

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

const (
	testGroup      = "58479189229801"
	testHookUrl    = "https://test2.oopla.tk/ok/webhook"
	testGraphToken = "tkn18PUTCPo1ciryWepr81aK54HObBhLMVCM8A8kgksXwcoENM79id7fMr7xoMdrjcOA2:CJLQNKJGDIHBABABA"
)

func TestGroupGraphApi_SetWebHook(t *testing.T) {
	api := NewGroupGraphApi(testGraphToken)
	err := api.SetWebHook(testHookUrl, true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGroupGraphApi_DropWebHook(t *testing.T) {
	api := NewGroupGraphApi(testGraphToken)
	err := api.DropWebHook(testHookUrl)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGroupGraphApi_GetWebHooks(t *testing.T) {
	api := NewGroupGraphApi(testGraphToken)
	err := api.GetWebHooks()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGroupGraphApi_GetUpdates(t *testing.T) {
	api := NewGroupGraphApi(testGraphToken)
	err := api.GetUpdates()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGroupGraphApi_GetChats(t *testing.T) {
	api := NewGroupGraphApi(testGraphToken)
	chats, err := api.GetChats()
	if err != nil {
		t.Fatal(err)
	}

	logs.Debug(chats)
}
