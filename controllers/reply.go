package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ReplyController struct {
	beego.Controller
}

func (rc *ReplyController) CreateReply() {
	var ReplyRequest models.ReplyRequestBody
	err := json.Unmarshal(rc.Ctx.Input.RequestBody, &ReplyRequest)
	if err != nil {
		rc.Data["json"] = "Invalid Json"
		rc.ServeJSON()
	}
	id, err := models.CreateReply(&ReplyRequest)
	rc.Data["json"] = map[string]int64{"reply_id": id}
	rc.ServeJSON()
}

func (rc *ReplyController) GetRepliesOfComment() {
	id, _ := strconv.Atoi(rc.Ctx.Input.Param(":CommentId"))
	fmt.Println(id, 9)
	rc.Data["json"] = models.GetReplies(id)
	rc.ServeJSON()
}
