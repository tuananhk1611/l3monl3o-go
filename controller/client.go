package controller

import (
	database "../db"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func Read(c * gin.Context){

	db := database.DBConn()
	rows, err := db.Query("SELECT id, name FROM user WHERE id = " + c.Param("id"))
	if err != nil{
		c.JSON(500, gin.H{
			"messages" : "Object not found",
		});
	}

	post := Post{}

	for rows.Next(){
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Name = name
	}

	c.JSON(200, post)
	defer db.Close() // close database connect after complete all action
}
