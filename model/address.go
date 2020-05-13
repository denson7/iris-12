package model

type Address struct {
	AddressId     int64  `xorm:"pk autoincr" json:"id"` // 主键自增
	Address       string `json:"address"`               //
	Phone         string `json:"phone"`                 //
	AddressDetail string `json:"address_detail"`        //
	IsValid       int    `json:"is_valid"`              //

}
