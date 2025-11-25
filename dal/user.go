package dal

import (
	"be-lab/model/do"
	"be-lab/model/req"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (d *Dal) User(c context.Context, id int32) (res *do.User, err error) {
	res = &do.User{}
	err = d.DB.WithContext(c).
		Where("id = ?", id).
		Take(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (d *Dal) UserByOpenid(c context.Context, openid string) (res *do.User, err error) {
	res = &do.User{}
	err = d.DB.WithContext(c).Model(&do.User{}).
		Where("openid = ?", openid).
		Take(res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (d *Dal) UserList(c context.Context, where map[string]any, p *req.Page) (res []do.User, cnt int64, err error) {
	err = d.DB.WithContext(c).Model(&do.User{}).
		Where(where).
		Where("is_del = 0").
		Offset(p.Offset()).
		Limit(p.Limit()).
		Find(&res).Error
	d.DB.WithContext(c).Model(&do.User{}).
		Where(where).
		Where("is_del = 0").
		Count(&cnt)
	return
}

func (d *Dal) UserAdd(c context.Context, u *do.User) error {
	return d.DB.WithContext(c).Create(u).Error
}

func (d *Dal) UserSave(c context.Context, u *do.User) error {
	return d.DB.Model(u).WithContext(c).Updates(u).Error
}

func (d *Dal) AllUser() (res []do.User, err error) {
	err = d.DB.Model(&do.User{}).Find(&res).Error
	return
}
