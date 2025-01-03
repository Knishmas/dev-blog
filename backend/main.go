package main 

import ( 
	"net/http"
	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
	"backend/database"
	"log"
)

func main() {
	// Initialize Database
	database.InitDB()
	defer database.CloseDB() // Close DB properly

	// Initialize Router
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // TODO: Replace with actual domain in production
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	// Test database route
	r.GET("/test-db", func(c *gin.Context) {
		rows, err := database.DB.Query("SELECT name FROM sqlite_master WHERE type='table';")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			return
		}
		defer rows.Close()
	
		var tables []string
		for rows.Next() {
			var table string
			if err := rows.Scan(&table); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read row"})
				return
			}
			tables = append(tables, table)
		}
	
		c.JSON(http.StatusOK, gin.H{"tables": tables})
	})
	

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}