package main

import (
	// "errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type account struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var accounts = []account{
	{ID: "1", Name: "Tino", Email: "tino@localhost"},
	{ID: "2", Name: "Sue", Email: "sue@localhost"},
}

var home = []account{
	{ID: "55", Name: "Test book", Email: "Test Author"},
}

func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts)
}

func createAccounts(c *gin.Context) {
	var newAccount account

	if err := c.BindJSON(&newAccount); err != nil {
		return
	}
	accounts = append(accounts, newAccount)
	c.IndentedJSON(http.StatusCreated, newAccount)
}

func getHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, home)
}

func main() {
	router := gin.Default()
	router.GET("/accounts", getAccounts)
	router.POST("/accounts", createAccounts)
	router.GET("/", getHome)
	router.Run("localhost:8080")
}
