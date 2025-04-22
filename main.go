package main

import (
	"callback-info/database"
	"callback-info/handlers"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Parse command line flags
	rootDir := flag.String("root", ".", "Root directory for the application")
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	// Ensure the root directory exists
	if err := os.MkdirAll(*rootDir, 0755); err != nil {
		log.Fatalf("Failed to create root directory: %v", err)
	}

	// Initialize database with the specified path
	dbPath := filepath.Join(*rootDir, "callback.db")
	if err := database.InitDB(dbPath); err != nil {
		log.Fatal(err)
	}

	// Create router
	r := gin.Default()

	staticFilesPath := filepath.Join(*rootDir, "/static")
	// Serve static files
	r.Use(static.Serve("/", static.LocalFile(staticFilesPath, false)))

	// API routes
	api := r.Group("/api")
	{
		api.POST("/callback", handlers.CallbackHandler)
		api.GET("/callbacks", handlers.GetCallbacksHandler)
	}

	// Start server
	addr := ":" + *port
	log.Printf("Server starting on %s with root directory at %s", addr, *rootDir)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
