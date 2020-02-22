package main

import (
	"fmt"

	"github.com/akshitGoel09/goSimpleUserService/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
