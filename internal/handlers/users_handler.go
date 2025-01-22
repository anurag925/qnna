// internal/handlers/auth.go

package handlers

import (
	"net/http"
	"time"

	"github.com/anurag925/qnna/internal/models"
	"github.com/anurag925/qnna/internal/utils"
	"github.com/anurag925/qnna/templates"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) SignUp(c echo.Context) error {
	if h.isGet(c) {
		return h.render(c, http.StatusOK, templates.SignUp())
	}
	var userInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Username string `json:"username" validate:"required"`
		Mobile   string `json:"mobile" validate:"required"`
		Age      int    `json:"age" validate:"required,min=0"`
	}
	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	user := &models.User{
		Email:    userInput.Email,
		Password: string(hashedPassword),
		Username: userInput.Username,
		Mobile:   userInput.Mobile,
		Age:      userInput.Age,
	}
	if err := h.userRepo.DB.Insert(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (h *handler) Login(c echo.Context) error {
	var loginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}
	if err := c.Bind(&loginInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&loginInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var user models.User
	if err := h.userRepo.DB.NewSelect().Model(&user).Where("email = ?", loginInput.Email).Scan(c.Request().Context()); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(utils.GetJwtKey())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
