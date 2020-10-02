// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"awesomeProject/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		//if strings.HasPrefix(ctx.Input.URL(), "/login/" ) {
		//	return
		//}
		if strings.HasPrefix(ctx.Input.URL(), "/login/") ||
			(ctx.Request.RequestURI == "/") || (ctx.Request.RequestURI == "/authorize/") {
			// just do nothing in the filter and complete the logic in controller
			return
		}

		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(307, "/authorize/")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	// hello-world route
	beego.Router("/", &controllers.HelloController{}, "get:HelloWorld")
	// Login
	beego.Router("/login/", &controllers.LoginController{}, "post:Login")
	// authorize
	beego.Router("/authorize/", &controllers.LoginController{}, "post:VerifyToken")
	// Logout
	//beego.Router("/logout", &controllers.LoginController{}, "post:Logout")

	// init namespace
	ns := beego.NewNamespace("/api/v1",

		beego.NSNamespace("/user",
			// get all users
			beego.NSRouter("/", &controllers.UserController{}, "get:GetAllUsers"),

			// add new user
			beego.NSRouter("/", &controllers.UserController{}, "post:AddNewUser"),

			//update an existing user
			beego.NSRouter("/", &controllers.UserController{}, "put:UpdateUser"),

			//// delete a user
			beego.NSRouter("/:id", &controllers.UserController{}, "delete:DeleteUser"),
			//// Authenticate User
			//beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			//// get a user with id
			//beego.NSRouter("/:id", &controllers.UserController{}, "get:GetUserById"),
		),
		beego.NSNamespace("/post",
			//Get All Posts
			beego.NSRouter("/", &controllers.PostController{}, "get:GetAllPosts"),
			//Create New Post
			beego.NSRouter("/", &controllers.PostController{}, "post:AddNewPost"),
			//update New Post
			beego.NSRouter("/", &controllers.PostController{}, "put:UpdatePost"),
			// Delete Post
			beego.NSRouter("/:id", &controllers.PostController{}, "delete:DeletePost"),
			// Get post
			beego.NSRouter("/:id", &controllers.PostController{}, "get:GetPost"),
			//get user posts
			beego.NSRouter(":id/posts/", &controllers.PostController{}, "get:GetUserPosts"),
		),
		beego.NSNamespace("/commento",
			//Create Comment
			beego.NSRouter("/", &controllers.CommentController{}, "post:CreateComment"),
			//Get All Comments
			beego.NSRouter("/", &controllers.CommentController{}, "get:GetAllComments"),
		),
		beego.NSNamespace("/reaction",
			// Create Reaction
			beego.NSRouter("/", &controllers.ReactionController{}, "post:CreateReaction"),
			// Get All Reactions
			beego.NSRouter("/", &controllers.ReactionController{}, "get:GetAllReactions"),
			// total reactions count
			beego.NSRouter("/count", &controllers.ReactionController{}, "get:TotalReactionsCount"),
			//get reaction metrics
			beego.NSRouter("/:id/metrics", &controllers.ReactionController{}, "post:GetReactionMetrics"),
			//get Post Reactions
			beego.NSRouter(":id/reactions", &controllers.ReactionController{}, "get:GetPostReactions"),
		),
		beego.NSNamespace("/error",
			// Create Error
			beego.NSRouter("/", &controllers.ErrorController{}, "get:ErrorHandling"),
		),
	)
	// register namespace
	beego.AddNamespace(ns)
}
