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
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func init() {
	//var Auth = func(w http.ResponseWriter, r *http.Request) (err error){
	//	_, err = ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		log.Printf("Error reading body: %v", err)
	//		http.Error(w, "can't read body", http.StatusBadRequest)
	//		return err
	//	}
	//	return nil
	//}
	//var Validate = func(token jwt.Token, tknstr string) (err error){
	//	claims :=&Claims{}
	//}
	var Authorize = func(ctx *context.Context) {
		//if strings.HasPrefix(ctx.Input.URL(), "/login/" ) {
		//	return
		//}
		if strings.HasPrefix(ctx.Input.URL(), "/login/") ||
			(ctx.Request.RequestURI == "/") || (ctx.Request.RequestURI == "/authorize/") {
			// just do nothing in the filter and complete the logic in controller
			fmt.Println("continued", time.Now())
			return
		}
		tknStr := ctx.Input.Header("Authorization")
		fmt.Println(tknStr, 2)
		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		fmt.Println(tkn, 2)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				MyErr := errors.New("401: Unauthorized")
				fmt.Println(MyErr)
				ctx.Output.SetStatus(401)
				ctx.Abort(401, "Unauthorized")
			}
			MyErr := errors.New("400: BadRequest")
			fmt.Println(MyErr)
			ctx.Abort(400, "Bad request")
		}
		if !tkn.Valid {
			MyErr := errors.New("401: Unauthorized")
			fmt.Println(MyErr)
			ctx.Abort(401, "Unauthorized")
		}
	}

	//beego.InsertFilter("/*", beego.BeforeRouter, Auth)
	beego.InsertFilter("/*", beego.BeforeRouter, Authorize)
	// hello-world route
	beego.Router("/", &controllers.HelloController{}, "get:HelloWorld")
	// Login
	beego.Router("/login/", &controllers.LoginController{}, "post:Login")
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
		beego.NSNamespace("/reply",
			// create Reply
			beego.NSRouter("/", &controllers.ReplyController{}, "post:CreateReply"),
			//Get All replies of a comment
			beego.NSRouter("/:CommentId", &controllers.ReplyController{}, "get:GetRepliesOfComment"),
		),
	)
	// register namespace
	beego.AddNamespace(ns)
}
