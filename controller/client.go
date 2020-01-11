package controller

import (
	database "../db"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Post struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var Table = "user"

func Get(c *gin.Context) {
	db := database.DBConn()
	getId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	defer db.Close() // close database connect after complete all action
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Cannot parse id",
		})
		return
	}
	if getId <= 0 {
		c.JSON(200, gin.H{
			"messages": "Id not found or not valid",
		})
		return
	}
	rows, err := db.Query("SELECT * FROM user WHERE id = " + c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{
			"messages": "Object not found",
		})
		return
	}
	rs := Post{}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name, password string

		err = rows.Scan(&id, &name, &password)
		if err != nil {
			panic(err.Error())
			return
		}

		rs.Id = id
		rs.Name = name
		rs.Password = password
	}

	c.JSON(200, rs)
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
	query := fmt.Sprintf(`INSERT INTO %v (%v) VALUES (%v)`, Table, queryKey, queryValue)
	_, err = db.Exec(query)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Internal serve 500",
		})
		return
	}
	c.JSON(200, gin.H{
		"messages": "inserted",
	})
}

func Update(c *gin.Context) {
	db := database.DBConn()
	var params *Post
	err := json.NewDecoder(c.Request.Body).Decode(&params)
	defer db.Close() // close database connect after complete all action
	if err != nil {
		panic(err.Error())
		return
	}
	queryKey, queryValue := BuildQueryFromParams(params)
	query := fmt.Sprintf(`UPDATE %v SET %v (%v) VALUES (%v) WHERE id = `, Table, queryKey, queryValue)
	_, err = db.Exec(query)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Internal serve 500",
		})
		return
	}
	c.JSON(200, gin.H{
		"messages": "inserted",
	})
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
		var id int64
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
	if params.Name != "" {
		listKey = append(listKey, "name")
		listValue = append(listValue, "'"+string(params.Name)+"'")
	}
	if params.Password != "" {
		listKey = append(listKey, "password")
		listValue = append(listValue, "'"+string(params.Password)+"'")
	}
	resKey = strings.Join(listKey, ",")
	resValue = strings.Join(listValue, ",")
	return resKey, resValue
}
