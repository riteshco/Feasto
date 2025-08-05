package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


func Run(){
	app := fiber.New()

	app.Get("/" , func (c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	fmt.Printf("Listening on http://localhost:3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start the http server : %v" , err)
	}	
}