package dal

import (
	"be-lab/model/do"
	"be-lab/model/req"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

func (d *Dal) Booking(c context.Context, id int32) (res *do.Booking, err error) {
	res = &do.Booking{}
	err = d.DB.WithContext(c).
		Where("id = ?", id).
		Where("is_del = 0").
		Take(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (d *Dal) BookingCntByUid(c context.Context, uid int32) int64 {
	var cnt int64
	_ = d.DB.WithContext(c).Model(&do.Booking{}).
		Where("is_del = 0").
		Where("status IN (1,2)").
		Where("uid = ?", uid).
		Count(&cnt)
	return cnt
}

func (d *Dal) BookingList(c context.Context, where map[string]any, p *req.Page) (res []do.Booking, cnt int64, err error) {
	err = d.DB.Model(&do.Booking{}).
		WithContext(c).
		Where(where).
		Where("is_del = 0").
		Offset(p.Offset()).
		Limit(p.Limit()).
		Order("id DESC").
		Find(&res).Error
	d.DB.WithContext(c).Model(&do.Booking{}).
		Where(where).
		Where("is_del = 0").
		Count(&cnt)
	return
}

func (d *Dal) BookingByDevices(c context.Context, ids []int32) (res []do.Booking, err error) {
	err = d.DB.Model(&do.Booking{}).
		WithContext(c).
		Where("device_id IN ?", ids).
		Where("is_del = 0").
		Where("status IN (1,2)").
		Order("stime ASC").
		Find(&res).Error
	return
}

func (d *Dal) BookingAdd(c context.Context, u *do.Booking) error {
	return d.DB.WithContext(c).Create(u).Error
}

func (d *Dal) BookingSave(c context.Context, u *do.Booking) error {
	return d.DB.Model(u).WithContext(c).Updates(u).Error
}

func (d *Dal) BookingPreToActive() (res []do.Booking, err error) {
	err = d.DB.Model(&do.Booking{}).
		Where("is_del = 0").
		Where("status = 1").
		Where("stime <= ?", time.Now().Format(time.DateTime)).
		Find(&res).Error
	return
}

func (d *Dal) BookingActiveToEnd() (res []do.Booking, err error) {
	err = d.DB.Model(&do.Booking{}).
		Where("is_del = 0").
		Where("status = 2").
		Where("etime <= ?", time.Now().Format(time.DateTime)).
		Find(&res).Error
	return
}
