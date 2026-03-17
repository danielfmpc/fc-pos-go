package handlers

import (
	"encoding/json"
	"net/http"
	"pos-go-api/internal/dto"
	"pos-go-api/internal/entity"
	"pos-go-api/internal/infra/database"
	"time"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// Create user godoc
// @Summary 		Create user
// @Description Create user
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request body 			dto.CreateUserInputDto true "user request"
// @Success 		201
// @Failure 		400 		{object} Error
// @Failure 		500 		{object} Error
// @Router 			/users 	[post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInputDto
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if _, err := h.UserDB.FindByEmail(user.Email); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "User already exists"}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get JWT godoc
// @Summary 		Get JWT
// @Description Get JWT
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			request body 					dto.LoginInputDto true "user request"
// @Success 		200			{object} 			dto.LoginOutputDto
// @Failure 		400 		{object} 			Error
// @Failure 		500 		{object} 			Error
// @Router 			/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwtAuth := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiration := r.Context().Value("jwtExpiration").(int)

	var user dto.LoginInputDto
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil || !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: "Invalid credentials"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, tokenString, _ := jwtAuth.Encode(map[string]any{
		"id":    u.ID.String(),
		"email": u.Email,
		"name":  u.Name,
		"exp":   time.Now().Add(time.Second * time.Duration(jwtExpiration)).Unix(),
	})

	response := dto.LoginOutputDto{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
