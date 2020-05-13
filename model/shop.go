package model

type Shop struct {
	ShopId int64 `xorm:"pk autoincr" json:"item_id"` // 主键自增
	Name string `xorm:"varchar(32)" json:"name"` //
	Address string `xorm:"varchar(128)" json:"address"` // 店铺地址
	Description string `xorm:"varchar(255)" json:"description"` // 店铺描述
	Phone string `json:"phone"` // 店铺电话
	New bool `json:"new"` // 新开店铺
	IsPremium bool `json:"is_premium"` // 品牌保障

}
