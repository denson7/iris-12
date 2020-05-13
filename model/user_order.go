package model

import "time"

type UserOrder struct {
	Id            int64        `xorm:"pk autoincr" json:"id"` // 主键自增
	SunMoney      int64        `xorm:"default 0" json:"sum_money"`
	Time          time.Time    `xorm:"DateTime" json:"time"`         // 时间
	OrderTime     uint64       `json:"order_time"`                   // 订单创建时间
	OrderStatusId int64        `xorm:"index" json:"order_status_id"` // 订单状态Id
	OrderStatus   *OrderStatus `xorm:"-"`                            // 订单对象
	UserId        int64        `xorm:"index" json:"user_id"`         // 用户Id
	User          *User        `xorm:"-"`                            // 用户信息
	ShopId        int64        `xorm:"index" json:"shop_id"`         // 商品编号
	Shop          *Shop        `xorm:"-"`                            // 商品信息
	AddressId     int64        `xorm:"index" json:"address_id"`      // 地址Id
	Address       *Address     `xorm:"-"`                            // 用户地址信息
	DelFlag       int64        `xorm:"default 0" json:"del_flag"`    // 删除标志
}

func (this *UserOrder) UserOrder2Repo() interface{} {
	return ""
}

