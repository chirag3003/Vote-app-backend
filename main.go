package main

import (
	"fmt"
	"github.com/chirag3003/vote-back/controllers"
	"github.com/chirag3003/vote-back/database"
	"github.com/chirag3003/vote-back/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := &database.Database{}
	db.Connect()

	app := fiber.New()

	routers.Setup(controllers.Setup(db), app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
