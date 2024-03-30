// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package game_model

import (
	"time"
)

const TableNameXLiveProvider = "x_live_provider"

// XLiveProvider mapped from table <x_live_provider>
type XLiveProvider struct {
	SellerID   int32     `gorm:"column:seller_id;comment:运营商" json:"seller_id"`                                // 运营商
	Provider   string    `gorm:"column:provider;comment:提供商" json:"provider"`                                  // 提供商
	PushURL    string    `gorm:"column:push_url;comment:推流地址" json:"push_url"`                                 // 推流地址
	PushKey    string    `gorm:"column:push_key;comment:推流鉴权key" json:"push_key"`                              // 推流鉴权key
	PullURL    string    `gorm:"column:pull_url;comment:拉流地址" json:"pull_url"`                                 // 拉流地址
	PullKey    string    `gorm:"column:pull_key;comment:拉流鉴权key" json:"pull_key"`                              // 拉流鉴权key
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
}

// TableName XLiveProvider's table name
func (*XLiveProvider) TableName() string {
	return TableNameXLiveProvider
}