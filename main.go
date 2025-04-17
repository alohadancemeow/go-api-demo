package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book // Slice to store books

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./views", ".html")

	// Initialize a new Fiber app
	app := fiber.New(fiber.Config{
		// Pass in Views Template Engine
		Views: engine,
	})

	// Initialize in-memory data
	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// CRUD routes
	app.Get("/books", getBooks)
	app.Get("/book/:id", getBook)
	app.Post("/book", createBook)
	app.Put("/book/:id", updateBook)
	app.Delete("/book/:id", deleteBook)

	app.Post("/upload", uploadFile)

	// Render HTML template
	app.Get("/html-test", func(c fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title":       "Go Fiber Template Example",
			"Description": "An example template",
			"Greeting":    "Hello, World!",
		})
	})

	// Get environment variables
	app.Get("/env", getEnv)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

// Upload Image file
func uploadFile(c fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully!")
}

// get environment variables
func getEnv(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET_KEY": os.Getenv("SECRET_KEY"),
	})
}
