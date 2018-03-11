package main

import (
	"fmt"
	"time"
	"./control"
	"github.com/kataras/iris"
)
const (
	user = "postgres"
	password = "123"
	database = "postgres"
)

func main() {
	fmt.Println(time.Now().Local().Format("2006-01-02"))
	control.ConnDb(user, password, database) 
 
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html").Reload(true))

	app.Get("/", func (ctx iris.Context) {
		ctx.View("menu.html")
	})
	app.Get("/post", control.GetAllPost)
	app.Get("/create", control.CreatePost)
	app.Post("/created", control.CreatedPost)
	app.Get("/preview", control.GetPublishPost)
	app.Get("/post/{id}",control.DetailsPost)
	app.Get("/post/edit/{id}", control.Edit)
	app.Post("/post/edited/{id}", control.Edited)
	app.Get("/post/delete/{id}", control.DeletePost)

	app.Get("/api", control.Api)

	app.Run(iris.Addr(":8080"))
}
