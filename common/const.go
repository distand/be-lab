package common

const (
	Deleted   = 1
	AuthAdmin = 1000
)

const (
	StatusUnknown = iota
	StatusActive
	StatusUsing
	StatusInactive
)

const (
	KeyUid    = "uid"
	KeyOpenid = "openid"
	KeyRole   = "role"
)

var DeviceFields = map[string]string{
	"memo":         "用途",
	"mobile_phase": "流动相",
	"condition":    "反应条件",
	"column":       "柱子类型",
	"bacterial":    "菌种",
	"volume":       "体积",
	"temperature":  "温度",
}
