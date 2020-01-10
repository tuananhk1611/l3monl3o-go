package controller

import (
	database "../db"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Post struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Get(c *gin.Context) {

	db := database.DBConn()
	rows, err := db.Query("SELECT id, name, password FROM user WHERE id = " + c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Object not found",
		})
	}

	post := Post{}

	for rows.Next() {
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

func Create(c *gin.Context) {
	db := database.DBConn()
	var params *Post
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	defer db.Close() // close database connect after complete all action
	if err != nil {
		panic(err.Error())
		return
	}
	queryKey, queryValue := BuildQueryFromParams(params)
	fmt.Println("test %v %v", queryKey, queryValue)
	res, err := db.Prepare("INSERT INTO user(" + queryKey + ") VALUES(?,?)")
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Internal serve 500",
		})
	}
	_, err = res.Exec(params.Name, params.Password)
	if err != nil {
		panic(err.Error())
		return
	}
	c.JSON(200, res)
}

func Update(c *gin.Context) {
	db := database.DBConn()
	rows, err := db.Query("SELECT id, name FROM user WHERE id = " + c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Object not found",
		})
	}

	post := Post{}

	for rows.Next() {
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

func Delete(c *gin.Context) {
	db := database.DBConn()
	rows, err := db.Query("SELECT id, name FROM user WHERE id = " + c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Object not found",
		})
	}

	post := Post{}

	for rows.Next() {
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

func BuildQueryFromParams(params *Post) (resKey string, resValue string) {
	listKey := make([]string, 0)
	listValue := make([]string, 0)
	if params.Id != 0 {
		listKey = append(listKey, "id")
		listValue = append(listValue, string(params.Id))
	}
	if params.Name != "" {
		listKey = append(listKey, "name")
		listValue = append(listValue, string(params.Name))
	}
	if params.Password != "" {
		listKey = append(listKey, "password")
		listValue = append(listValue, string(params.Password))
	}
	resKey = strings.Join(listKey, ",")
	resValue = strings.Join(listValue, ",")
	return resKey, resValue
}
