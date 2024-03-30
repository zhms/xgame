// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package order_model

import (
	"time"
)

const TableNameXUser = "x_user"

// XUser mapped from table <x_user>
type XUser struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SellerID        int32     `gorm:"column:seller_id;comment:运营商" json:"seller_id"`                                // 运营商
	Account         string    `gorm:"column:account;comment:用户账号" json:"account"`                                   // 用户账号
	Password        string    `gorm:"column:password;comment:用户密码" json:"password"`                                 // 用户密码
	IsVisitor       int32     `gorm:"column:is_visitor;comment:是否是游客" json:"is_visitor"`                            // 是否是游客
	State           int32     `gorm:"column:state;default:1;comment:状态 1开启,2关闭" json:"state"`                       // 状态 1开启,2关闭
	Agent           string    `gorm:"column:agent;comment:所属管理员" json:"agent"`                                      // 所属管理员
	Token           string    `gorm:"column:token;comment:token" json:"token"`                                      // token
	LoginIP         string    `gorm:"column:login_ip;comment:登录ip" json:"login_ip"`                                 // 登录ip
	LoginIPLocation string    `gorm:"column:login_ip_location;comment:登录ip地区" json:"login_ip_location"`             // 登录ip地区
	LoginCount      int32     `gorm:"column:login_count;comment:登录次数" json:"login_count"`                           // 登录次数
	LoginTime       time.Time `gorm:"column:login_time;comment:登录时间" json:"login_time"`                             // 登录时间
	ChatState       int32     `gorm:"column:chat_state;default:1;comment:禁言 1是,2否" json:"chat_state"`               // 禁言 1是,2否
	IsOnline        int32     `gorm:"column:is_online;default:2;comment:是否在线" json:"is_online"`                     // 是否在线
	CreateTime      time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
}

// TableName XUser's table name
func (*XUser) TableName() string {
	return TableNameXUser
}