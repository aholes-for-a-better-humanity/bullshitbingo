/*
Package server provides a simple Go server to prototype against
*/
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Static("/", ".")
	app.Listen(":8080")
}
