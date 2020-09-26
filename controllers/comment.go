package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (cc *CommentController) CreateComment() {
	var CommentRequestBody models.CommentRequestBody
	err := json.Unmarshal(cc.Ctx.Input.RequestBody, &CommentRequestBody)
	if err != nil {
		fmt.Println(err)
	}
	if id, err := models.CreateComment(&CommentRequestBody); err == nil {
		cc.Data["json"] = map[string]int64{"CommentId": id}
	} else {
		cc.Data["json"] = err.Error()
	}
	cc.ServeJSON()
}

func (cc *CommentController) GetAllComments() {
	cc.Data["json"] = models.GetAllComments()
	cc.ServeJSON()
}