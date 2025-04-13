package main

import (
	"fmt"

	"github.com/y-ichiuji/udemy-go/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.Create()

	// user, _ := models.GetUser(1)
	// fmt.Println(user)

	// user.Name = "test2"
	// user.Email = "test2@example.com"
	// user.Update()

	// user, _ = models.GetUser(1)
	// fmt.Println(user)

	// user.Delete()
	// user, _ = models.GetUser(1)
	// fmt.Println(user)

	user, _ := models.GetUser(2)
	user.CreateTodo("First Todo")
}