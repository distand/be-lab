package req

type Page struct {
	Pn int `form:"pn,omitempty" json:"pn,omitempty"`
	Ps int `form:"ps,omitempty" json:"ps,omitempty"`
}

func (p *Page) Offset() int {
	if p.Pn < 1 {
		p.Pn = 1
	}
	if p.Ps > 100 {
		p.Ps = 100
	}
	return (p.Pn - 1) * p.Ps
}

func (p *Page) Limit() int {
	if p.Ps < 1 {
		p.Ps = 100
	}
	return p.Ps
}

type ListReq struct {
	Page
	Status   int32 `form:"status"`
	Type     int32 `form:"type"`
	Uid      int32 `form:"uid"`
	DeviceId int32 `form:"device_id"`
}

func (p *ListReq) Where() map[string]any {
	mp := make(map[string]any)
	if p.Uid > 0 {
		mp["uid"] = p.Uid
	}
	if p.Status > 0 {
		mp["status"] = p.Status
	}
	if p.Type > 0 {
		mp["type"] = p.Type
	}
	if p.DeviceId > 0 {
		mp["device_id"] = p.DeviceId
	}
	return mp
}
