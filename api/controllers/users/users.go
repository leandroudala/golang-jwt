package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leandroudala/golang_jwt/api/database"
	"github.com/leandroudala/golang_jwt/api/models"
	"github.com/leandroudala/golang_jwt/api/repository"
	"github.com/leandroudala/golang_jwt/api/repository/crud"
	"github.com/leandroudala/golang_jwt/api/responses"
)

// All List all Users
func All(w http.ResponseWriter, r *http.Request) {
	// connecting to the database
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// inserting into the database
	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		users, err := userRepository.All()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.JSON(w, http.StatusOK, users)
	}(repo)
}

// Create a new User
func Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	// loading request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// converting to json
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// connecting to the database
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// inserting into the database
	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		var status int = http.StatusCreated
		user, status, err = userRepository.Save(user)
		if err != nil {
			responses.ERROR(w, status, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, status, user)
	}(repo)
}

// Update a user
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	publicID := vars["public_id"]

	user := models.User{}
	// loading request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// converting to json
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// setting the public ID
	user.PublicID = publicID

	// connecting to the database
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// inserting into the database
	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		var status int = 200
		user, status, err = userRepository.Update(user)
		if err != nil {
			responses.ERROR(w, status, err)
			return
		}
		responses.JSON(w, status, user)
	}(repo)
}

// Delete a user
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// connecting to the database
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// inserting into the database
	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		var status int
		status, err = userRepository.Delete(vars["public_id"])
		if err != nil {
			responses.ERROR(w, status, err)
			return
		}
		responses.JSON(w, http.StatusOK, nil)
	}(repo)
}

// Get a user
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := models.User{}
	// connecting to the database
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// inserting into the database
	repo := crud.NewRepositoryUsersCRUD(db)

	func(userRepository repository.UserRepository) {
		var status int
		user, status, err = userRepository.FindByID(vars["public_id"])
		if err != nil {
			responses.ERROR(w, status, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)
	}(repo)
}
