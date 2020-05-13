package model

import "time"

type User struct {
	Id int64 `xorm:"pk autoincr" json:"id"` // 主键自增
	UserName string `xorm:"varchar(12)" json:"username"` // 用户名
	RegisterTime time.Time `json:"register_time"` // 注册时间
	Mobile string `xorm:"varchar(11)" json:"mobile"` // 手机号
	IsActive int64 `json:"is_active"` // 用户是否激活
	Balance int64 `json:"balance"` // 账户余额
	Pwd string `json:"password"` // 密码
	DelFlag int64 `json:"del_flag"` // 是否删除
	CityName string `xorm:"varchar(12)" json:"city_name"`
	//City *City `xorm:"- <- ->"`
}
