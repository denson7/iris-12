package model

import "time"

// 管理员结构体
type Admin struct {
	// 如果field为Id,类型为int64,没有定义tag,xorm会视为主键,并拥有自增属性
	AdminId    int64     `xorm:"pk autoincr" json:"id"` // 主键 自增
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`
	CityName   string    `xorm:"varchar(12)" json:"city_name"`
	CityId     int64     `xorm:"index" json:"city_id"`
	//City *City `xorm:"- <- ->"`
}

func (this *Admin) AdminToRespDesc() interface{} {
	respDesc := map[string]interface{}{
		"user_name":   this.AdminName,
		"id":          this.AdminId,
		"create_time": this.CreateTime,
		"status":      this.Status,
		"avatar":      this.Avatar,
		"city":        this.CityName,
		"admin":       "admin",
	}
	return respDesc
}
