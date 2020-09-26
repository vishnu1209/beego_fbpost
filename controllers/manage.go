package controllers
//
//import (
//	"awesomeProject/models"
//	_ "awesomeProject/models"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//	"strconv"
//)
//type UserController struct {
//	beego.Controller
//}
//
//func (manage *ManageController) Delete() {
//	// convert the string value to an int
//	userId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))
//	o := orm.NewOrm()
//	o.Using("default")
//	user := models.User{}
//	if exist := o.QueryTable(user.TableName()).Filter("Id", userId).Exist(); exist {
//		if num, err := o.Delete(&models.User{Id: userId}); err == nil {
//			beego.Info("Record Deleted. ", num)
//		} else {
//			beego.Error("Record couldn't be deleted. Reason: ", err)
//		}
//	} else {
//		beego.Info("Record Doesn't exist.")
//	}
//}