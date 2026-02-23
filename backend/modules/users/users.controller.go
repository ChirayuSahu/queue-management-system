package users

import (
	"errors"

	"github.com/chirayusahu/queue-management-system/backend/common"
	"github.com/chirayusahu/queue-management-system/backend/config"
	"github.com/chirayusahu/queue-management-system/backend/database"
	"github.com/chirayusahu/queue-management-system/backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"
)

func CreateUser(c *fiber.Ctx) error {

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

	return common.Respond(c, fiber.StatusCreated, true, "User created successfully", nil)
}

func GetAllUsers(c *fiber.Ctx) error {

	var users []models.User

	query := database.DB

	isAdmin := c.Query("is_admin")

	if isAdmin != "" {
		query = query.Where("is_admin = ?", isAdmin)
	}

	if err := query.Find(&users).Error; err != nil {
		return common.Respond(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	if len(users) == 0 {
		return common.Respond(c, fiber.StatusNotFound, false, "No users found", nil)
	}

	return common.Respond(c, fiber.StatusOK, true, "Users retrieved successfully", users)
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
		"id": user.ID.String(),
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

func DeleteUser(c *fiber.Ctx) error {

	id := c.Params("id")
	query := database.DB

	var user models.User

	if err := query.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.Respond(c, fiber.StatusNotFound, false, "User not found", nil)
		}
		return common.Respond(c, fiber.StatusInternalServerError, false, "Failed to retrieve user", nil)
	}

	if err := query.Delete(&user).Error; err != nil {
		return common.Respond(c, fiber.StatusInternalServerError, false, "Failed to delete user", nil)
	}

	return common.Respond(c, fiber.StatusOK, true, "User deleted successfully", nil)
}
