package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"testMekarApp/main/models"
	"testMekarApp/main/usecases"
	"testMekarApp/utils"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func UserController(r *mux.Router, s usecases.UserUsecase) {
	userHandler := UserHandler{s}
	user := r.PathPrefix("/user").Subrouter()
	users := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/{id}", userHandler.User).Methods(http.MethodGet)
	user.HandleFunc("", userHandler.PostUser).Methods(http.MethodPost)
	user.HandleFunc("", userHandler.PutUser).Methods(http.MethodPut)
	user.HandleFunc("", userHandler.DeleteUser).Methods(http.MethodDelete)
	users.HandleFunc("", userHandler.Users).Methods(http.MethodGet)
}

func (s *UserHandler) User(w http.ResponseWriter, r *http.Request) {
	var userResponse utils.Response
	tes := mux.Vars(r)
	id, err := strconv.Atoi(tes["id"])
	user, err := s.UserUsecase.GetUserById(id)
	w.Header().Set("content-type", "apllication/json")
	if err != nil {
		userResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWriter(&userResponse, w)
		log.Println(err)
	} else {
		userResponse = utils.Response{Status: http.StatusOK, Message: "Success", Data: user}
		utils.ResponseWriter(&userResponse, w)
	}
	log.Println("Endpoint hit: Get User")
}

func (s *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	users, err := s.UserUsecase.GetUsers()
	var userResponse utils.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		userResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWriter(&userResponse, w)
		log.Println(err)
	} else {
		userResponse = utils.Response{Status: http.StatusOK, Message: "Get All User Success", Data: users}
		utils.ResponseWriter(&userResponse, w)
	}
	log.Println("Endpoint hit: Get All User")
}

func (s *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userResponse utils.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	err = s.UserUsecase.PostUser(&user)
	if err != nil {
		userResponse = utils.Response{Status: http.StatusBadRequest, Message: "error", Data: err}
		utils.ResponseWriter(&userResponse, w)
		log.Println(err)
	} else {
		userResponse = utils.Response{Status: http.StatusAccepted, Message: "Success", Data: user}
		utils.ResponseWriter(&userResponse, w)
	}
	log.Println("Endpoint hit: Post User")
}

func (s *UserHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userResponse utils.Response
	w.Header().Set("content-type", "application /json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	err = s.UserUsecase.PutUser(&user)
	if err != nil {
		userResponse = utils.Response{Status: http.StatusBadRequest, Message: "error", Data: err}
		utils.ResponseWriter(&userResponse, w)
		log.Println(err)
	} else {
		userResponse = utils.Response{Status: http.StatusAccepted, Message: "Success", Data: user}
		utils.ResponseWriter(&userResponse, w)
	}
	log.Println("Endpoint hit: Put User")
}

func (s *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userResponse utils.Response
	tes := mux.Vars(r)
	id, err := strconv.Atoi(tes["id"])
	err = s.UserUsecase.DeleteUser(id)
	w.Header().Set("content-type", "apllication/json")
	if err != nil {
		userResponse = utils.Response{Status: http.StatusNotFound, Message: "Not Found", Data: err.Error()}
		utils.ResponseWriter(&userResponse, w)
		log.Println(err)
	} else {
		userResponse = utils.Response{Status: http.StatusOK, Message: "Success", Data: id}
		utils.ResponseWriter(&userResponse, w)
	}
	log.Println("Endpoint hit: Delete User")
}
