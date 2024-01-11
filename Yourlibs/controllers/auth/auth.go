package auth

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/libraryManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UsersInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// registrasi admin
// POST /admin/register
// func RegisterAdmin(c *gin.Context) {
// 	var input UsersInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// autogenerated id, created_at, is_deleted
// 	id := uuid.New().String()
// 	id2 := uuid.New().String()
// 	created_at := time.Now().UTC().Add(7 * time.Hour)

// 	// check validity of username
// 	var existingUsername models.Users
// 	if err := models.DB.Where("username = ?", input.Username).First(&existingUsername).Error; err == nil {
// 		c.JSON(http.StatusConflict, gin.H{"error": "Username already in use"})
// 		return
// 	}

// 	// check validity of password
// 	if len(input.Password) < 6 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "password must containt at least 6 characters"})
// 		return
// 	}

// 	// check validity of token
// 	if input.Token != "nietzsche" {
// 		c.JSON(http.StatusConflict, gin.H{"error": "Token invalid"})
// 		return
// 	}

// 	dataUser := models.Users{
// 		ID:        id,
// 		CreatedAt: created_at,
// 		IsDeleted: false,
// 		// Username:  input.Username,
// 		Password: input.Password,
// 		// Role:      true,
// 	}

// 	dataProfile := models.Profile{
// 		ID:      id2,
// 		ID_User: id,
// 	}

// 	models.DB.Create(&dataUser)
// 	models.DB.Create(&dataProfile)

// 	c.JSON(http.StatusCreated, gin.H{"data": dataUser})
// }

// registrasi user
// POST /user/register
func RegisterUser(c *gin.Context) {
	var input UsersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// autogenerated id, created_at, is_deleted
	id := uuid.New().String()
	created_at := time.Now().UTC().Add(7 * time.Hour)

	// check validity of email
	var existingUsername models.Users
	if err := models.DB.Where("username = ?", input.Email).First(&existingUsername).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already in use"})
		return
	}

	// check validity of password
	if len(input.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must containt at least 6 characters"})
		return
	}

	dataUser := models.Users{
		ID:        id,
		CreatedAt: created_at,
		IsDeleted: false,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
	}

	models.DB.Create(&dataUser)

	c.JSON(http.StatusCreated, gin.H{"data": dataUser})
}