package req

import (
	"be-lab/common"
	"be-lab/model/do"
	"be-lab/model/vo"
	"encoding/json"
	"time"
)

type DeviceSave struct {
	ID     int32          `json:"id"`
	Type   int32          `json:"type"`
	Name   string         `json:"name"`
	Memo   string         `json:"memo"`
	Status int32          `json:"status"`
	Rule   *vo.DeviceRule `json:"rule,omitempty"`
	Config *vo.DeviceCfg  `json:"config,omitempty"`
}

func (d *DeviceSave) BuildDo(res *do.Device) *do.Device {
	if res == nil {
		res = &do.Device{
			Ctime:  time.Now(),
			Status: common.StatusActive,
		}
	}
	if d.ID > 0 {
		res.ID = d.ID
	}
	if d.Type > 0 {
		res.Type = d.Type
	}
	if d.Name != "" {
		res.Name = d.Name
	}
	if d.Memo != "" {
		res.Memo = d.Memo
	}
	if d.Status >= 0 {
		res.Status = d.Status
	}
	if d.Rule != nil {
		rule, _ := json.Marshal(d.Rule)
		res.Rule = string(rule)
	}
	if d.Config != nil {
		cfg, _ := json.Marshal(d.Config)
		res.Rule = string(cfg)
	}
	return res
}
