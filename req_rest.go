package ok

import (
	"net/url"
	"strings"
)

type requestRest interface {
	Method() string
	Values() url.Values
}

type getCurrentUserReq struct {
}

func (r getCurrentUserReq) Method() string {
	return "users.getCurrentUser"
}

func (r getCurrentUserReq) Values() url.Values {
	vals := url.Values{}
	vals.Set("fields", "URL_PROFILE,FIRST_NAME,LAST_NAME,NAME,ODKL_LOGIN,PIC128X128")
	return vals
}

type getUserGroupsReq struct {
}

func (r getUserGroupsReq) Method() string {
	return "group.getUserGroupsV2"
}

func (r getUserGroupsReq) Values() url.Values {
	return url.Values{}
}

type getGroupsReq struct {
	ids    []string
	fields []GroupBeanField
}

func (r getGroupsReq) Method() string {
	return "group.getInfo"
}

func (r getGroupsReq) Values() url.Values {
	v := url.Values{}
	v.Set("uids", strings.Join(r.ids, ","))
	var fieldsBuilder strings.Builder
	for _, f := range r.fields {
		fieldsBuilder.WriteString(string(f))
		fieldsBuilder.WriteString(",")
	}
	fields := fieldsBuilder.String()
	fields = strings.TrimSuffix(fields, ",")
	v.Set("fields", fields)
	return v

}
