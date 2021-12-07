package controller

import "github.com/astaxie/beego"

type Security struct {
	beego.Controller
}

func (s *Security) Login()  {
	s.Ctx.ResponseWriter.Write([]byte("login controller"))
}

func (s *Security) Logout()  {
	s.Ctx.ResponseWriter.Write([]byte("logout controller"))
}


func (s *Security) Authenticate()  {
	s.Ctx.ResponseWriter.Write([]byte("authenticate controller"))
}

