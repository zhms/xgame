// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameXRobot = "x_robot"

// XRobot mapped from table <x_robot>
type XRobot struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                 // id
	SellerID   int32     `gorm:"column:seller_id;comment:运营商" json:"seller_id"`                                // 运营商
	Account    string    `gorm:"column:account;comment:账号" json:"account"`                                     // 账号
	State      int32     `gorm:"column:state;default:1;comment:状态 1离线,2在线" json:"state"`                       // 状态 1离线,2在线
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
}

// TableName XRobot's table name
func (*XRobot) TableName() string {
	return TableNameXRobot
}