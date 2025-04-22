package handlers

import (
	"callback-info/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CallbackHandler(c *gin.Context) {
	// Get request parameters
	params := make(map[string]interface{})
	if err := c.ShouldBindJSON(&params); err != nil {
		params = make(map[string]interface{})
	}

	// Convert params to JSON string
	paramsJSON, _ := json.Marshal(params)

	// Save to database
	err := database.SaveCallback(
		c.ClientIP(),
		c.Request.Host,
		string(paramsJSON),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetCallbacksHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	response, err := database.GetCallbacks(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
