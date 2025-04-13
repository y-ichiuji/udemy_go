package main

import (
	"fmt"

	"github.com/y-ichiuji/udemy-go/app/controllers"
	"github.com/y-ichiuji/udemy-go/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}