package main

import (
	"webapp/e-commerce/config"
	"webapp/e-commerce/routes"
)

func main() {
	// Configure and connect to DB
	config.ConnectToDB()
	// Register routes
	r := routes.RegisterRoutes()
	r.Run(":9000")
}
