package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Item represents a product with ID, Name, and Price
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Sample data
var items = []Item{
	{ID: 1, Name: "Item1", Price: 10.5},
	{ID: 2, Name: "Item2", Price: 20.0},
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Logging middleware
	router.Use(loggingMiddleware)

	// Define routes
	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.GET("/items/:id", getItemByID)
	router.PUT("/items/:id", updateItemByID)
	router.DELETE("/items/:id", deleteItemByID)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}

// loggingMiddleware logs incoming requests and their processing time
func loggingMiddleware(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	method := c.Request.Method

	c.Next()

	statusCode := c.Writer.Status()
	log.Printf("[%s] %s %d %v", method, path, statusCode, time.Since(start))
}

// getItems handles GET /items and returns the list of items
func getItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, items)
	log.Println("Returned all items")
}

// addItem handles POST /items and adds a new item
func addItem(ctx *gin.Context) {
	var newItem Item
	if err := ctx.ShouldBindJSON(&newItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Invalid input: %v", err)
		return
	}

	newItem.ID = len(items) + 1
	items = append(items, newItem)

	ctx.JSON(http.StatusCreated, newItem)
	log.Printf("Added new item: %+v", newItem)
}

// getItemByID handles GET /items/:id and returns a single item by ID
func getItemByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		log.Printf("Invalid ID: %s", ctx.Param("id"))
		return
	}

	for _, item := range items {
		if item.ID == id {
			ctx.JSON(http.StatusOK, item)
			log.Printf("Returned item with ID %d", id)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	log.Printf("Item not found: ID %d", id)
}

// updateItemByID handles PUT /items/:id to update an existing item by ID
func updateItemByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		log.Printf("Invalid ID: %s", ctx.Param("id"))
		return
	}

	var updatedItem Item
	if err := ctx.ShouldBindJSON(&updatedItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Invalid input: %v", err)
		return
	}

	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = item.ID // Preserve the original ID
			items[i] = updatedItem
			ctx.JSON(http.StatusOK, updatedItem)
			log.Printf("Updated item with ID %d: %+v", id, updatedItem)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	log.Printf("Item not found: ID %d", id)
}

// deleteItemByID handles DELETE /items/:id to delete an item by ID
func deleteItemByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		log.Printf("Invalid ID: %s", ctx.Param("id"))
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			log.Printf("Deleted item with ID %d", id)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	log.Printf("Item not found: ID %d", id)
}
