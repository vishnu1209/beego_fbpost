package test

import (
	_ "awesomeProject/common"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

	"awesomeProject/models"
	"testing"
)

func TestHello(t *testing.T) {
	emptyResult := models.Hello("")
	if emptyResult != "Welcome user.!" {
		t.Errorf("hello(\"\") failed expected %v but got %v", "Welcome user.!", emptyResult)
	} else {
		t.Logf("hello(\"\") success expected %v and got %v", "Welcome user.!", emptyResult)
	}
	result := models.Hello("qwerty")
	if result != "Welcome qwerty" {
		t.Errorf("hello(\"qwerty\") failed! Expected %v but got %v", "Welcome qwerty", result)
	} else {
		t.Logf("hello(\"qwerty\") Success! Expected %v and got %v", "Welcome qwerty", result)
	}
}

func TestUser(t *testing.T) {

	user := models.User{
		Id:        9,
		FirstName: "vishnu",
		LastName:  "k v",
	}
	Id, ErrorInInsert := models.InsertOneUser(&user)
	if ErrorInInsert != nil {
		fmt.Println(ErrorInInsert)
		t.Errorf("User was unable to be created %v", ErrorInInsert)
	}
	o := orm.NewOrm()
	var u models.User
	InvalidPostId := o.QueryTable("User").Filter("Id", Id).One(&u)
	if InvalidPostId != nil {
		fmt.Println(InvalidPostId)
		t.Errorf("User Not created err %v", InvalidPostId)
	}
	if Id == 9 && u.Id == 9 && u.FirstName == user.FirstName && u.LastName == user.LastName {
		t.Logf("User is created succesfully")
	}
}
