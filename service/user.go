package service

import (
	"be-lab/common"
	"be-lab/common/code"
	"be-lab/common/utils"
	"be-lab/model/do"
	"be-lab/model/req"
	"be-lab/model/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *Service) UserList(c *gin.Context, p *req.ListReq) *vo.Page {
	rsp := &vo.Page{}
	list, cnt, err := s.Dal.UserList(c, p.Where(), &p.Page)
	if err != nil {
		return rsp
	}
	dl := make([]*vo.User, 0)
	for _, v := range list {
		dl = append(dl, v.ToVO())
	}
	rsp.List = dl
	rsp.Total = cnt
	return rsp
}

func (s *Service) UserInfo(c *gin.Context) (*vo.User, error) {
	value, ok := s.OpenidMap.Load(utils.Openid(c))
	if !ok {
		return nil, errors.New(code.UserErr)
	}
	return value.(*vo.User), nil
}

func (s *Service) nickname(uid int32) string {
	value, ok := s.UidMap.Load(uid)
	if !ok {
		return "未知用户"
	}
	return value.(*vo.User).Nickname
}

func (s *Service) UserSave(c *gin.Context, req *req.UserSave) error {
	uid := utils.Uid(c)
	if utils.IsAdmin(c) && req.ID > 0 {
		uid = req.ID
	}
	u, err := s.Dal.User(c, uid)
	if err != nil {
		return errors.New(code.SysErr)
	}
	u.Nickname = req.Nickname
	if utils.IsAdmin(c) {
		if req.Status > 0 {
			u.Status = req.Status
		}
		if req.Role > 0 {
			u.Role = req.Role
		}
	}
	s.storeUser(u)
	return s.Dal.UserSave(c, u)
}

func (s *Service) Login(c *gin.Context, openid string) (*vo.User, error) {
	u, err := s.Dal.UserByOpenid(c, openid)
	if err != nil {
		return nil, errors.New(code.SysErr)
	}
	if u == nil {
		//创建用户
		u = &do.User{
			Openid:   openid,
			Nickname: "新用户" + utils.RandStr(6),
			Status:   common.StatusActive,
			Ctime:    time.Now(),
		}
		err = s.Dal.UserAdd(c, u)
		if err != nil {
			return nil, errors.New(code.SysErr)
		}
		u, err = s.Dal.UserByOpenid(c, openid)
		if err != nil {
			return nil, errors.New(code.SysErr)
		}
	}
	if u.Status != common.StatusActive || u.IsDel == common.Deleted {
		return nil, errors.New(code.FreezeErr)
	}
	u.Ltime = time.Now()
	err = s.Dal.UserSave(c, u)
	s.storeUser(u)
	return u.ToVO(), err
}

func (s *Service) ReloadUser() {
	go func() {
		us, err := s.Dal.AllUser()
		if err != nil {
			return
		}
		for _, u := range us {
			s.storeUser(&u)
		}
	}()
}

func (s *Service) storeUser(u *do.User) {
	s.OpenidMap.Store(u.Openid, u.ToVO())
	s.UidMap.Store(u.ID, u.ToVO())
}
