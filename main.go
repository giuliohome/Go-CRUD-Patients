// File: main.go
// Version: 2021-04-19
// Description: The purpose of this file is to develop a CRUD web application
// where users can Create, Update, and Delete patients. 

package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	LastName string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Age string `json:"age"`
}

var Users = []User{
	User{uuid.NewV1().String(), "Edgar", "Elias","+1 226 972-3049", "24"},
	User{uuid.NewV1().String(), "Adrian", "Johnson","+1 226 972-3049", "21"},
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", GetUsers)

	userRooutes := r.Group("/users")
	{
		userRooutes.GET("/", GetUsers)
		userRooutes.POST("/", CreateUser)
		userRooutes.PUT("/:id", EditUser)
		userRooutes.DELETE("/:id", DeleteUser)
	}

	if err := r.Run(":3000"); err != nil{
		log.Fatal(err.Error())
	}
}

func GetUsers(c *gin.Context){
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":"Patients",
		"Users":Users,
	})
}

func CreateUser(c *gin.Context){
	var requestBody User

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(422, gin.H{
			"error":true,
			"message": "Invalid Request Body",
		})
		return
	}

	if requestBody.Name != "" && requestBody.LastName != "" && requestBody.PhoneNumber != "" && requestBody.Age != "" {		
		requestBody.ID = uuid.NewV1().String()
		Users = append(Users, requestBody)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":"Patients",
			"Users":Users,
		})
	} else {
		c.JSON(422, gin.H{
			"error":true,
			"message": "Invalid Request Body",
		})
		return
	}

}

func EditUser(c *gin.Context){
	
	id := c.Param("id")

	var requestBody User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(422, gin.H{
			"error":true,
			"message": "Invalid Request Body",
		})
		return
	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = requestBody.Name
			Users[i].LastName = requestBody.LastName
			Users[i].PhoneNumber = requestBody.PhoneNumber
			Users[i].Age = requestBody.Age

			c.JSON(200, gin.H{
				"error": false,
				"message": "User Updated.",
				"user": Users[i],
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":true,
		"message": "Invalid User ID",
	})
}


func DeleteUser(c *gin.Context){
	
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i + 1:]...)

			c.JSON(200, gin.H{
				"error": false,
				"message": "User Deleted.",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":true,
		"message": "Invalid User ID",
	})
}