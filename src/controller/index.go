package controller

import (
	"dujitang/src/service"
	"github.com/kataras/iris"
	"log"
)

func Index(ctx iris.Context) {
	result, err := service.Index()

	if err != nil {
		log.Fatal(err)
	}
	ctx.ViewData("title", "Home page")
	ctx.ViewData("body", result)
	ctx.View("index.html")
}
