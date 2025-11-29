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
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Service) BookingList(c *gin.Context, p *req.ListReq) *vo.Page {
	var (
		rsp = &vo.Page{
			List: make([]*vo.Booking, 0),
		}
	)
	list, cnt, err := s.Dal.BookingList(c, p.Where(), &p.Page)
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
		res = append(res, v.ToVO(s.nickname(v.Uid), names[v.DeviceID]))
	}
	rsp.List = res
	rsp.Total = cnt
	return rsp
}

func (s *Service) BookingSave(c *gin.Context, p *req.BookingSave) error {
	if err := s.checkBooking(c, p); err != nil {
		return err
	}
	if p.ID > 0 {
		d, err := s.Dal.Booking(c, p.ID)
		if err != nil || d == nil {
			return err
		}
		//非本人或者已取消的预约不能修改
		if (!utils.IsAdmin(c) && d.Uid != utils.Uid(c)) || d.Status == common.StatusInactive {
			return errors.New(code.Forbidden)
		}
		return s.Dal.BookingSave(c, p.BuildDo(utils.Uid(c), d))
	}
	return s.Dal.BookingAdd(c, p.BuildDo(utils.Uid(c), nil))
}

func (s *Service) checkBooking(c *gin.Context, p *req.BookingSave) error {
	if p.ID > 0 && p.IsDel > 1 {
		return nil
	}
	if p.Stime >= p.Etime || p.DeviceID == 0 || p.Stime < time.Now().Unix() {
		return errors.New(code.ParamErr)
	}
	device, err := s.Dal.Device(c, p.DeviceID)
	if err != nil || device == nil || device.Status == common.StatusInactive {
		return errors.New("仪器不可用")
	}
	rule := device.ToVO().Rule
	if rule.MaxOnce > 0 && p.Etime-p.Stime > rule.MaxOnce {
		return errors.New("超过单次最大时间限制")
	}
	for _, field := range rule.RequireFields {
		if _, ok := common.DeviceFields[field]; !ok {
			return errors.New("不支持的参数：" + field)
		}
		if v, ok := p.Ext[field]; !ok || v == "" {
			return errors.New("缺少必填参数：" + field)
		}
	}
	res, err := s.Dal.BookingByDevices(c, []int32{p.DeviceID})
	if err != nil {
		return errors.New(code.SysErr)
	}
	if len(res) == 0 {
		return nil
	}
	var selfCnt int64
	uid := utils.Uid(c)
	for _, b := range res {
		if b.Uid == uid {
			selfCnt++
		}
		if b.ID == p.ID {
			continue
		}
		if p.Stime < b.Etime.Unix() && p.Etime > b.Stime.Unix() {
			return errors.New("预约时间段仪器被占用")
		}
	}
	if selfCnt >= rule.MaxContinuous {
		return errors.New(fmt.Sprintf("超过连续最大预约次数: %d", rule.MaxContinuous))
	}
	return nil
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
