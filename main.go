package main

import (
	"github.com/0xjeffro/tx-parser/solana"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/solana", solanaHandler)
	err := r.Run()
	if err != nil {
		log.Println(err)
		return
	}
}

func solanaHandler(c *gin.Context) {
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Fail(c, err)
		return
	}

	res, err := solana.Parser(bytes)
	if err != nil {
		Fail(c, err)
		return
	}
	Success(c, res)
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
	})
}

func Fail(c *gin.Context, err error) {
	c.JSON(400, gin.H{
		"message": "error",
		"error":   err.Error(),
	})
}
