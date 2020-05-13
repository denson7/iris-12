package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"learnIris/service"
	"strings"
)

type StatisController struct {
	// iris框架自动为每个亲求绑定上下文对象
	Ctx iris.Context

	// 统计功能实体
	Service service.StatisService

	// session对象
	Session *sessions.Session
}

func (sc *StatisController) GetCount() mvc.Result {
	// eg: /statis/user/2020-05-04/count
	path := sc.Ctx.Path()
	var pathSlice []string
	if path != "" {
		pathSlice = strings.Split(path, "/")
	}

	// 不符合请求格式
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": 200,
				"count":  0,
			},
		}
	}

	pathSlice = pathSlice[1:]
	model := pathSlice[1]
	date := pathSlice[2]
	var result int64
	switch model {
	case "user":
		userResult := sc.Session.Get("username" + date)
		if userResult != nil {
			userResult = userResult.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": 200,
					"count":  userResult,
				},
			}
		} else {
			iris.New().Logger().Info(date)
			result = sc.Service.GetUserDailyCount(date)
			// 设置缓存
			sc.Session.Set("usrename"+date, result)
		}
	case "order":
		orderResult := sc.Session.Get("order" + date)
		if orderResult != nil {
			orderResult = orderResult.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": 200,
					"count":  orderResult,
				},
			}
		} else {
			result = sc.Service.GetOrderDailyCount(date)
			sc.Session.Set("order"+date, result)
		}
	case "admin":
		adminResult := sc.Session.Get("admin" + date)
		if adminResult != nil {
			adminResult = adminResult.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": 200,
					"count":  adminResult,
				},
			}
		} else {
			result = sc.Service.GetAdminDailyCount(date)
			sc.Session.Set("admin"+date, result)
		}
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"status": 200,
			"count": result,
		},
	}

}
