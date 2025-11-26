package vo

type Device struct {
	ID     int32      `json:"id"`
	Type   int32      `json:"type"`
	Name   string     `json:"name"`
	Memo   string     `json:"memo"`
	Status int32      `json:"status"`
	Rule   DeviceRule `json:"rule"`
	Config DeviceCfg  `json:"config"`
	List   []*Booking `json:"list"`
}

type DeviceRule struct {
	MaxOnce       int64    `json:"max_once,omitempty"`
	MaxContinuous int64    `json:"max_continuous,omitempty"`
	RequireFields []string `json:"require_fields,omitempty"`
}

type DeviceCfg struct {
}
