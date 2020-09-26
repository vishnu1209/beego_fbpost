package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
)

type Reaction struct {
	Id           int
	Post         *Post     `orm:"null;rel(fk);on_delete(cascade)"`
	Comment      *Commento `orm:"null;rel(fk);on_delete(cascade)"`
	User         *User     `orm:"rel(fk);on_delete(cascade)"`
	ReactionType string
	ReactedAt    string `orm:"auto_now_add;type(datetime)"`
}

type ReactionRequestBody struct {
	PostId       int
	CommentId    int
	UserId       int
	ReactionType string
}

func init() {
	fmt.Println("type:", reflect.TypeOf(Reaction{}))
	orm.RegisterModel(new(Reaction))
}

func CreateReaction(reaction *ReactionRequestBody) (Id int64, err error) {
	o := orm.NewOrm()
	var user User
	o.QueryTable("User").Filter("Id", reaction.UserId).One(&user)
	if reaction.PostId != 0 {
		var post Post
		o.QueryTable("Post").Filter("Id", reaction.PostId).One(&post)
		reaction := Reaction{
			Post:         &post,
			Comment:      nil,
			User:         &user,
			ReactionType: reaction.ReactionType,
		}
		Id, err = o.Insert(&reaction)
	} else {
		var comment Commento
		o.QueryTable("Commento").Filter("Id", reaction.CommentId).One(&comment)
		reaction := Reaction{
			Post:         nil,
			Comment:      &comment,
			User:         &user,
			ReactionType: reaction.ReactionType,
		}
		Id, err = o.Insert(&reaction)
	}
	return
}

func GetAllReactions() []*Reaction {
	o := orm.NewOrm()
	var reactions []*Reaction
	num, err := o.QueryTable("Reaction").All(&reactions)
	fmt.Println(reactions)
	if num > 0 && err == nil {
		return reactions
	} else {
		return nil
	}
}

func TotalReactionsCount() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("Reaction").Count()
	return
}

func ReactionMetricsForGivenPost(PostId int) (reactions map[string]int, err error) {
	o := orm.NewOrm()
	var ReactionsOfPost []*Reaction
	qs := o.QueryTable("Reaction")
	_, err = qs.Filter("Post", PostId).All(&ReactionsOfPost)
	if err != nil {
		return
	}
	reactions = make(map[string]int)
	for i := 0; i < len(ReactionsOfPost); i++ {
		reactions[ReactionsOfPost[i].ReactionType] += 1
	}
	return
}

type PostReaction struct {
	UserId   int    `json:"user_id"` // to change attribute names in output
	PostId   int    `json:"post_id"`
	Reaction string `json:"reaction_type"`
}

func GetReactionDetailsOfPost(id int) []PostReaction {
	o := orm.NewOrm()
	qs := o.QueryTable("Reaction")
	var reactions []*Reaction
	qs.Filter("Post", id).All(&reactions, "Id", "Post", "ReactionType", "User")
	ReactionsList := []PostReaction{}
	for i := 0; i < len(reactions); i++ {
		d := PostReaction{
			UserId:   reactions[i].User.Id,
			PostId:   reactions[i].Post.Id,
			Reaction: reactions[i].ReactionType,
		}
		ReactionsList = append(ReactionsList, d)
	}
	return ReactionsList
}