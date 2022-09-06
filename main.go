package main

import (
	"log"

	"how-to-init-sql.DB/dao"
	"how-to-init-sql.DB/dao/model"
)

func main() {
	err := dao.User.Create(&model.User{Username: "spike014", Password: "你猜猜", Email: "l1ng14@foxmail.com"})
	if err != nil {
		log.Println("Insert user fail:", err)
	}
	log.Println("User created!")
}
