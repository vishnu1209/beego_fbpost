package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ReactionController struct {
	beego.Controller
}

func (rc *ReactionController) GetAllReactions() {
	rc.Data["json"] = models.GetAllReactions()
	rc.ServeJSON()
}

func (rc *ReactionController) TotalReactionsCount() {
	num, err := models.TotalReactionsCount()
	if err == nil {
		rc.Data["json"] = map[string]int64{"number of reactions": num}
	} else {
		fmt.Println("error")
		rc.Data["json"] = err.Error()
	}
	rc.ServeJSON()
}

func (rc *ReactionController) CreateReaction() {
	var Reaction models.ReactionRequestBody
	fmt.Println(string(rc.Ctx.Input.RequestBody), 1)
	err := json.Unmarshal(rc.Ctx.Input.RequestBody, &Reaction)
	if err != nil {
		fmt.Println(err)
	}
	id, err := models.CreateReaction(&Reaction)
	if err == nil {
		rc.Data["json"] = map[string]int64{"Id": id}
	} else {
		fmt.Println("error")
		rc.Data["json"] = err.Error()
	}
	rc.ServeJSON()
}

func (rc *ReactionController) GetReactionMetrics() {
	PostId, _ := strconv.Atoi(rc.Ctx.Input.Param(":id"))
	fmt.Println(PostId)
	reactions, err := models.ReactionMetricsForGivenPost(PostId)
	if err != nil {
		rc.Data["json"] = err.Error()
	} else {
		rc.Data["json"] = reactions
	}
	rc.ServeJSON()
}

func (rc *ReactionController) GetPostReactions() {
	id, _ := strconv.Atoi(rc.Ctx.Input.Param(":id"))
	rc.Data["json"] = models.GetReactionDetailsOfPost(id)
	rc.ServeJSON()
}
