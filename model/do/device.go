package do

import (
	"be-lab/model/vo"
	"encoding/json"
	"time"
)

type Device struct {
	ID     int32     `json:"id" gorm:"primaryKey"`
	Type   int32     `json:"type" gorm:"column:type"`
	Name   string    `json:"name" gorm:"column:name"`
	Memo   string    `json:"memo" gorm:"column:memo"`
	Rule   string    `json:"rule" gorm:"column:rule"`
	Status int32     `json:"status" gorm:"column:status"`
	IsDel  int32     `json:"is_del" gorm:"column:is_del"`
	Config string    `json:"config" gorm:"column:config"`
	Ctime  time.Time `json:"ctime" gorm:"column:ctime"`
}

func (m *Device) TableName() string {
	return "lab_device"
}

func (m *Device) ToVO() *vo.Device {
	var (
		rule vo.DeviceRule
		cfg  vo.DeviceCfg
	)
	_ = json.Unmarshal([]byte(m.Rule), &rule)
	_ = json.Unmarshal([]byte(m.Config), &cfg)
	res := &vo.Device{
		ID:     m.ID,
		Type:   m.Type,
		Name:   m.Name,
		Memo:   m.Memo,
		Status: m.Status,
		Rule:   rule,
		Config: cfg,
		List:   make([]*vo.Booking, 0),
	}
	return res
}
