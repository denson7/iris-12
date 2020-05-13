package model

type OrderDetail struct {
	UserOrder   `xorm:"extends"`
	User        `xorm:"extends"`
	OrderStatus `xorm:"extends"`
	Shop        `xorm:"extends"`
	Address     `xorm:"extends"`
}

func (detail *OrderDetail) OrderDetail2Resp() interface{} {
	respDesc := map[string]interface{}{
		"id":           detail.UserOrder.Id,
		"total_amount": detail.UserOrder.SunMoney,
		"user_id":      detail.User.Id,
		"status":       detail.OrderStatus.StatusDesc,
		"address_id":   detail.Address.AddressId,
	}
	return respDesc
}
