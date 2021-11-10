package main

//go:generate go run github.com/arsmn/fastgql
import (
	"fmt"
	"log"
	"os"

	"besrabasant/fiber/graphql"
	"besrabasant/fiber/graphql/generated"

	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graphql.Resolver{}},
		),
	)
	gqlHandler := srv.Handler()
	playgroundHandler := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	app.All("/graphql-console", func(c *fiber.Ctx) error {
		playgroundHandler(c.Context())
		return nil
	})

	appPort := os.Getenv("APP_PORT")

	log.Fatal(app.Listen(fmt.Sprintf(":%s", appPort)))
}
