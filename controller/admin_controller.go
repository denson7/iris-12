package controller

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"learnIris/model"
	"learnIris/service"
)

type AdminController struct {
	// iris框架自动为每个亲求绑定上下文对象
	Ctx iris.Context

	// admin 功能实体
	Service service.AdminService

	// session对象
	Session *sessions.Session
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// 登录
// 接口：POST /admin/login
func (ac *AdminController) PostLogin(context iris.Context) mvc.Result {
	iris.New().Logger().Info("admin login")
	var adminLogin AdminLogin
	// 获取提交的参数
	ac.Ctx.ReadJSON(&adminLogin)
	// 数据校验
	if adminLogin.UserName == "" && adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": "0",
				"success": "登录失败",
				"message": "用户密码为空,请重新输入",
			},
		}
	}

	// 校验
	admin, exist := ac.Service.GetByAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)
	// 管理员不存在
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": "0",
				"success": "登录失败",
				"message": "用户密码错误,请重新输入",
			},
		}
	}

	// 验证有效,设置session
	userByte, _ := json.Marshal(admin)
	ac.Session.Set(ADMIN, userByte)
	return mvc.Response{
		Object: map[string]interface{}{
			"status": "1",
			"success": "登录成功",
			"message": "登录成功",
		},
	}



}

// 退出登录
// 接口： GET /admin/signout
func (ac *AdminController) GetSignout() mvc.Result {
	// 删除session
	ac.Session.Delete(ADMIN)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  200,
			"message": "",
		},
	}
}

// 获取管理员信息
// 接口：GET /admin/info
func (ac *AdminController) GetInfo() mvc.Result {
	// 从session中获取数据
	userByte := ac.Session.Get(ADMIN)
	if userByte == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  200,
				"message": "",
			},
		}
	}
	// 解析数据到admin结构体
	var admin model.Admin
	err := json.Unmarshal(userByte.([]byte), &admin)
	// 解析失败
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  200,
				"message": "解析失败",
			},
		}
	}
	// 解析成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  200,
			"success": "解析成功",
			"data": admin.AdminToRespDesc(),
		},
	}
}

// 接口： GET /admin/count
func (ac *AdminController) GetCount() mvc.Result {
	count, err := ac.Service.GetAdminCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  200,
				"message": "发生错误",
				"count": 0,
			},
		}
	}
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  200,
			"message": "获取成功",
			"count": count,
		},
	}
}