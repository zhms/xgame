// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package order_model

const TableNameXHostSeller = "x_host_seller"

// XHostSeller mapped from table <x_host_seller>
type XHostSeller struct {
	Host     string `gorm:"column:host;primaryKey" json:"host"`
	SellerID int32  `gorm:"column:seller_id" json:"seller_id"`
}

// TableName XHostSeller's table name
func (*XHostSeller) TableName() string {
	return TableNameXHostSeller
}
