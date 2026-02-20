package handlers

import (
	"net/http"
	"student-grade-api/config"
	"student-grade-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("secret_key")

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	c.ShouldBindJSON(&input)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input LoginInput
	c.ShouldBindJSON(&input)

	var user models.User
	config.DB.Where("email = ?", input.Email).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}