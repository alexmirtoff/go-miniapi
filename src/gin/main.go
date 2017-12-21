package main

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type users struct {
	id int	`json:"id"`
	name string `json:"name"`
}

type posts struct {
	id int
	title string
	body string
}

func main() {
	db, err := sql.Open("sqlite3", "../../db/test.db")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		log.Fatal(err)
	}
	newUser := new(users)
	aaa := make(map[int]string)
	for rows.Next() {
		err = rows.Scan(&newUser.id, &newUser.name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(newUser.id, newUser.name)
		aaa[newUser.id] = newUser.name
		
	}
	
	

	fmt.Println(aaa)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, aaa)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

//gin.H{"kek": "lel"}