package test

import (
	_ "awesomeProject/common"

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
