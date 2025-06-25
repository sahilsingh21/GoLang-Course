package main

//go mod tidy //automatic fix (command)
// go get "package"// install

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sahilsingh21/golangCourse/auth"
	"github.com/sahilsingh21/golangCourse/user"
)

func main() {
	auth.LoginWithCredentials("Bob", "secret")

	session := auth.GetSession()
	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "Jogn Doe",
	}

	// fmt.Println(user.Email, user.Name)

	//3rd party packages
	color.Green(user.Email)

}
