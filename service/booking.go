package service

import (
	"be-lab/common"
	"be-lab/common/code"
	"be-lab/common/utils"
	"be-lab/model/do"
	"be-lab/model/req"
	"be-lab/model/vo"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Service) BookingList(c *gin.Context, p *req.ListReq) *vo.Page {
	var (
		rsp = &vo.Page{
			List: make([]*vo.Booking, 0),
		}
		uid      = utils.Uid(c)
		nickname = s.nickname(uid)
	)
	where := p.Where()
	where["uid"] = uid
	list, cnt, err := s.Dal.BookingList(c, where, &p.Page)
	if err != nil || cnt == 0 {
		return rsp
	}
	var (
		ids []int32
		res []*vo.Booking
	)
	for _, v := range list {
		ids = append(ids, v.DeviceID)
	}
	names := s.deviceNames(c, ids)
	for _, v := range list {
		res = append(res, v.ToVO(nickname, names[v.DeviceID]))
	}
	rsp.List = res
	rsp.Total = cnt
	return rsp
}

func (s *Service) BookingSave(c *gin.Context, p *req.BookingSave) error {
	if p.Stime >= p.Etime || p.DeviceID == 0 || p.Stime < time.Now().Unix() {
		return errors.New(code.ParamErr)
	}
	device, err := s.Dal.Device(c, p.DeviceID)
	if err != nil || device == nil || device.Status == common.StatusInactive {
		return errors.New("仪器不可用")
	}
	rule := device.ToVO().Rule
	if rule.MaxOnce > 0 && p.Etime-p.Stime > rule.MaxOnce*common.Hour {
		return errors.New("单次预约时间超过限制")
	}
	if device.Status == common.StatusUsing || s.isConflict(c, p) {
		return errors.New("预约时间与他人冲突")
	}
	if p.ID > 0 {
		d, err := s.Dal.Booking(c, p.ID)
		if err != nil || d == nil {
			return err
		}
		//非本人或者已取消的预约不能修改
		if d.Uid != utils.Uid(c) || d.Status == common.StatusInactive {
			return errors.New(code.Forbidden)
		}
		return s.Dal.BookingSave(c, p.BuildDo(d).SetUser(utils.Uid(c)))
	}
	return s.Dal.BookingAdd(c, p.BuildDo(nil).SetUser(utils.Uid(c)))
}

func (s *Service) isConflict(c *gin.Context, p *req.BookingSave) bool {
	res, err := s.Dal.BookingByDevices(c, []int32{p.DeviceID})
	if err != nil {
		return true
	}
	if len(res) == 0 {
		return false
	}
	for _, b := range res {
		if b.ID == p.ID {
			continue
		}
		if p.Stime < b.Etime.Unix() && p.Etime > b.Stime.Unix() {
			return true
		}
	}
	return false
}

func (s *Service) BookingJob() {
	c := context.Background()
	active, _ := s.Dal.BookingPreToActive()
	for _, v := range active {
		_ = s.Dal.BookingSave(c, &do.Booking{ID: v.ID, Status: common.StatusUsing})
		_ = s.Dal.DeviceSave(c, &do.Device{ID: v.DeviceID, Status: common.StatusUsing})
	}
	end, _ := s.Dal.BookingActiveToEnd()
	for _, v := range end {
		_ = s.Dal.BookingSave(c, &do.Booking{ID: v.ID, Status: common.StatusInactive})
		_ = s.Dal.DeviceSave(c, &do.Device{ID: v.DeviceID, Status: common.StatusActive})
	}
}
