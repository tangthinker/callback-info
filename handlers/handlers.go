package handlers

import (
	"callback-info/database"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getClientIP(c *gin.Context) string {
	// Try to get IP from X-Forwarded-For header
	xff := c.GetHeader("X-Forwarded-For")
	if xff != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// Try to get IP from X-Real-IP header
	xri := c.GetHeader("X-Real-IP")
	if xri != "" {
		return xri
	}

	// Fallback to client IP
	return c.ClientIP()
}

func CallbackHandler(c *gin.Context) {
	// Get request parameters
	params := make(map[string]interface{})
	if err := c.ShouldBindJSON(&params); err != nil {
		params = make(map[string]interface{})
	}

	// Convert params to JSON string
	paramsJSON, _ := json.Marshal(params)

	// Get client IP
	clientIP := getClientIP(c)

	// Save to database
	err := database.SaveCallback(
		clientIP,
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
