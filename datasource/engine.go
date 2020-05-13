package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"learnIris/model"
	"xorm.io/xorm"
)

// 数据库
func NewMysqlEngine() *xorm.Engine{
	// 创建数据库连接对象
	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	// 关闭数据库引擎
	defer engine.Close()


	// 设置名称映射规则, 生成的表名为驼峰式命名规则, 默认
	//engine.SetMapper(names.SnakeMapper{})
	err = engine.Sync2(
		new(model.Admin),
		//new(model.User),
	)
	if err != nil {
		panic(err.Error())
	}

	//engine.ShowSQL(true)
	//engine.Logger().SetLevel(0) // debug
	//engine.SetMaxOpenConns(10)
	//engine.SetMaxIdleConns(2)
	return engine
}
