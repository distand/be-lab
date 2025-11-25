package vo

type Booking struct {
	ID         int32  `json:"id"`
	Uid        int32  `json:"uid"`
	Nickname   string `json:"nickname"`
	DeviceID   int32  `json:"device_id"`
	DeviceName string `json:"device_name"`
	Memo       string `json:"memo"`
	Stime      int64  `json:"stime"`
	Etime      int64  `json:"etime"`
	Ctime      int64  `json:"ctime"`
}
