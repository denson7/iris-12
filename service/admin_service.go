package service

import (
	"learnIris/model"
	"xorm.io/xorm"
)

type AdminService interface {
	// 通过用户名和密码获取管理员实体
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)

	// 获取管理员总数
	GetAdminCount() (int64, error)
}

// 结构体adminService需要实现接口AdminService包含的所有方法
func NewAdminService(db *xorm.Engine) AdminService  {
	return &adminService{
		engine: db,
	}
}

// 管理员的服务实现结构体
type adminService struct {
	engine *xorm.Engine
}

// 查询管理员的总数
// GetAdminCount()的具体实现
func (ac *adminService) GetAdminCount() (int64, error) {
	count, err := ac.engine.Count(new(model.Admin))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	return count, nil
}

// 通过用户名和密码查询管理员
// GetByAdminNameAndPassword()的具体实现
func (ac *adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin
	ac.engine.Where("admin_name = ? and pwd = ? ", username, password).Get(&admin)
	return admin, admin.AdminId != 0
}