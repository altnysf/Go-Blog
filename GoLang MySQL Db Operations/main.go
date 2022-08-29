package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func main() {
	dsn := "superadmin:PfjHZ!YPryjMIrPR@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Migrate a schema

	// db.AutoMigrate(&User{}) Authomatically create a table with 'User' struct informations

	// ----------------------------------------------------------------------------------------------------

	// Create a data

	// db.Create(&User{Username: "admin", Password: "sifre123"})

	// ----------------------------------------------------------------------------------------------------

	// Read Single Data - First Way

	var userRead User
	db.First(&userRead, 1)
	fmt.Println(userRead)
	fmt.Println(userRead.Username)
	fmt.Println(userRead.Password)

	/* OUTPUT
	{{1 2022-08-29 13:18:53.968 +0300 +03 2022-08-29 13:18:53.968 +0300 +03 {0001-01-01 00:00:00 +0000 UTC false}} admin sifre123}
	admin
	sifre123
	*/

	// Read Single Data - Second Way

	db.First(&userRead, "Username = ?", "admin")
	fmt.Println(userRead.Username, userRead.Password, userRead.CreatedAt)

	/* OUTPUT
	admin sifre123 2022-08-29 13:18:53.968 +0300 +03
	*/

	// Read Multiple Data
	var usersRead []User
	db.Find(&usersRead, "Password = ?", "sifre123")
	fmt.Println(usersRead)
	/* OUTPUTS
	[{{1 2022-08-29 13:18:53.968 +0300 +03 2022-08-29 13:18:53.968 +0300 +03 {0001-01-01 00:00:00 +0000 UTC false}} admin sifre123} {{2 2022-08-29 13:39:11.097 +0300 +03 2022-08-29 13:39:11.097 +0300 +03 {0001-01-01 00:00:00 +0000 UTC false}} goLang sifre123}]
	*/

	// Update Single Data

	var userUpt User
	db.First(&userUpt, 2)
	db.Model(&userUpt).Update("username", "gorm")

	// Update Multiple Data

	var usersUpt User
	db.First(&usersUpt, 2)
	db.Model(&userUpt).Updates(User{Username: "goLang", Password: "go1234"})

	// Delete Data

	db.Delete(&User{}, 2)
}
