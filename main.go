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
	dbPath := flag.String("db", "./callback.db", "Path to the SQLite database file")
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	// Ensure the directory exists
	dbDir := filepath.Dir(*dbPath)
	if dbDir != "." {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			log.Fatalf("Failed to create database directory: %v", err)
		}
	}

	// Initialize database with the specified path
	if err := database.InitDB(*dbPath); err != nil {
		log.Fatal(err)
	}

	// Create router
	r := gin.Default()

	// Serve static files
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	// API routes
	api := r.Group("/api")
	{
		api.POST("/callback", handlers.CallbackHandler)
		api.GET("/callbacks", handlers.GetCallbacksHandler)
	}

	// Start server
	addr := ":" + *port
	log.Printf("Server starting on %s with database at %s", addr, *dbPath)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
