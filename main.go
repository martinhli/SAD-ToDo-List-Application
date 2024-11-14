package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"to-do-app/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ToDo struct {
	ID          string    `json: "id"`
	Title       string    `json: "title"`
	Description string    `json: "description"`
	Completed   bool      `json: "completed"`
	CreatedAt   time.Time `json: "createdat"`
}

/*				REST API				*/

var items = []ToDo{}

func GetItems(c *gin.Context) {
	if len(items) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Hi! Welcome to the To-Do-App! Post your first item to get started."})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "All items are retrieved",
			"items":   items,
		})
	}
}

func GetItembyID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Item retrieved",
				"id":      id,
				"item":    item,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func CreateItem(c *gin.Context) {
	var item ToDo

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = fmt.Sprintf("%d", len(items)+1)
	item.CreatedAt = time.Now()
	items = append(items, item)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Item was created",
		"item":    item,
	})
}

func UpdateItembyID(c *gin.Context) {
	id := c.Param("id")
	var updateditem ToDo

	if err := c.ShouldBindJSON(&updateditem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i].ID = updateditem.ID
			items[i].Title = updateditem.Title
			items[i].Description = updateditem.Description
			items[i].Completed = updateditem.Completed
			c.JSON(http.StatusOK, items)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

/*				LOGGING				*/

// Middleware for handling user interactions
func LogUserInteractions() gin.HandlerFunc {
	return func(c *gin.Context) { // Returning a function inside the function is useful in this case since we are making a HTTP handler.
		//We need to provide additional data (like the database db, api etc.) without changing the signature of the function
		logger.Log.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"ip":     c.ClientIP(),
		}).Info("User interaction logged") //Updating with info that the user interaction has been logged
		c.Next()
	}
}

func main() {
	// Access the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Print for debugging purposes
	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_USER:", dbUser)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DB_NAME:", dbName)

	logger.Initialize() // Initialize the logger settings

	router := gin.Default()
	itemsGroup := router.Group("/items")
	itemsGroup.Use(LogUserInteractions())
	{
		itemsGroup.GET("", GetItems)
		itemsGroup.GET("/:id", GetItembyID)
		itemsGroup.POST("", CreateItem)
		itemsGroup.PUT("/:id", UpdateItembyID)
		itemsGroup.DELETE("/:id", DeleteItem)
	}

	// Root route, shows the welcome message if there are no items in the list
	router.GET("/", func(c *gin.Context) {
		if len(items) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the To-Do App! Post your first item to get started."})
		} else {
			c.Redirect(http.StatusMovedPermanently, "/items")
		}
	})

	router.Run(":8080")
}
