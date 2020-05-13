package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"learnIris/config" // mod名的包名
	"learnIris/controller"
	"learnIris/datasource"
	"learnIris/service"
	"time"
)

func main() {
	app := newApp()

	// 配置设置
	configuration(app)
	// 路由设置
	mvcHandle(app)

	// 读取自定义配置
	conf := config.InitConfig()
	addr := ":" + conf.Port

	app.Run(
		// 监听端口
		iris.Addr(addr),
		// 无服务错误提示
		iris.WithoutServerError(iris.ErrServerClosed),
		// 对json数据序列化更快的配置
		iris.WithOptimizations,
	)

}

// 构建App
func newApp() *iris.Application {
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel("debug")

	// 注册静态资源

	// 注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	// http://localhost:8089
	app.Get("/", func(context iris.Context) {
		// 渲染模板文件
		context.View("index.html")
	})
	return app
}

// 配置设置
func configuration(app *iris.Application) {
	// 配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	// 发生错误时配置
	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"errMsg": iris.StatusNotFound,
			"msg":    "not found",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context iris.Context) {
		context.JSON(iris.Map{
			"errMsg": iris.StatusInternalServerError,
			"msg":    "internal error",
			"data":   iris.Map{},
		})
	})
}

// mvc 架构模式处理
func mvcHandle(app *iris.Application) {
	// 启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessionCookie",
		Expires: 24 * time.Hour,
	})

	// 实例化redis
	redis := datasource.NewRedis()
	// 设置session的同步位置为redis
	sessManager.UseDatabase(redis)

	// 实例化mysql数据库
	engine := datasource.NewMysqlEngine()

	// 管理员模块
	adminService := service.NewAdminService(engine)
	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
	)
	admin.Handle(new(controller.AdminController))

	// 统计功能模块
	statisService := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{date}"))
	statis.Register(
		statisService,
		sessManager.Start,
	)
	statis.Handle(new(controller.StatisController))

	// 订单模块
	orderService := service.NewOrderService(engine)
	order := mvc.New(app.Party("/order/"))
	order.Register(
		orderService,
		sessManager.Start,
	)
	order.Handle(new(controller.OrderController))


	//// 用户模块
	//userService := service.NewUserService(engine)
	//user := mvc.New(app.Party("/v1/users"))
	//user.Register(userService, sessManager.Start)
	//user.Handle(new(controller.UserController))
	//

}
