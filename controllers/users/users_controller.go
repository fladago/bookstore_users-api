package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//function to handle get user request
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
