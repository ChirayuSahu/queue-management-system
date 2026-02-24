package auth

import (
	"errors"
	"time"

	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/chirayusahu/queue-management-system/backend/database"
	"github.com/chirayusahu/queue-management-system/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {

	type UserBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user UserBody

	if err := c.BodyParser(&user); err != nil {
		return common.Respond(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	var existingUser models.User

	err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error

	if err == nil {
		return common.Respond(c, fiber.StatusConflict, false, "User with this email already exists", nil)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return common.Respond(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.Respond(c, fiber.StatusInternalServerError, false, "Failed to hash password", nil)
	}

	hashedPasswordStr := string(hashedPassword)
	newUser := models.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: &hashedPasswordStr,
		AuthProvider: models.ProviderCredentials,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return common.Respond(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return common.Respond(c, fiber.StatusCreated, true, "User registered successfully", nil)
}

func LoginUser(c *fiber.Ctx) error {

	type LoginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body LoginBody

	if err := c.BodyParser(&body); err != nil {
		return common.Respond(c, fiber.StatusBadRequest, false, "Invalid request body", nil)
	}

	var user models.User

	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.Respond(c, fiber.StatusUnauthorized, false, "Invalid email or password", nil)
		}
		return common.Respond(c, fiber.StatusInternalServerError, false, "Failed to retrieve user", nil)
	}

	if user.AuthProvider != models.ProviderCredentials {
		return common.Respond(c, fiber.StatusUnauthorized, false, "User does not use credentials authentication", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(body.Password)); err != nil {
		return common.Respond(c, fiber.StatusUnauthorized, false, "Invalid email or password", nil)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID.String(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte(config.LoadConfig().JWTSecret)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return common.Respond(c, fiber.StatusInternalServerError, false, "Failed to generate token", nil)
	}

	return common.Respond(c, fiber.StatusOK, true, "Login successful", fiber.Map{
		"token": tokenString,
	})
}
