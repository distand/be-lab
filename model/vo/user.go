package vo

type User struct {
	ID       int32  `json:"id"`
	Openid   string `json:"openid"`
	Nickname string `json:"nickname"`
	Status   int32  `json:"status"`
	Role     int32  `json:"role"`
	Ltime    int64  `json:"ltime"`
	Ctime    int64  `json:"ctime"`
}
