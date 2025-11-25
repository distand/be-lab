package service

import (
	"be-lab/common"
	"be-lab/common/utils"
	"be-lab/dal"
	"be-lab/model/vo"
	"github.com/gin-gonic/gin"
	"sync"
)

type Service struct {
	Dal       *dal.Dal
	OpenidMap sync.Map
	UidMap    sync.Map
}

func NewService(dal *dal.Dal) *Service {
	s := &Service{
		Dal:       dal,
		OpenidMap: sync.Map{},
		UidMap:    sync.Map{},
	}
	s.ReloadUser()
	return s
}

func (s *Service) Index(c *gin.Context) *vo.Stats {
	res := &vo.Stats{
		BookingMy: s.Dal.BookingCntByUid(c, utils.Uid(c)),
	}
	cnt := s.Dal.DeviceCount(c)
	for _, v := range cnt {
		res.DeviceAll += v.Cnt
		if v.Status == common.StatusUsing {
			res.DeviceUsing += v.Cnt
		}
		if v.Status == common.StatusInactive {
			res.DeviceInactive += v.Cnt
		}
	}
	return res
}
