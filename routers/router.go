package routers

import (
	"beego_training/controller"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
	"net/http"
)

func init()  {
	beego.Any("/**", func(context *context.Context) {
		_, err := context.ResponseWriter.Write([]byte("Welcome to user API"))
		if err != nil {
			log.Println(err)
			return
		}
	})


	beego.Get("/admin", func(c *context.Context) {
		c.ResponseWriter.WriteHeader(http.StatusOK)
		_, err := c.ResponseWriter.Write([]byte("welcome to admin api"))
		if err != nil {
			return
		}
	})

	beego.Router("/user/:userid([0-9]+)", &controller.UserController{}, "Delete:DeleteUser")
	beego.Router("/user/:id([0-9]+)", &controller.UserController{}, "Put:UpdateUser")
	beego.Router("/user", &controller.UserController{}, "Post:CreateUsers;Get:GetUsers")
	beego.AutoRouter(&controller.Security{})

}
