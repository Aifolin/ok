package ok

import (
	"encoding/json"
	"time"
)

type CommonResp struct {
	ErrorCode json.Number `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
}

type GetGroupsResp []GroupFullInfo

type GroupFullInfo struct {
	Abbreviation              string     `json:"abbreviation"`
	AccessType                string     `json:"access_type"` //CLOSED|OPEN|SECRET
	Address                   string     `json:"address"`
	AdminId                   string     `json:"admin_id"`
	Attrs                     GroupAttrs `json:"attrs"`
	Business                  bool       `json:"business"`
	Category                  string     `json:"category"` //ARMY|COLLEAGE|CUSTOM|FACULTY|HAPPENING|HOLIDAY|MOIMIR|OTHER|SCHOOL|UNIVERSITY|WORKPLACE
	City                      string     `json:"city"`
	CommentAsOfficial         bool       `json:"comment_as_official"`
	Community                 bool       `json:"community"`
	ContentAsOfficial         bool       `json:"content_as_official"`
	Country                   string     `json:"country"`
	CreatedMs                 int64      `json:"created_ms"`
	Description               string     `json:"description"`
	EndDate                   int64      `json:"end_date"`
	FeedSubscription          bool       `json:"feed_subscription"`
	FollowersCount            int        `json:"followers_count"`
	FriendsCount              int        `json:"friends_count"`
	GraduateYear              int        `json:"graduate_year"`
	HomepageName              string     `json:"homepage_name"`
	HomepageUrl               string     `json:"homepage_url"`
	LocationId                int64      `json:"location_id"`
	LocationLatitude          float64    `json:"location_latitude"`
	LocationLongitude         float64    `json:"location_longitude"`
	LocationZoom              int        `json:"location_zoom"`
	MainPageTab               string     `json:"main_page_tab"`
	MainPhoto                 GroupPhoto `json:"main_photo"`
	MembersCount              int        `json:"members_count"`
	MentionsSubscription      bool       `json:"mentions_subscription"`
	MessagesAllowed           bool       `json:"messages_allowed"`
	MinAge                    int        `json:"min_age"`
	Name                      string     `json:"name"`
	NotificationsSubscription bool       `json:"notifications_subscription"`
	PartnerProgramStatus      string     `json:"partner_program_status"`
	Phone                     string     `json:"phone"`
	PhotoId                   string     `json:"photo_id"`
	PhotosTabHidden           bool       `json:"photos_tab_hidden"`
	PicAvatar                 string     `json:"picAvatar"`
	PossibleMembersCount      int        `json:"possible_members_count"`
	Premium                   bool       `json:"premium"`
	Private                   bool       `json:"private"`
	ProductsTabHidden         bool       `json:"products_tab_hidden"`
	Ref                       string     `json:"ref"`
	Role                      string     `json:"role"`
	ScopeId                   string     `json:"scope_id"`
	ShortName                 string     `json:"short_name"`
	StartDate                 int64      `json:"start_date"`
	Status                    string     `json:"status"` //ACTIVE|DELETED|DISABLED|MODERATION
	SubcategoryId             string     `json:"subcategory_id"`
	Tags                      []string   `json:"tags"`
	TransfersAllowed          bool       `json:"transfers_allowed"`
	Uid                       string     `json:"uid"`
	VideoTabHidden            bool       `json:"video_tab_hidden"`
	YearFrom                  int        `json:"year_from"`
	YearTo                    int        `json:"year_to"`
}

type GroupPhoto struct {
	PhotoBig   string `json:"pic1024x768"`
	PhotoSmall string `json:"pic128x128"`
}

type GroupAttrs struct {
	Flags string `json:"flags"`
}

type GetUserGroupsResp struct {
	CommonResp
	Anchor string           `json:"anchor"` //pagination identifier
	Groups []GroupSmallInfo `json:"groups"`
	Status string           `json:"status"`
	UserId string           `json:"userId"`
}

type GroupSmallInfo struct {
	BlockReason string `json:"block_reason"`
	GroupId     string `json:"groupId"`
	Role        string `json:"role"` //[ANALYST|EDITOR|MODERATOR|SUPER_MODERATOR]+
}

type GetCurrentUserResp struct {
	CommonResp
	Accessible                    bool           `json:"accessible"`
	Age                           int            `json:"age"`
	AllowsMessagingOnlyForFriends bool           `json:"allows_messaging_only_for_friends"`
	Birthday                      string         `json:"birthday"`
	BirthdaySet                   bool           `json:"birthdaySet"`
	Blocked                       bool           `json:"blocked"`
	Blocks                        bool           `json:"blocks"`
	Business                      bool           `json:"business"`
	CanVCall                      bool           `json:"can_vcall"`
	CanVMail                      bool           `json:"can_vmail"`
	Capabilities                  string         `json:"capabilities"`
	CurrentStatus                 string         `json:"current_status"`
	CurrentStatusDate             time.Time      `json:"current_status_date"`
	Email                         string         `json:"email"`
	Executor                      bool           `json:"executor"`
	FirstName                     string         `json:"first_name"`
	Gender                        string         `json:"gender"`
	LastName                      string         `json:"last_name"`
	LastOnline                    string         `json:"last_online"`
	Locale                        string         `json:"locale"`
	Location                      UserLocation   `json:"location"`
	Login                         string         `json:"login"`
	Mobile                        string         `json:"mobile"`
	Name                          string         `json:"name"`
	Online                        string         `json:"online"`
	PhotoID                       string         `json:"photo_id"`
	PhotoBig                      string         `json:"pic1024x768"`
	PhotoSmall                    string         `json:"pic128x128"`
	PicBase                       string         `json:"pic_base"`
	PicFull                       string         `json:"pic_full"`
	PicMax                        string         `json:"pic_max"`
	PossibleRelations             []UserRelation `json:"possible_relations"`
	Premium                       bool           `json:"premium"`
	//todo check presents
	//presents [{}]
	Private bool `json:"private"`
	//todo check Relations
	//Relations [{}]
	//relationship {}
	Uid              string `json:"uid"`
	UrlChat          string `json:"url_chat"`
	UrlChatMobile    string `json:"url_chat_mobile"`
	UrlProfile       string `json:"url_profile"`
	UrlProfileMobile string `json:"url_profile_mobile"`
	Vip              bool   `json:"vip"`
}

type UserLocation struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
}

type UserRelation string //ALL|BROTHERSISTER|CHILD|CHILDINLAW|CLASSMATE|CLOSEFRIEND|COLLEGUE|COMPANIONINARMS|CURSEMATE|GODCHILD|GODPARENT|GRANDCHILD|GRANDPARENT|LOVE|NEPHEW|PARENT|PARENTINLAW|PLAYINGTOGETHER|RELATIVE|SPOUSE|UNCLEAUNT
