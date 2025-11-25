package req

type UserSave struct {
	ID       int32  `json:"id"`
	Nickname string `json:"nickname"`
	Status   int32  `json:"status"`
	Role     int32  `json:"role"`
}
