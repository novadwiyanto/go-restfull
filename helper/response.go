package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func SendData(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(400, gin.H{"message": fmt.Sprint(err)})
		return
	}
	c.JSON(200, Response{Status: 200, Data: data, Message: "SUCCESS"})
	return
}

func SendStatus(c *gin.Context, err error) {
	if err != nil {
		c.JSON(400, gin.H{"message": fmt.Sprint(err)})
		return
	}
	c.JSON(200, Response{Status: 200, Message: "SUCCESS"})
	return
}
