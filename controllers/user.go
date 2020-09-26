package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

// UserController represents controller for user API
type UserController struct {
	beego.Controller
}

// GetAllUsers function gets all the users
func (uc *UserController) GetAllUsers() {
	fmt.Println(models.GetAllUsers())
	uc.Data["json"] = models.GetAllUsers()
	uc.ServeJSON()
}

//
//func ValidateRequestData(c *gin.Context) {
//	var v models.User
//	if err := uc.ShouldBindJSON(&v); err != nil {
//		v.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
//		return
//	}
//}

// AddNewUser adds new user
func (uc *UserController) AddNewUser() {
	var user models.User
	fmt.Println(string(uc.Ctx.Input.RequestBody), user)
	err:=json.Unmarshal(uc.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(1, user, 1)
	if id, err := models.InsertOneUser(&user); err == nil {
		uc.Data["json"] = map[string]int64{"UserId": id}
	} else {
		uc.Data["json"] = err.Error()
	}
	uc.ServeJSON()
}

//var user = models.User{}
//fmt.Println(string(uc.Ctx.Input.RequestBody))
//id:= uc.GetString("FirstName")
//json.Unmarshal(uc.Ctx.Input.RequestBody, &user)
//fmt.Println(0, id, 2)
//fmt.Println(&user, user)
//user_id := models.InsertOneUser(&user)
//uc.Data["json"] = map[string]interface{}{"userid":user_id, "FirstName": user.FirstName}
//uc.ServeJSON()
//return
//request := this.Ctx.Request
//this.Ctx.WriteString(request.FirstName)
//jsoninfo := this.GetString("FirstName")
//if jsoninfo == "" {
//	this.Ctx.WriteString("jsoninfo is empty")
//	return
//}
//this.Ctx.WriteString("Id")
//o := orm.NewOrm()
//var user models.User
//user.Id = 2
//user.FirstName = "slene"
//id, err := o.Insert(&user)
//if err == nil {
//	fmt.Println(id)
//}

//UpdateUser updates an existing user
func (uc *UserController) UpdateUser() {

	var u models.User
	err := json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	if err != nil {
		fmt.Println(err)
	}
	user := models.UpdateUser(u)
	uc.Data["json"] = user
	uc.ServeJSON()
}

// DeleteUser deletes an existing user
func (uc *UserController) DeleteUser() {
	// get id from query string and convert it to int
	id, _ := strconv.Atoi(uc.Ctx.Input.Param(":id"))
	fmt.Println(id)

	// delete user
	deleted := models.DeleteUser(id)

	// generate response
	uc.Data["json"] = map[string]bool{"deleted": deleted}
	uc.ServeJSON()
}

//
//// GetUserById gets a single user with the given id
//func (uc *UserController) GetUserById() {
//	// get the id from query string
//	id, _ := strconv.Atoi(uc.Ctx.Input.Param(":id"))
//
//	// get user
//	user := models.GetUserById(id)
//
//	// generate response
//	uc.Data["json"] = user
//	uc.ServeJSON()
//}
