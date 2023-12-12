package routes

import (
	"template-starter-mysql/app/controller"
	"template-starter-mysql/app/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db, err := repository.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inisialisasi repository dan handler
	userRepo := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepo)

	api := app.Group("/api")

	userRoute := api.Group("/users")
	userRoute.Get("/", userController.GetAllUsers)

}