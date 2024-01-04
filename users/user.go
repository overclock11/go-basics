package users

import (
	"fmt"
	"goltest/models"
	"time"
)

func CreateUser() {
	var newUser = new(models.User)
	newUser.AddUser(1, "Juan", time.Now(), true)
	fmt.Println(newUser)
}
