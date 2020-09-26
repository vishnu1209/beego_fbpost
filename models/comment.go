package models

import (
	"github.com/astaxie/beego/orm"
)

type Commento struct {
	Id             int    `json:"Id"`
	CommentedBy    *User  `orm:"rel(fk)"`
	CommentedAt    string `orm:"auto_now_add;type(datetime)"`
	Post           *Post  `orm:"rel(fk)"`
	CommentContent string
}

func init() {
	orm.RegisterModel(new(Commento))
}

type CommentRequestBody struct {
	CommentedById  int
	CommentedAt    string `orm:"auto_now_add;type(datetime)"`
	PostId         int
	CommentContent string
}

func CreateComment(CommentRequestBody *CommentRequestBody) (Id int64, err error) {
	o := orm.NewOrm()
	var user User
	o.QueryTable("User").Filter("Id", CommentRequestBody.CommentedById).One(&user)
	var post Post
	o.QueryTable("Post").Filter("Id", CommentRequestBody.PostId).One(&post)
	o.Read(&user)
	o.Read(&post)
	comment := Commento{CommentContent: CommentRequestBody.CommentContent, CommentedAt: CommentRequestBody.CommentedAt, CommentedBy: &user, Post: &post}
	Id, err = o.Insert(&comment)
	var comments []*Commento
	_, err = o.QueryTable("Commento").All(&comments)
	return
}

func GetAllComments() []*Commento {
	o := orm.NewOrm()
	var commentos []*Commento
	num, err := o.QueryTable("Commento").All(&commentos)
	if num > 0 && err == nil {
		return commentos
	} else {
		return nil
	}
}
