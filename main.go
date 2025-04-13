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

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// todo, _ := models.GetTodo(1)
	// fmt.Println(todo)

	// user3, _ := models.GetUser(3)
	// user3.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, todo := range todos {
	// 	fmt.Println(todo)
	// }

	// user2, _ := models.GetUser(3)
	// todos2, _ := user2.GetTodosByUser()
	// for _, todo := range todos2 {
	// 	fmt.Println(todo)
	// }

	todo, _ := models.GetTodo(1)
	fmt.Println(todo)
	todo.Content = "Updated Todo"
	todo.Update()
	todo, _ = models.GetTodo(1)
	fmt.Println(todo)
}