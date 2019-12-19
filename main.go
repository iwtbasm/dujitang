package main

import (
	"dujitang/src/controller"
	"github.com/kataras/iris"
)

func newApp() *iris.Application {
	app := iris.New()
	app.StaticEmbedded("/static", "./assets", Asset, AssetNames)
	return app
}

func main() {
	//app := iris.New()
	app := newApp()
	// 静态资源
	// app.StaticWeb("./public", "/")
	// 网页模板
	tmpl := iris.HTML("./views", ".html")

	app.RegisterView(tmpl)

	app.Get("/", controller.Index)

	app.Run(iris.Addr(":8088"))

}
