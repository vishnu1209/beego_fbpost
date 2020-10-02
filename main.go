package main

import (
	"awesomeProject/models"
	_ "awesomeProject/models"
	_ "awesomeProject/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:data.db")
}

var (
	router = gin.Default()
)

func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func main() {
	//router.POST("/login", Login)
	//log.Fatal(router.Run(":8080"))
	//authPlugin := auth.NewBasicAuthenticator(SecretAuth, "My Realm")
	//beego.InsertFilter("*", beego.BeforeRouter, authPlugin)

	// if in develop mode
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.WebConfig.Session.SessionOn = true

	orm.Debug = true

	// autosync
	// db alias
	name := "default"

	// drop table and re-create
	force := false

	// print log
	verbose := true

	// error

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	GreetingMessage := models.Hello("vishnu")
	fmt.Println(GreetingMessage)

	beego.Run()
}

func SecretAuth(username, password string) bool {
	// The username and password parameters comes from the request header,
	// make a database lookup to make sure the username/password pair exist
	// and return true if they do, false if they dont.

	// To keep this example simple, lets just hardcode "hello" and "world" as username,password
	if username == "hello" && password == "world" {
		return true
	}
	return false
}
