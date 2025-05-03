// Go HTMX CRUD Application
//
// This application demonstrates a modern approach to building interactive web applications
// using Go for the backend and HTMX for frontend interactivity. It implements a basic
// CRUD (Create, Read, Update, Delete) system for managing items with image uploads.
//
// Key Features:
// - Server-side rendering with Go templates
// - HTMX for dynamic content updates without full page reloads
// - File upload handling
// - RESTful API endpoints
// - Chi router for HTTP routing
// - In-memory data storage (for demonstration purposes)

package main

import (
	"fmt"
	"html/template" // For server-side HTML templating
	"io"            // For file copying operations
	"log"
	"net/http"
	"os"
	"path/filepath"

	//  go mod download github.com/go-chi/chi/v5
	//  go mod download github.com/go-chi/cors

	"github.com/go-chi/chi/v5"            // Lightweight, idiomatic router for Go
	"github.com/go-chi/chi/v5/middleware" // Common middleware packages
	"github.com/go-chi/cors"              // CORS middleware for cross-origin requests
)

// Item represents a single item in our CRUD application
// Each item has a unique identifier, name, description, and an optional image
type Item struct {
	ID          string // Unique identifier for the item
	Name        string // Name of the item
	Description string // Description of the item
	ImagePath   string // Path to the uploaded image file
}

// items is our in-memory database
// In a production environment, this would be replaced with a proper database
var items = make(map[string]Item)

func main() {
	// Initialize Chi router
	r := chi.NewRouter()

	// Middleware Setup
	// Logger middleware logs all HTTP requests
	r.Use(middleware.Logger)
	// CORS middleware allows cross-origin requests - important for HTMX
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},                                       // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow all standard HTTP methods
	}))

	// Create uploads directory for storing images
	os.MkdirAll("uploads", os.ModePerm)

	// Serve static files from the uploads directory
	// This allows browsers to access uploaded images
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// Route Definitions
	// Each route maps to a specific handler function
	r.Get("/", handleIndex)                   // Serves the main page
	r.Get("/items", handleListItems)          // Returns list of items (used by HTMX)
	r.Post("/items", handleCreateItem)        // Creates a new item
	r.Delete("/items/{id}", handleDeleteItem) // Deletes an item

	// Start the server
	fmt.Println("Server starting on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

// handleIndex serves the main page of the application
// It renders the index.html template with the current list of items
func handleIndex(w http.ResponseWriter, r *http.Request) {
	// template.Must is a helper that wraps template.ParseFiles() and panics if there's an error
	// template.ParseFiles reads and parses the HTML template file from disk
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Execute renders the template, replacing any {{variables}} with values from 'items'
	// The rendered HTML is written to 'w' (http.ResponseWriter)
	tmpl.Execute(w, items)
}

// handleListItems returns a partial HTML template containing the list of items
// This is used by HTMX to dynamically update the items list without a full page reload
func handleListItems(w http.ResponseWriter, r *http.Request) {
	// template.Must panics if ParseFiles returns an error - useful during development
	// as templates should always be valid
	tmpl := template.Must(template.ParseFiles("templates/items-list.html"))

	// Execute combines the template with our 'items' data
	// In the HTML, you can access the data using {{range .}} for the items map
	tmpl.Execute(w, items)
}

// handleCreateItem processes the creation of a new item
// It handles both the form data and file upload in a single request
func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form with a maximum size of 10MB
	// This prevents users from uploading extremely large files
	r.ParseMultipartForm(10 << 20) // 10 MB max

	// Extract form values
	name := r.FormValue("name")
	description := r.FormValue("description")
	// Generate a simple incremental ID (in production, use UUID or similar)
	id := fmt.Sprintf("%d", len(items)+1)

	// Handle image file upload
	file, handler, err := r.FormFile("image")
	var imagePath string
	if err == nil {
		defer file.Close()

		// Create the file path in the uploads directory
		filename := filepath.Join("uploads", handler.Filename)
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err == nil {
			defer f.Close()
			// Copy the uploaded file to our local filesystem
			io.Copy(f, file)
			// Store the public path to the image
			imagePath = "/uploads/" + handler.Filename
		}
	}

	// Create and store the new item
	item := Item{
		ID:          id,
		Name:        name,
		Description: description,
		ImagePath:   imagePath,
	}
	items[id] = item

	// Return the updated items list as a partial HTML template
	// HTMX will use this to update the UI
	tmpl := template.Must(template.ParseFiles("templates/items-list.html"))
	tmpl.Execute(w, items)
}

// handleDeleteItem removes an item from the in-memory database
// It's triggered by HTMX delete request and returns the updated items list
func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	// chi.URLParam extracts named parameters from the URL
	// For example, from '/items/123', URLParam(r, "id") returns "123"
	id := chi.URLParam(r, "id")

	// delete is a built-in Go function that removes a key-value pair from a map
	// If the key doesn't exist, delete is a no-op (no error)
	delete(items, id)

	// Parse and execute the template just like in other handlers
	// template.Must wraps ParseFiles and panics on error
	tmpl := template.Must(template.ParseFiles("templates/items-list.html"))
	// Execute merges the template with our data and writes to the ResponseWriter
	tmpl.Execute(w, items)
}
