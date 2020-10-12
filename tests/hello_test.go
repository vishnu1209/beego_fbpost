package mainTest

import (
	_ "awesomeProject/common"
	"awesomeProject/models"
	_ "github.com/mattn/go-sqlite3"
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

//func TestGetAllUsers(t *testing.T) {
//	assert := assert.New(t)
//	t.Run("should return data", func(t *testing.T) {
//		testDb := new(mocks.OrmDB)
//		testDb.On("GetAllUsers").Return(nil)
//		assert.Equal(t, "Message is: foofofofof", models.GetAllUsers())
//		testDb.AssertNumberOfCalls(t, "GetAllUsers", 1)
//		testDb.AssertExpectations(t)
//	})
//}

/*
type MockDB struct {
	mock.Mock
}

func (db *MockDB) QueryTable(s string) orm.QuerySeter {
	args := db.Called()
	fmt.Println(args)
	return nil
}

func (db *MockDB) All(i interface{}, cols string) (int, error) {
	args := db.Called()
	fmt.Println(args)
	return args.Int(0), nil
}

func (db *MockDB) One(i interface{}, cols string) (int, error) {
	args := db.Called()
	fmt.Println(args)
	return args.Int(0), nil
}
func (db *MockDB) Filter(i interface{}, cols string) (int, error) {
	args := db.Called()
	fmt.Println(args)
	return args.Int(0), nil
}
func (db *MockDB) Insert(i interface{}, cols string) (int, error) {
	args := db.Called()
	fmt.Println(args)
	return args.Int(0), nil
}

//func TestGetAllUsers(t *testing.T) {
//	orm.RegisterDriver("sqlite3", orm.DRSqlite)
//	orm.RegisterDataBase("default", "sqlite3", "file:data.db")
//	o := MockDB{}
//	o.On("QueryTable").Return(nil)
//	o.On("All").Return(nil)
//	var users []*models.User
//	users, err := models.GetAllUsers(&o)
//}

//func TestUser(t *testing.T) {
//
//	orm.RegisterDriver("sqlite3", orm.DRSqlite)
//	orm.RegisterDataBase("default", "sqlite3", "file:data.db")
//
//	user := models.User{
//		Id:        9,
//		FirstName: "vishnu",
//		LastName:  "k v",
//	}
//	o := models.MockDB{}
//	o.On("Insert").Return(nil)
//
//	Id, ErrorInInsert := models.CreateUser(&user)
//	if ErrorInInsert != nil {
//		fmt.Println(ErrorInInsert)
//		t.Errorf("User was unable to be created %v", ErrorInInsert)
//	}
//	o.On("QueryTable").Return(nil)
//	o.On("One").Return(nil)
//	o.On("Filter").Return(nil)
//
//	var u models.User
//	InvalidPostId := o.QueryTable("User").Filter("Id", Id).One(&u)
//	if InvalidPostId != nil {
//		fmt.Println(InvalidPostId)
//		t.Errorf("User Not created err %v", InvalidPostId)
//	}
//	if Id == 9 && u.Id == 9 && u.FirstName == user.FirstName && u.LastName == user.LastName {
//		t.Logf("User is created succesfully")
//	}
//}
*/
