package mainTest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type User struct {
	mock.Mock
}

func (m *User) CreateUser(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

func TestSomething(t *testing.T) {

	// create an instance of our test object
	testObj := new(User)

	// setup expectations
	testObj.On("CreateUser", 123).Return(true, nil)

	// call the code we are testing
	//models.InsertOneUser(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)

}

func TestSomething1(t *testing.T) {

	assert.True(t, true, "True is true!")

	var a string = "Hello"
	var b string = "Hello"
	var c string = "Hi"

	assert.Equal(t, a, b, "The two words should be the same.")
	assert.NotEqual(t, c, b, "The two words should not be the same")
}

//func TestGetAllPosts(t *testing.T) {
//	//orm.RegisterDriver("sqlite3", orm.DRSqlite)
//	//orm.RegisterDataBase("default", "sqlite3", "file:data.db")
//	db := new(MockDB)
//
//	db.On("QueryTable").Return(nil)
//	db.On("All").Return(nil)
//
//	response := models.GetAllPosts()
//
//	assert.Equal(t, response, nil, "No error")
//}

//type Storage interface {
//	CreateUser(user User) error
//}
//
//func NewMockStorage() Storage {
//	return &mockStorage{users: make(map[int64]User)}
//}
//
//type mockStorage struct {
//	users map[int64]User
//	lastID int64
//}
//
//func (m *mockStorage) CreateUser(user User) error {
//	m.users[lastID] = user
//	return nil
//}
