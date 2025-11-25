package service

import (
	"be-lab/common"
	"be-lab/model/req"
	"be-lab/model/vo"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Service) DeviceList(c *gin.Context, p *req.ListReq) *vo.Page {
	rsp := &vo.Page{}
	list, cnt, err := s.Dal.DeviceList(c, p.Where(), &p.Page)
	if err != nil {
		return rsp
	}
	dl := make([]*vo.Device, 0)
	for _, v := range list {
		dl = append(dl, v.ToVO())
	}
	s.handleBooking(c, dl)
	rsp.List = dl
	rsp.Total = cnt
	return rsp
}

func (s *Service) DeviceSave(c *gin.Context, p *req.DeviceSave) error {
	if p.ID > 0 {
		d, err := s.Dal.Device(c, p.ID)
		if err != nil || d == nil {
			return err
		}
		return s.Dal.DeviceSave(c, p.BuildDo(d))
	}
	return s.Dal.DeviceAdd(c, p.BuildDo(nil))
}

func (s *Service) deviceNames(c *gin.Context, ids []int32) map[int32]string {
	res := make(map[int32]string)
	if len(ids) == 0 {
		return res
	}
	for _, id := range ids {
		res[id] = "未知设备"
	}
	list, err := s.Dal.Devices(c, ids)
	if err != nil || len(list) == 0 {
		return res
	}
	for _, v := range list {
		res[v.ID] = v.Name
	}
	return res
}

func (s *Service) handleBooking(c *gin.Context, dl []*vo.Device) {
	if len(dl) == 0 {
		return
	}
	var ids []int32
	for _, v := range dl {
		ids = append(ids, v.ID)
	}
	res, err := s.Dal.BookingByDevices(c, ids)
	if err != nil || len(res) == 0 {
		return
	}
	now := time.Now()
	for _, d := range dl {
		for _, b := range res {
			if b.DeviceID != d.ID {
				continue
			}
			if b.Stime.Before(now) && b.Etime.After(now) {
				//有进行中的预约
				d.Status = common.StatusUsing
			}
			d.List = append(d.List, b.ToVO(s.nickname(b.Uid), d.Name))
		}
	}
}
