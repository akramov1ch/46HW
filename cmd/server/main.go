package main

import (
	"46HW/pb"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/hash", func(c *gin.Context) {
		var req pb.Request
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hash := sha256.Sum256([]byte(req.Message))
		resp := pb.Response{
			Message: hex.EncodeToString(hash[:]),
		}
		c.JSON(http.StatusOK, resp)
	})
	router.Run(":8080")
}
