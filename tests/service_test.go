package mainTest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type db struct{}

/*
scenario:
Greet and GreetInDefaultMsg uses greeter
greeter had DB as attribute and DB has Fetchmsg and fetchdefaultmsg functions

what we do:
Mock DB
mock Fetchmsg and fetchdefaultmsg
*/

type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type greeter struct {
	database DB
	lang     string
}

type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}

func (d *db) FetchMessage(lang string) (string, error) {
	if lang == "en" {
		return "Hello", nil
	}
	if lang == "hi" {
		return "Namaste", nil
	}
	return "bzzz", nil
}

func (d *db) FetchDefaultMessage() (string, error) {
	return "default message", nil
}

func (g greeter) Greet() string {
	msg, _ := g.database.FetchMessage(g.lang) // call database to get the message based on the lang
	return "Message is: " + msg
}

func (g greeter) GreetInDefaultMsg() string {
	msg, _ := g.database.FetchDefaultMessage() // call database to get the default message
	return "Message is: " + msg
}

type dbMock struct {
	mock.Mock
}

func (d *dbMock) FetchMessage(lang string) (string, error) {
	args := d.Called(lang)
	return args.String(0), args.Error(1)
}

func (d *dbMock) FetchDefaultMessage() (string, error) {
	args := d.Called()
	return args.String(0), args.Error(1)
}

func TestMockMethodWithoutArgs(t *testing.T) {
	theDBMock := dbMock{}                                            // create the mock
	theDBMock.On("FetchDefaultMessage").Return("foofofofof", nil)    // mock the expectation
	g := greeter{&theDBMock, "en"}                                   // create greeter object using mocked db
	assert.Equal(t, "Message is: foofofofof", g.GreetInDefaultMsg()) // assert what actual value that will come
	theDBMock.AssertNumberOfCalls(t, "FetchDefaultMessage", 1)       // we can assert how many times the mocked method will be called
	theDBMock.AssertExpectations(t)                                  // this method will ensure everything specified with On and Return was in fact called as expected
}
