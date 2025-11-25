package req

import (
	"be-lab/common"
	"be-lab/model/do"
	"time"
)

type BookingSave struct {
	ID       int32  `json:"id"`
	DeviceID int32  `json:"device_id"`
	Memo     string `json:"memo"`
	IsDel    int32  `json:"is_del"`
	Stime    int64  `json:"stime"`
	Etime    int64  `json:"etime"`
}

func (d *BookingSave) BuildDo(res *do.Booking) *do.Booking {
	if res == nil {
		res = &do.Booking{
			Ctime:  time.Now(),
			Status: common.StatusActive,
		}
	}
	if d.ID > 0 {
		res.ID = d.ID
	}
	if d.DeviceID > 0 {
		res.DeviceID = d.DeviceID
	}
	if d.Memo != "" {
		res.Memo = d.Memo
	}
	if d.IsDel > 0 {
		res.IsDel = d.IsDel
	}
	if d.Stime > 0 {
		res.Stime = time.Unix(d.Stime, 0)
	}
	if d.Etime > 0 {
		res.Etime = time.Unix(d.Etime, 0)
	}
	return res
}
