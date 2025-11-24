package main

import (
	"event-api-app/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=2"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginResponse struct {
	Token string        `json:"token"`
	User  database.User `json:"user"`
}

// login authenticates user and returns JWT token
//
// @Summary User login
// @Description Authenticate user with email and password
// @Tags authentication
// @Accept json
// @Produce json
// @Param credentials body loginRequest true "User login credentials"
// @Success 200 {object} database.User
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (app *application) login(c *gin.Context) {
	var auth loginRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := app.models.Users.GetByEmail(auth.Email)

	if existingUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(auth.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.Id,
		"expr":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(app.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, not able to generate token"})
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: tokenString, User: *existingUser})
}

// registerUser creates a new user account
//
// @Summary User registration
// @Description Create a new user account
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body registerRequest true "User registration data"
// @Success 201 {object} loginResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (app *application) registerUser(c *gin.Context) {
	var register registerRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	register.Password = string(hashedPassword)

	user := database.User{
		Email:    register.Email,
		Password: register.Password,
		Name:     register.Name,
		Role:     "user", // Default role is "user"
	}

	err = app.models.Users.Insert(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

type updateUserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Name     string `json:"name" binding:"omitempty,min=2"`
	Password string `json:"password" binding:"omitempty,min=8"`
}

// updateUser updates user information
//
// @Summary Update user information
// @Description Update user email, name, and password
// @Tags user
// @Accept json
// @Produce json
// @Param user body updateUserRequest true "User update data"
// @Success 200 {object} database.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/user [put]
func (app *application) updateUser(c *gin.Context) {
	var updateReq updateUserRequest

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("user_id").(int)

	updatedUser, err := app.models.Users.Update(userId, updateReq.Email, updateReq.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update user"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
