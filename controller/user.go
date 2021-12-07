package controller

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
	"net/http"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Name string
	Age int
}

func (u *UserController) GetUsers()  {
  //values := u.Input()

  //fmt.Println(values.Get("name"))
	//fmt.Println(values.Get("age"))

 // fmt.Println(u.GetString("name"))
	//fmt.Println(u.GetInt("age"))
 // fmt.Println(values.Get("i"))

	//names := u.GetStrings("name")

	//fmt.Println(names)

 // fmt.Println(values.Get("name"))
  var user User
  u.Ctx.Input.Bind(&user, "user")

  fmt.Println(user.Age, user.Name)
	u.Ctx.ResponseWriter.Write([]byte("Get users"))



}

type Address struct {
	City []string
	Zip int
}

type Person struct {
	UserName *string `json:"username" valid:"Required"`
	FirstName string `json:"firstname" valid:"Required;MaxSize(10)"`
	Age int  `json:"age" valid:"Required;Min(18)"`
	Address map[string]interface{} `json:"preferences" valid:"Required"`
}

func (u *UserController) CreateUsers()  {
	var user Person
	fmt.Println(string(u.Ctx.Input.RequestBody))
 err :=   json.Unmarshal(u.Ctx.Input.RequestBody,&user)

 if err!=nil{
	 log.Println(err)
	 u.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	 u.Ctx.ResponseWriter.Write([]byte("request parameters are not valid"))
	 return
 }


  if user.UserName == nil {
	  u.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	  return
  }

 valid := validation.Validation{}



  isValid, _ := valid.Valid(&user)
	fmt.Println(user)
	fmt.Println(valid.ErrorMap())

  if !isValid {
	  u.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	 res, _ := json.Marshal(valid.ErrorMap())
	  u.Ctx.ResponseWriter.Write(res)
	  return
  }

fmt.Println(isValid)

	fmt.Println(user.UserName)


	u.Data["json"] = user
	u.Ctx.ResponseWriter.WriteHeader(http.StatusCreated)
	u.ServeJSON()

}

func (u *UserController) UpdateUser()  {

	_, err := u.Ctx.ResponseWriter.Write([]byte("update user"))
	if err != nil {
		return
	}
}


func (u *UserController) DeleteUser()  {
   value := u.Input()
   fmt.Println(value.Get(":userid"))
	_, err := u.Ctx.ResponseWriter.Write([]byte(""))
	if err != nil {
		return
	}
}
