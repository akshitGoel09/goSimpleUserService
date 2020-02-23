package main

import (
	"fmt"

	"github.com/akshitGoel09/goSimpleUserService/model"
)

func main() {
	u := model.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
