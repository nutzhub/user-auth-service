package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)


const (
	CREATE_USER_TABLE = `CREATE TABLE IF NOT EXISTS user (
		id  INT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NUll,
		userName TEXT UNIQUE,
		password TEXT NOT NULL
	)`// multiline string
)

type User struct {
	id       int
	name     string
	email    string
	userName string
	password string
}


func createDB() *sql.DB {
	os.Remove("./temp.db")
	db, err := sql.Open("sqlite3", "./temp.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	r := gin.Default()
	db := createDB()
	defer db.Close()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is runing")
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})
	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})

	r.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})

	r.POST("/users/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "waiting for implementation", "status": http.StatusOK})
	})

	log.Fatal(r.Run(":8080"))
}
