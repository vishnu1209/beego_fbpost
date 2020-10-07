package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Reply struct {
	Id           int64
	RepliedBy    *User     `orm:"rel(fk)"`
	RepliedAt    string    `orm:"auto_now_add;type(datetime)"`
	RepliedFor   *Commento `orm:"rel(fk)"`
	ReplyContent string
}

type ReplyRequestBody struct {
	Id           int64
	RepliedById  int
	RepliedForId int
	ReplyContent string
}

func CreateReply(body *ReplyRequestBody) (id int64, err error) {
	o := orm.NewOrm()
	var user User
	err = o.QueryTable("User").Filter("Id", body.RepliedById).One(&user)
	if err != nil {
		return
	}
	var comment Commento
	err = o.QueryTable("Commento").Filter("Id", body.RepliedForId).One(&comment)
	if err != nil {
		return 0, err
	}
	reply := Reply{Id: body.Id, RepliedBy: &user, RepliedFor: &comment, ReplyContent: body.ReplyContent}
	id, err = o.Insert(&reply)
	return
}

func GetReplies(id int) []*Reply {
	o := orm.NewOrm()
	var comment Commento
	err := o.QueryTable("Commento").Filter("Id", id).One(&comment)
	fmt.Println(comment)
	if err != nil {
		return nil
	}
	var replies []*Reply
	_, err = o.QueryTable("Reply").Filter("RepliedFor", comment).All(&replies)
	fmt.Println(replies)
	return replies
}

func init() {
	orm.RegisterModel(new(Reply))
}
