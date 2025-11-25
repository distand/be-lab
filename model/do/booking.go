package do

import (
	"be-lab/model/vo"
	"time"
)

type Booking struct {
	ID       int32     `json:"id" gorm:"primaryKey"`
	Uid      int32     `json:"uid" gorm:"column:uid"`
	DeviceID int32     `json:"device_id" gorm:"column:device_id"`
	Memo     string    `json:"memo" gorm:"column:memo"`
	Status   int32     `json:"status" gorm:"column:status"`
	IsDel    int32     `json:"is_del" gorm:"column:is_del"`
	Stime    time.Time `json:"stime" gorm:"column:stime"`
	Etime    time.Time `json:"etime" gorm:"column:etime"`
	Ctime    time.Time `json:"ctime" gorm:"column:ctime"`
}

func (m *Booking) TableName() string {
	return "lab_booking"
}

func (m *Booking) ToVO(nickname, deviceName string) *vo.Booking {
	res := &vo.Booking{
		ID:         m.ID,
		Uid:        m.Uid,
		Nickname:   nickname,
		DeviceID:   m.DeviceID,
		DeviceName: deviceName,
		Memo:       m.Memo,
		Stime:      m.Stime.Unix(),
		Etime:      m.Etime.Unix(),
		Ctime:      m.Ctime.Unix(),
	}
	return res
}

func (m *Booking) SetUser(uid int32) *Booking {
	m.Uid = uid
	return m
}
