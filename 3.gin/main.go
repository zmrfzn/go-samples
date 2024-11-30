package main

import (
	"net/http"
	"strconv"

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

	// Define routes
	router.GET("/items", getItems)              // Get all items
	router.GET("/items/:id", getItemByID)       // Get a single item by ID
	router.POST("/items", addItem)              // Add a new item
	router.PUT("/items/:id", updateItemByID)    // Update an item by ID
	router.DELETE("/items/:id", deleteItemByID) // Delete an item by ID

	// Start the server
	router.Run(":8080") // By default, listens on http://localhost:8080
}

// getItems handles GET /items and returns the list of items
func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

// getItemByID handles GET /items/:id and returns a single item by ID
func getItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// addItem handles POST /items and adds a new item to the list
func addItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new item to the slice
	newItem.ID = len(items) + 1 // Assign a new ID
	items = append(items, newItem)

	c.JSON(http.StatusCreated, newItem)
}

// updateItemByID handles PUT /items/:id to update an existing item by ID
func updateItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = item.ID // Ensure ID remains unchanged
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// deleteItemByID handles DELETE /items/:id to delete an item by ID
func deleteItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...) // Remove the item
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
