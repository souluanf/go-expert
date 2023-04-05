package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/souluanf/go-expert/7-apis/internal/dto"
	"github.com/souluanf/go-expert/7-apis/internal/entity"
	"github.com/souluanf/go-expert/7-apis/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	JwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !p.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub":   p.ID.String(),
		"exp":   time.Now().Add(time.Second * time.Duration(JwtExpiresIn)).Unix(),
		"email": p.Email,
	})
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
