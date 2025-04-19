package main

import (
	"Hackathon/backend/infrastructure/database"
	"Hackathon/backend/infrastructure/router"
	"Hackathon/backend/interface/controllers"
	"Hackathon/backend/interface/repository"
	"fmt"
	"net/http"
)

var (
	httpRouter     router.Router = router.NewMuxRouter()
	dbHandler, err               = database.NewDBHandler("")
)

func getHabitController() controllers.HabitController {
	habitRepository := repository.NewHabitRepository(dbHandler)
	habitController := controllers.NewHabitController(habitRepository)
	return *habitController
}

func main() {
	if err != nil {
		fmt.Println("Error while initializing database handler")
		fmt.Println(err)
		return
	}

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("The Application is up and running...")
	})

	habitController := getHabitController()
	httpRouter.POST("/Habit/add", habitController.AddHabit)

	httpRouter.SERVE(":8000")
}
