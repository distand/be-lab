package do

import (
	"be-lab/model/vo"
	"time"
)

type User struct {
	ID       int32     `json:"-" gorm:"primaryKey"`
	Openid   string    `json:"openid" gorm:"column:openid"`
	Nickname string    `json:"nickname" gorm:"column:nickname"`
	Status   int32     `json:"-" gorm:"column:status"`
	Role     int32     `json:"role" gorm:"column:role"`
	IsDel    int32     `json:"-" gorm:"column:is_del"`
	Ltime    time.Time `json:"ltime" gorm:"column:ltime"`
	Ctime    time.Time `json:"ctime" gorm:"column:ctime"`
}

func (m *User) TableName() string {
	return "lab_user"
}

func (m *User) ToVO() *vo.User {
	return &vo.User{
		ID:       m.ID,
		Openid:   m.Openid,
		Nickname: m.Nickname,
		Role:     m.Role,
		Status:   m.Status,
		Ltime:    m.Ltime.Unix(),
		Ctime:    m.Ctime.Unix(),
	}
}
