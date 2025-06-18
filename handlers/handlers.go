package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santori/ping-checker/checker"
)

type CheckRequest struct {
	URLs []string `json:"urls"`
}

func CheckReq(c *gin.Context) {
	var request CheckRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	response := checker.CheckURLs(request.URLs)

	c.JSON(http.StatusOK, response)
}
