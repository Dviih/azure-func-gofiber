package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("*", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")

		return c.Send([]byte(fmt.Sprintf(`{"hello": "world", "from": "%s", "env": "%s", "rootdir": "%s"}`, os.Getenv("FUNCTIONS_CUSTOMHANDLER_WORKER_ID"), os.Getenv("AZURE_FUNCTIONS_ENVIRONMENT"), os.Getenv("FUNCTIONS_APP_ROOT_PATH"))))
	})

	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT"))); err != nil {
		panic(err)
	}
}
