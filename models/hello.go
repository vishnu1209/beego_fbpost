package models

import "fmt"

func Hello(user string) string {

	if len(user) == 0 {
		welcomeMessage := "Welcome user.!"
		fmt.Println(welcomeMessage)
		return welcomeMessage
	}
	welcomeMessage := "Welcome " + user
	return welcomeMessage
}
