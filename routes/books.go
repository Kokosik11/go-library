package routes

import (
	"github.com/Kokosik11/go-library/controllers" // replace
	"github.com/gofiber/fiber/v2"
)

func CatchphrasesRoute(route fiber.Router) {
	route.Get("/", controllers.GetBooks)
	// route.Get("/:id", controllers.GetCatchphrase)
	// route.Post("/", controllers.AddCatchphrase)
	// route.Put("/:id", controllers.UpdateCatchphrase)
	// route.Delete("/:id", controllers.DeleteCatchphrase)
}
