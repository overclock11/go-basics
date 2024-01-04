package models

import "time"

// es estandar definir todos los modelos en el paquete de models
type User struct {
	id        int
	name      string
	createdAt time.Time
	status    bool
}

func (user *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	user.id = id
	user.name = name
	user.createdAt = createdAt
	user.status = status
}
