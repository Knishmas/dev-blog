package main 

import ( 
	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB()
	defer database.DB.Close()
	defer db.Close()

	r := ginDefault() 

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // TODO: Replace later with actual domain
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.run(":8080") 


}