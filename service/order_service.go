package service

import (
	"github.com/kataras/iris/v12"
	"learnIris/model"
	"xorm.io/xorm"
)

type OrderService interface {
	GetCount() (int64, error)
	GetOrderList(offset, limit int) []model.OrderDetail
}

// 结构体adminService需要实现接口AdminService包含的所有方法
func NewOrderService(db *xorm.Engine) OrderService {
	return &orderService{
		engine: db,
	}
}

type orderService struct {
	engine *xorm.Engine
}

// 获取订单列表
func (service *orderService) GetOrderList(offset, limit int) []model.OrderDetail {
	orderList := make([]model.OrderDetail, 0)
	// 查询订单信息
	err := service.engine.Table("user_order").
		Join("INNER", "order_status", "order_status.status_id = user_order.order_status_id").
		Join("INNER", "user", "user.id = user_order.user_id").
		Join("INNER", "shop", "shop.shop_id = user_order.shop_id").
		Join("INNER", "address", "address.address_id = user_order.address_id").
		Find(&orderList)
	if err != nil {
		iris.New().Logger().Info("error getOrderList")
		panic(err.Error())
		return nil
	}
	return orderList
}

func (service *orderService) GetCount() (int64, error) {
	return 1, nil
}
