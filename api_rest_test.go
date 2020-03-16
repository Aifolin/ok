package ok

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

const (
	accessKey = "-s-csWuR-5vQpXub-4eOE7QbD8zNLZtK942UHY05D5VUIT3"
	appPublic = "CLOOBKJGDIHBABABA"
	appSecret = "3321675860F37563EC8EDDD0"
)

func init() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}

func TestOKApi_GetUserInfo(t *testing.T) {
	api := NewOKApi(accessKey, appPublic, appSecret)
	info, err := api.GetUserInfo()
	if err != nil {
		t.Fatal(err)
	}

	logs.Debug("%+v", info)
	logs.Debug("%+v", info.UrlProfile)
	logs.Debug("%+v", info.Uid)
	logs.Debug("%+v", info.Name)
	logs.Debug("%+v", info.PhotoSmall)
}

func TestOKApi_GetUserGroups(t *testing.T) {
	api := NewOKApi(accessKey, appPublic, appSecret)
	groups, err := api.GetUserGroups()
	if err != nil {
		t.Fatal(err)
	}

	logs.Debug("%+v", groups)

	var groupIds []string
	for _, g := range groups.Groups {
		groupIds = append(groupIds, g.GroupId)
	}

	fullGroups, err := api.GetGroups(
		groupIds,
		[]GroupBeanField{
			ABBREVIATION,
			ADDRESS,
			//ADMIN_ID,
			CATEGORY,
			CREATED_MS,
			//DESCRIPTION,
			FOLLOWERS_COUNT,
			HOMEPAGE_NAME,
			HOMEPAGE_URL,
			MAIN_PHOTO,
			MEMBERS_COUNT,
			NAME,
			PIC_AVATAR,
			SHORTNAME,
			UID,
			MANAGE_MEMBERS,
			STATUS,
		})
	if err != nil {
		t.Fatal(err)
	}

	logs.Debug("%+v", fullGroups)

}
