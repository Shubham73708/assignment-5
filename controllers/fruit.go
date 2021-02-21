package controllers

import (
	"bytes"
	"fmt"
	"fruit-api/db"
	"fruit-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func fruit(c *gin.Context) {
// 	tmpl := template.Must(template.ParseFiles("fruit.html"))
// 	tmpl.Execute(nil, c)
// }

//Create new fruit details
func Create(c *gin.Context) {
	// config.TPL.ExecuteTemplate( c,"create.gohtml", nil)
	var buffer bytes.Buffer
	// id := c.PostForm("id")
	name := c.PostForm("name")
	colour := c.PostForm("colour")
	price := c.PostForm("price")
	quantity := c.PostForm("quantity")
	descriptions := c.PostForm("descriptions")

	stmt, err := db.Init().Prepare("insert into fruit(name,colour,price,quantity,descriptions)values(?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(name, colour, price, quantity, descriptions)
	// Fastest way to append strings
	buffer.WriteString(name)
	buffer.WriteString(" ")
	defer stmt.Close()
	fruitname := buffer.String()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully created:", fruitname),
	})

}

// func signup(c *gin.Context) {
// 	tmpl := template.Must(template.ParseFiles("signup.html"))
// 	tmpl.Execute(w, nil)
// }

func Signup(c *gin.Context) {
	var buffer bytes.Buffer
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	stmt, err := db.Init().Prepare("insert into users(username,email,password)values(?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(username, email, password)
	// Fastest way to append strings
	buffer.WriteString(username)
	buffer.WriteString(" ")
	defer stmt.Close()
	fruitname := buffer.String()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully created:", fruitname),
	})

}

// //Create new fruit details
func Signin(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	passwordFromDB := ""
	err := db.Init().QueryRow("select password FROM users WHERE email = ?", email).Scan(&passwordFromDB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf(" Unsuccessfull"),
		})
	} else {
		if password == passwordFromDB {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf(" successfully logged in."),
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf(" Unsuccessfull!, data does not match"),
			})
		}
	}

}

// GET all fruits
func GetAllFruit(c *gin.Context) {
	var (
		fruit  models.Fruit
		fruits []models.Fruit
	)
	rows, err := db.Init().Query("select * from fruit;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&fruit.Id, &fruit.Name, &fruit.Colour, &fruit.Price, &fruit.Quantity, &fruit.Descriptions)
		fruits = append(fruits, fruit)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": fruits,
	})
}

// GET a fruit detail
func GetFruit(c *gin.Context) {
	var (
		fruit  models.Fruit
		result gin.H
	)
	id := c.Param("id")
	err := db.Init().QueryRow("select * from fruit where id = ?;", id).Scan(&fruit.Id, &fruit.Name, &fruit.Colour, &fruit.Price, &fruit.Quantity, &fruit.Descriptions)

	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": fruit,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Delete fruit
func DeleteFruit(c *gin.Context) {
	id := c.Param("id")
	// fmt.Println(id)
	stmt, err := db.Init().Prepare("delete from fruit where id= ?;")

	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted fruit: %s", id),
	})
}
