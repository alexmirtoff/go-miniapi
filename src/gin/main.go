package main

import (
	"database/sql"
	"fmt"
	//"reflect" 
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type users struct {
	//id   int    `json:"id"`
	//name string `json:"name"`
	id   int
	name string
	nick string
}

type posts struct {
	id    int
	title string
	body  string
}

func main() {
	// init router
	r := gin.Default()

	// init root

	r.GET("/userlist", func(c *gin.Context) {
		userList := getUserList()
		c.JSON(200, userList)
	})

	// create user
	r.GET("/createuser", func(c *gin.Context) {
		nameOnCreate := c.Query("name")
		nickOnCreate := c.Query("nick")
		cuMap := make(map[string]string)
		cuMap["name"] = nameOnCreate
		cuMap["nick"] = nickOnCreate
		newUser(cuMap)
	})

	// drop user
	r.GET("/dropuser", func(c *gin.Context) {
		dropIds := c.Query("id")
		dropIdi, _ := strconv.Atoi(dropIds)
		dropUser(dropIdi)
		
	})

	// run server
	r.Run()
}

func initDb() (db *sql.DB) {
	// init db
	db, err := sql.Open("sqlite3", "../../db/test.db")
	checkErr(err)
	return db
}

// SHOW ALL USERS
func getUserList() (out map[int]string) {
	db := initDb()
	rows, err := db.Query("SELECT * FROM USERS")
	checkErr(err)
	// int user and map
	userCursor := new(users)
	userMap := make(map[int]string)

	for rows.Next() {
		err = rows.Scan(&userCursor.id, &userCursor.name, &userCursor.nick)
		checkErr(err)
		userMap[userCursor.id] = userCursor.name

	}
	rows.Close()
	db.Close()
	return userMap
}

// SHOW USER BY NICK
// func getUserByNick(db *sql.DB) (out map[int]string) {

// }

// INSERT NEW USER
func newUser(newUser map[string]string) {
	db := initDb()
	ex, err := db.Prepare("INSERT INTO users(name, nick) values(?, ?)")
	checkErr(err)
	res, err := ex.Exec(newUser["name"], newUser["nick"])
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	db.Close()
}

func dropUser(id int) {
	db := initDb()
	ex, err := db.Prepare("DELETE FROM users where id=?")
	checkErr(err)
	res, err := ex.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("User dropped: ", affect)
	db.Close()
}

// errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//gin.H{"kek": "lel"}
