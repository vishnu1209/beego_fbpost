package controllers

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

// Post API
type PostController struct {
	beego.Controller
}

// @Title GetAllPosts
// @Description Get Post list
// @Success 200 {object} models.post.GetAllPosts
// @Failure 400 no enough input
// @Failure 500 get posts common error
// @Failure 404 Not found
// @Accept json
// @router /post [get]
func (pc *PostController) GetAllPosts() {
	fmt.Println(models.GetAllPosts())
	pc.Data["json"] = models.GetAllPosts()
	pc.ServeJSON()
}

func (pc *PostController) AddNewPost() {
	var PostRequestBody models.PostRequestBody
	err := json.Unmarshal(pc.Ctx.Input.RequestBody, &PostRequestBody)
	if err != nil {
		fmt.Println(err)
	}
	id, err := models.CreateNewPost(&PostRequestBody)

	if err == nil {
		pc.Data["json"] = map[string]int64{"Id": id}
	} else {
		fmt.Println("error")
		pc.Data["json"] = err.Error()
	}
	pc.ServeJSON()

	/*
		To provide output
		data, err := json.MarshalIndent(models.GetAllPosts(),"","   ")
			fmt.Println(string(data), err)

		To read input
		var v interface{}

			json.Unmarshal(pc.Ctx.Input.RequestBody, &v)
			data := v.(map[string]interface{})
			fmt.Println(data)
			for k, v := range data {
				switch v := v.(type) {
				case string:
					fmt.Println(k, v, "(string)")
				case float64:
					fmt.Println(k, v, "(float64)")
				case []interface{}:
					fmt.Println(k, "(array):")
					for i, u := range v {
						fmt.Println("    ", i, u)
					}
				default:
					fmt.Println(k, v, "(unknown)")
				}
			}
	*/

}

// DeleteUser deletes an existing user
func (pc *PostController) DeletePost() {
	// get id from query string and convert it to int
	id, _ := strconv.Atoi(pc.Ctx.Input.Param(":id"))
	fmt.Println(id)

	// delete user
	deleted := models.DeletePost(id)

	// generate response
	pc.Data["json"] = map[string]bool{"deleted": deleted}
	pc.ServeJSON()
}

func (pc *PostController) UpdatePost() {
	var post models.Post
	err := json.Unmarshal(pc.Ctx.Input.RequestBody, &post)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&post)
	p := models.UpdatePost(post)
	pc.Data["json"] = p
	pc.ServeJSON()
}

func (pc *PostController) GetPost() {
	// id := this.Ctx.Input.Param(":id")

	id, _ := strconv.Atoi(pc.Ctx.Input.Param(":id"))
	PostDetails, err := models.GetPostDetails(id)
	data, err := json.MarshalIndent(PostDetails, "", "   ")
	fmt.Println(string(data), err)
	if err != nil {
		fmt.Println(err)
		pc.Data["json"] = err.Error()
	} else {
		pc.Data["json"] = PostDetails
	}
	pc.ServeJSON()
}

func (pc *PostController) GetUserPosts() {
	id, _ := strconv.Atoi(pc.Ctx.Input.Param(":id"))
	UserPostDetails, err := models.GetUserPostDetails(id)
	data, _ := json.MarshalIndent(UserPostDetails, "", "   ")
	fmt.Println(string(data), err)
	if err != nil {
		fmt.Println(err)
		pc.Data["json"] = err.Error()
	} else {
		pc.Data["json"] = UserPostDetails
	}
	pc.ServeJSON()
}
