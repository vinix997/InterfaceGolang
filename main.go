package main

import (
	"interface/entity"
	"interface/service"
	"time"
)

func main() {
	userSvc := service.NewUserService()

	userSvc.Register(&entity.User{Id: 1, Username: "Delon", Email: "bambang@gmail.com", Password: "bambank", Age: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()})
}
