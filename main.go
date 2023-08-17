package main

import (
	"log"

	"github.com/dangrichardyang/htmx-go_movie-app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
	}
	app := fiber.New()
	
	app.Get("/", handlers.Home)

	log.Fatal(app.Listen("127.0.0.1:5000"))
}