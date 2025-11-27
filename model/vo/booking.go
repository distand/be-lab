package vo

type Booking struct {
	ID         int32             `json:"id"`
	Uid        int32             `json:"uid"`
	Nickname   string            `json:"nickname"`
	DeviceID   int32             `json:"device_id"`
	DeviceName string            `json:"device_name"`
	Status     int32             `json:"status"`
	Ext        map[string]string `json:"ext"`
	Stime      int64             `json:"stime"`
	Etime      int64             `json:"etime"`
	Ctime      int64             `json:"ctime"`
}
