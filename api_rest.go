package ok

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//OKApi - API odnoklassniki
type OKApi struct {
	accessToken  string //user session token (got from OAuth)
	appPublicKey string
	appSecretKey string
	client       *http.Client
}

//NewOKApi - creates new OKApi with user accessToken and application public and secret keys
func NewOKApi(accessToken, appPublic, appSecret string) *OKApi {
	api := new(OKApi)
	api.accessToken = accessToken
	api.appPublicKey = appPublic
	api.appSecretKey = appSecret
	api.client = &http.Client{Timeout: time.Second * 15}
	return api
}

//GetUserGroups - returns the current users's groups
func (c *OKApi) GetUserGroups() (resp GetUserGroupsResp, err error) {
	err = c.call(getUserGroupsReq{}, &resp)
	return resp, err
}

//GetUserInfo - returns the current users's info
func (c *OKApi) GetUserInfo() (info GetCurrentUserResp, err error) {
	err = c.call(getCurrentUserReq{}, &info)
	if info.CommonResp.ErrorCode.String() != "" {
		return GetCurrentUserResp{}, fmt.Errorf("get user info: error code: %s, msg: %s",
			info.CommonResp.ErrorCode, info.CommonResp.ErrorMsg)
	}
	return info, err
}

//GetGroup - returns group info by it's id
func (c *OKApi) GetGroups(ids []string, fileds []GroupBeanField) (info GetGroupsResp, err error) {
	err = c.call(getGroupsReq{
		ids: ids, fields: fileds,
	},
		&info)
	return info, err
}

//call - signs req and calls http API endpoint
func (c *OKApi) call(req requestRest, resp interface{}) (err error) {
	values := req.Values()
	values.Set("application_key", c.appPublicKey)
	values.Set("format", "JSON")
	values.Set("method", req.Method())

	sign, err := c.sign(values)
	if err != nil {
		return fmt.Errorf("sign req: %s", err)
	}

	values.Set("sig", sign)
	values.Set("access_token", c.accessToken)
	r, err := http.NewRequest(
		"GET", fmt.Sprintf("%s?method=%s&%s", baseUrlREST, req.Method(), values.Encode()), nil)

	logs.Debug(r)
	httpResp, err := c.client.Do(r)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	logs.Debug("%s", data)
	var commonResp CommonResp
	//specially ignore the unmarshal errors,
	//because of some resps are arrays, but error resps are json
	json.Unmarshal(data, &commonResp)
	if commonResp.ErrorCode.String() != "" || commonResp.ErrorMsg != "" {
		return fmt.Errorf("call %s code:%v error: %s",
			req.Method(), commonResp.ErrorCode, commonResp.ErrorMsg)
	}

	return json.Unmarshal(data, resp)
}

//sign - calculates signature for req
func (c *OKApi) sign(values url.Values) (sign string, err error) {
	secret := fmt.Sprintf("%x", md5.Sum([]byte(c.accessToken+c.appSecretKey)))
	valuesUnescaped, err := url.QueryUnescape(values.Encode())
	if err != nil {
		return "", fmt.Errorf("unescape url vals:%s", err)
	}
	sortedParams := strings.ReplaceAll(valuesUnescaped, "&", "")
	sign = fmt.Sprintf("%x", md5.Sum([]byte(sortedParams+secret)))
	return sign, nil
}
