package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/leandroudala/golang_jwt/api/database"
	"github.com/leandroudala/golang_jwt/api/models"
	"github.com/leandroudala/golang_jwt/api/repository"
	"github.com/leandroudala/golang_jwt/api/repository/crud"
	"github.com/leandroudala/golang_jwt/api/responses"
)

// All List all Users
func All(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
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
		user, err = userRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

// Update a user
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

// Delete a user
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}

// Get a user
func Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a user"))
}
