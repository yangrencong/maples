package entity

import (
	"time"
)

type User struct {
	Id         int       `gorm:"column:id;type:int(11);AUTO_INCREMENT;NOT NULL" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(20);comment:用户名;NOT NULL" json:"name"`
	Sex        int       `gorm:"column:sex;type:int(11);default:1;comment:性别，男1，女2;NOT NULL" json:"sex"`
	Phone      string    `gorm:"column:phone;type:varchar(20);comment:电话号;NOT NULL" json:"phone"`
	Birthday   string    `gorm:"column:birthday;type:varchar(30);default:0000-00-00" json:"birthday"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP" json:"update_time"`
}

func (m *User) TableName() string {
	return "user"
}
