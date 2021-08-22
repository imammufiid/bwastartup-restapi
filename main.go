package main

import (
	"bwastartup/user"
	_ "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connection to database
	// dsn := "root:@tcp(127.0.0.1:3306)/learn_bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connection to database is GOOD!!!")

	// var users []user.User

	// // get data users from db
	// db.Find(&users)
	
	// // show user
	// for i, user := range users {
	// 	fmt.Println("ID : ",  user.ID)
	// 	fmt.Println("Name : ",  user.Name)
	// 	fmt.Println("Email : ",  user.Email)
	// 	fmt.Println("Role : ",  user.Role)
		
	// 	if i != len(users)-1 {
	// 		fmt.Println()
	// 	}
	// }

}

func handler(c *gin.Context)  {
	
	// connection to database
	dsn := "root:@tcp(127.0.0.1:3306)/learn_bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User

	// get data users from db
	db.Find(&users)

	// show users to JSON in web API
	c.JSON(http.StatusOK, users)

}
