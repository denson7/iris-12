package model

type OrderStatus struct {
	StatusId   int64  `xorm:"pk autoincr" json:"id"`
	StatusDesc string `xorm:"varchar(255)"` // 未支付, 已支付, 已发货,正在配送,已接收,发起退款,正在退款,退款成功,取消订单
}
