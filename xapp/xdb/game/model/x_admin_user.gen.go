// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameXAdminUser = "x_admin_user"

// XAdminUser mapped from table <x_admin_user>
type XAdminUser struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增Id" json:"id"`               // 自增Id
	SellerID    int32     `gorm:"column:seller_id;comment:运营商" json:"seller_id"`                                // 运营商
	Account     string    `gorm:"column:account;comment:账号" json:"account"`                                     // 账号
	Password    string    `gorm:"column:password;comment:登录密码" json:"password"`                                 // 登录密码
	RoleName    string    `gorm:"column:role_name;comment:角色" json:"role_name"`                                 // 角色
	LoginGoogle string    `gorm:"column:login_google;comment:登录谷歌验证码" json:"login_google"`                      // 登录谷歌验证码
	OptGoogle   string    `gorm:"column:opt_google;comment:渠道商" json:"opt_google"`                              // 渠道商
	Agent       string    `gorm:"column:agent;comment:上级代理" json:"agent"`                                       // 上级代理
	State       int32     `gorm:"column:state;default:1;comment:状态 1开启,2关闭" json:"state"`                       // 状态 1开启,2关闭
	Token       string    `gorm:"column:token;comment:最后登录的token" json:"token"`                                 // 最后登录的token
	LoginCount  int32     `gorm:"column:login_count;comment:登录次数" json:"login_count"`                           // 登录次数
	LoginTime   time.Time `gorm:"column:login_time;default:CURRENT_TIMESTAMP;comment:最后登录时间" json:"login_time"` // 最后登录时间
	LoginIP     string    `gorm:"column:login_ip;comment:最后登录Ip" json:"login_ip"`                               // 最后登录Ip
	Memo        string    `gorm:"column:memo;comment:备注" json:"memo"`                                           // 备注
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	RoomID      int32     `gorm:"column:room_id" json:"room_id"`
}

// TableName XAdminUser's table name
func (*XAdminUser) TableName() string {
	return TableNameXAdminUser
}