package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id          int
	PostContent string
	PostedBy    *User  `orm:"rel(fk)"`
	PostedAt    string `orm:"auto_now_add;type(datetime)"`
}

type PostRequestBody struct {
	Id          int
	PostedById  int
	PostContent string
	PostedAt    string
}

func init() {
	orm.RegisterModel(new(Post))
}

func GetAllPosts() []*Post {
	o := orm.NewOrm()
	var posts []*Post
	num, err := o.QueryTable("Post").All(&posts)
	if num >= 0 && err != orm.ErrNoRows {
		return posts
	} else {
		return nil
	}
}

func CreateNewPost(PostRequestBody *PostRequestBody) (Id int64, err error) {
	beego.Info("Creating new post")
	o := orm.NewOrm()
	var user User
	o.QueryTable("User").Filter("Id", PostRequestBody.PostedById).One(&user)
	o.Read(&user)
	post := Post{Id: PostRequestBody.Id, PostContent: PostRequestBody.PostContent, PostedAt: PostRequestBody.PostedAt, PostedBy: &user}
	Id, err = o.Insert(&post)
	var posts []*Post
	num, err := o.QueryTable("Post").All(&posts)
	fmt.Println(num)
	return
}

func UpdatePost(post Post) *Post {
	o := orm.NewOrm()
	p := Post{Id: post.Id}
	var UpdatedPost Post

	if o.Read(&p) == nil {
		p = post
		_, err := o.Update(&p)
		if err == nil {
			UpdatedPost = Post{Id: post.Id}
			o.Read(&UpdatedPost)
		}
	}
	return &UpdatedPost
}

// DeleteUser deletes a user
func DeletePost(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Post{Id: id})
	if err == nil {
		// successfull
		return true
	}

	return false
}

type PostDetails struct {
	PostId        int           `json:"post_id"`
	PostedBy      User          `json:"posted_by"`
	PostedAt      string        `json:"posted_at"`
	PostContent   string        `json:"post_content"`
	Reactions     Reactions     `json:"reactions"`
	Comments      []CommentDict `json:"comments"`
	CommentsCount int64         `json:"comments_count"`
}

type Reactions struct {
	Count     int64    `json:"count"`
	Reactions []string `json:"reactions"`
}

type CommentDict struct {
	Id             int         `json:"comment_id"`
	CommentedBy    User        `json:"commenter"`
	CommentedAt    string      `json:"commented_at"`
	CommentContent string      `json:"comment_content"`
	ReactionsDict  Reactions   `json:"reactions"`
	Replies        []ReplyDict `json:"replies"`
	RepliesCount   int64       `json:"replies_count"`
}

type ReplyDict struct {
	Id           int64  `json:"comment_id"`
	RepliedBy    User   `json:"commenter"`
	RepliedAt    string `json:"commented_at"`
	ReplyContent string `json:"comment_content"`
}

func GetPostDetails(id int) (PostDetails, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Post")
	var post Post
	InvalidPostId := qs.Filter("Id", id).One(&post)
	if InvalidPostId != nil {
		return PostDetails{}, InvalidPostId
	}
	qs = o.QueryTable("User")
	var user User
	InvalidUserId := qs.Filter("Id", post.PostedBy.Id).One(&user)
	if InvalidUserId != nil {
		return PostDetails{}, InvalidUserId
	}
	var PostReactionsList []*Reaction
	var ReactionsList []string
	PostReactionsCount, _ := o.QueryTable("Reaction").Filter("Post", post).All(&PostReactionsList, "ReactionType")
	for i := 0; i < len(PostReactionsList); i++ {
		ReactionsList = append(ReactionsList, PostReactionsList[i].ReactionType)
	}

	var comments []*Commento
	qs = o.QueryTable("Commento")
	CommentsCount, _ := qs.Filter("Post", post).All(&comments, "Id", "CommentedBy", "CommentedAt", "CommentContent")

	ListOfComments := []CommentDict{}
	for i := 0; i < len(comments); i++ {
		beego.Info("Comment dict")
		var CommentReactions []*Reaction
		var CommentReactionTypesList []string
		CommentReactionCount, _ := o.QueryTable("Reaction").Filter("Comment", comments[i]).All(&CommentReactions, "ReactionType")
		for i := 0; i < len(CommentReactions); i++ {
			beego.Info("Comment Reactions")
			CommentReactionTypesList = append(CommentReactionTypesList, CommentReactions[i].ReactionType)
		}
		CommentReactionsDict := Reactions{Count: CommentReactionCount, Reactions: CommentReactionTypesList}
		var CommentedUser User
		qs = o.QueryTable("User")
		qs.Filter("Id", comments[i].CommentedBy.Id).One(&CommentedUser)

		var CommentReplies []*Reply
		CommentRepliesCount, _ := o.QueryTable("Reply").Filter("RepliedFor", comments[i].Id).All(&CommentReplies)
		CommentRepliesList := []ReplyDict{}
		for ReplyNumber := 0; ReplyNumber < len(CommentReplies); ReplyNumber++ {
			var RepliedBy User
			o.QueryTable("User").Filter("Id", CommentReplies[ReplyNumber].RepliedBy.Id).One(&RepliedBy)
			beego.Info("Comment Replies")

			ReplyD := ReplyDict{
				Id:           CommentReplies[ReplyNumber].Id,
				RepliedBy:    RepliedBy,
				RepliedAt:    CommentReplies[ReplyNumber].RepliedAt,
				ReplyContent: CommentReplies[ReplyNumber].ReplyContent,
			}
			CommentRepliesList = append(CommentRepliesList, ReplyD)
		}

		CommentD := CommentDict{
			Id:             comments[i].Id,
			CommentedBy:    CommentedUser,
			CommentedAt:    comments[i].CommentedAt,
			CommentContent: comments[i].CommentContent,
			ReactionsDict:  CommentReactionsDict,
			Replies:        CommentRepliesList,
			RepliesCount:   CommentRepliesCount,
		}
		ListOfComments = append(ListOfComments, CommentD)
	}

	PostReactionsDict := Reactions{Count: PostReactionsCount, Reactions: ReactionsList}
	PostDict := PostDetails{
		PostId:        post.Id,
		PostedBy:      user,
		PostedAt:      post.PostedAt,
		PostContent:   post.PostContent,
		Reactions:     PostReactionsDict,
		Comments:      ListOfComments,
		CommentsCount: CommentsCount,
	}
	return PostDict, nil
}

func GetUserPostDetails(id int) ([]PostDetails, error) {
	o := orm.NewOrm()
	var user User
	InvalidUserId := o.QueryTable("User").Filter("Id", id).One(&user)
	fmt.Println(InvalidUserId)
	if InvalidUserId != nil {
		return []PostDetails{}, InvalidUserId
	}
	var posts []*Post
	o.QueryTable("Post").Filter("PostedBy", user).All(&posts)
	var UserPostDetails []PostDetails
	for i := 0; i < len(posts); i++ {
		PostDetails, _ := GetPostDetails(posts[i].Id)
		UserPostDetails = append(UserPostDetails, PostDetails)
	}
	return UserPostDetails, nil
}
