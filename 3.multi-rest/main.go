package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
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
	http.HandleFunc("/items", logMiddleware(handleItems))
	http.HandleFunc("/items/", logMiddleware(handleItemByID))

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// logMiddleware adds logging for each request
func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}
}

// handleItems handles GET and POST requests for /items
func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		addItem(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("Method not allowed: %s", r.Method)
	}
}

// handleItemByID handles GET, PUT, and DELETE requests for /items/:id
func handleItemByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/items/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid ID: %s", idStr)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getItemByID(w, r, id)
	case http.MethodPut:
		updateItemByID(w, r, id)
	case http.MethodDelete:
		deleteItemByID(w, r, id)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("Method not allowed: %s", r.Method)
	}
}

// getItems returns all items
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		log.Printf("Error encoding items: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Println("Returned all items")
}

// addItem adds a new item
func addItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Invalid input: %v", err)
		return
	}

	newItem.ID = len(items) + 1
	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newItem); err != nil {
		log.Printf("Error encoding new item: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Printf("Added new item: %+v", newItem)
}

// getItemByID returns a single item by ID
func getItemByID(w http.ResponseWriter, r *http.Request, id int) {
	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(item); err != nil {
				log.Printf("Error encoding item: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			log.Printf("Returned item with ID %d", id)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
	log.Printf("Item not found: ID %d", id)
}

// updateItemByID updates an existing item by ID
func updateItemByID(w http.ResponseWriter, r *http.Request, id int) {
	var updatedItem Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		log.Printf("Invalid input: %v", err)
		return
	}

	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = item.ID
			items[i] = updatedItem
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(updatedItem); err != nil {
				log.Printf("Error encoding updated item: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			log.Printf("Updated item with ID %d: %+v", id, updatedItem)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
	log.Printf("Item not found: ID %d", id)
}

// deleteItemByID deletes an item by ID
func deleteItemByID(w http.ResponseWriter, r *http.Request, id int) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(map[string]string{"message": "Item deleted"}); err != nil {
				log.Printf("Error encoding delete response: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			log.Printf("Deleted item with ID %d", id)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
	log.Printf("Item not found: ID %d", id)
}
