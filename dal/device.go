package dal

import (
	"be-lab/model/do"
	"be-lab/model/req"
	"be-lab/model/vo"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (d *Dal) Device(c context.Context, id int32) (res *do.Device, err error) {
	res = &do.Device{}
	err = d.DB.WithContext(c).
		Where("id = ?", id).
		Where("is_del = 0").
		Take(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (d *Dal) Devices(c context.Context, ids []int32) (res []do.Device, err error) {
	err = d.DB.WithContext(c).Model(&do.Device{}).
		Where("id IN ?", ids).
		Where("is_del = 0").
		Find(&res).Error
	return
}

func (d *Dal) DeviceList(c context.Context, where map[string]any, p *req.Page) (res []do.Device, cnt int64, err error) {
	err = d.DB.WithContext(c).Model(&do.Device{}).
		Where(where).
		Where("is_del = 0").
		Offset(p.Offset()).
		Limit(p.Limit()).
		Find(&res).Error
	d.DB.WithContext(c).Model(&do.Device{}).
		Where(where).
		Where("is_del = 0").
		Count(&cnt)
	return
}

func (d *Dal) DeviceCount(c context.Context) (res []vo.StatusCnt) {
	_ = d.DB.WithContext(c).Model(&do.Device{}).
		Select("status, COUNT(*) as cnt").
		Where("is_del = 0").
		Group("status").
		Find(&res).Error
	return
}

func (d *Dal) DeviceAdd(c context.Context, u *do.Device) error {
	return d.DB.WithContext(c).Create(u).Error
}

func (d *Dal) DeviceSave(c context.Context, u *do.Device) error {
	return d.DB.Model(u).WithContext(c).Updates(u).Error
}
