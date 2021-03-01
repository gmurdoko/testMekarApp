package main

import (
	"database/sql"
	"testMekarApp/config"
	"testMekarApp/main/controllers"
	"testMekarApp/main/repositories"
	"testMekarApp/main/usecases"

	"github.com/gorilla/mux"
)

func main() {
	db := config.ReadEnvironmentConnection()
	router := config.CreateRouter()

	Init(router, db)
	config.RunServer(router)
}

func Init(r *mux.Router, db *sql.DB) {
	userRepo := repositories.InitUserRepositoryImpl(db)
	userUsecase := usecases.InitUserusecaseImpl(userRepo)
	controllers.UserController(r, userUsecase)
	// log.Println("hhere")
}
