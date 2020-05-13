package datasource

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"learnIris/config"
	"time"
)

func NewRedis() *redis.Database {
	var db *redis.Database

	// 项目配置
	cmsConfig := config.InitConfig()
	if cmsConfig != nil {
		iris.New().Logger().Info("hello")
		rd := cmsConfig.Redis
		iris.New().Logger().Info(rd)

		db = redis.New(redis.Config{
			Network:   rd.NetWork, // "tcp",
			Addr:      rd.Addr + ":" + rd.Port, // "127.0.0.1:6379",
			Timeout:   time.Duration(30) * time.Second,
			MaxActive: 10,
			Password:  rd.Password,
			Database:  "",
			Prefix:    "",
			Delim:     "-",
			Driver:    redis.Redigo(), // redis.Radix() can be used instead.
		})
		// 当执行 control+C/cmd+C  时关闭连接
		iris.RegisterOnInterrupt(func() {
			db.Close()
		})
	} else {
		iris.New().Logger().Info("redis connect error")
	}
	return db
}
