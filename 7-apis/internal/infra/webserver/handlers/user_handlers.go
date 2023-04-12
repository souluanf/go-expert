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

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// GetJWT get jwt
// @Summary      get user JWT
// @Description  get user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.GetJWTInput   true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
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
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	if !p.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub":   p.ID.String(),
		"exp":   time.Now().Add(time.Second * time.Duration(JwtExpiresIn)).Unix(),
		"email": p.Email,
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.CreateUserInput   true  "user request"
// @Success      201
// @Failure      500  {object}  Error
// @Router       /users [post]
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
		badError := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(badError)
		return
	}
	err = h.UserDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		internalError := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(internalError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
