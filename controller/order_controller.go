package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"learnIris/service"
)

type OrderController struct {
	Ctx iris.Context
	Service service.OrderService
	Session *sessions.Session
}

// 获取订单列表
func (oc *OrderController) GetOrderList() mvc.Result {
	iris.New().Logger().Info("查询订单列表")
	offsetStr := oc.Ctx.FormValue("offset")
	limitStr := oc.Ctx.FormValue("limit")
	var offset int
	var limit int  
	// 判断offset和limit不能为""
	if offsetStr == "" || limitStr == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": 600,
				"type": "",
				"message": "参数错误",
			},
		}
	}
	//
	orderList := oc.Service.GetOrderList(offset, limit)
	if len(orderList) == 0 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": 200,
				"type": "",
				"message": "没有数据",
			},
		}
	}
	// 将数据返回前台
	var respList []interface{}
	for _, detail := range orderList{
		respList = append(respList, detail.OrderDetail2Resp)
	}
	return mvc.Response{
		Object: &respList,
	}
}