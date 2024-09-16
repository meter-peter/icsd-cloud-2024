package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meter-peter/icsd-cloud-2024/internal/models"
	"github.com/meter-peter/icsd-cloud-2024/internal/validators"
	"gorm.io/gorm"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		result := db.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetUser godoc
// @Summary Get a single user
// @Description Get details of a specific user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		var user models.User
		result := db.First(&user, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User details"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validators.ValidateUser(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&newUser)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusCreated, newUser)
	}
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update details of an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "Updated user details"
// @Success 200 {object} models.User
// @Failure 400,404 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		var user models.User
		result := db.First(&user, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		var updateData models.User
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.IdentityNumber = updateData.IdentityNumber
		user.FirstName = updateData.FirstName
		user.LastName = updateData.LastName
		user.Gender = updateData.Gender
		user.Addresses = updateData.Addresses
		user.PhoneNumbers = updateData.PhoneNumbers
		if err := validators.ValidateUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&user)
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400,404 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		result := db.Delete(&models.User{}, id)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

// SearchUsers godoc
// @Summary Search for users
// @Description Search for users based on various criteria
// @Tags users
// @Accept  json
// @Produce  json
// @Param identity_number query string false "Identity Number"
// @Param first_name query string false "First Name"
// @Param last_name query string false "Last Name"
// @Param gender query string false "Gender"
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]string
// @Router /users/search [get]
func SearchUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		query := db.Order("last_name, first_name, identity_number")

		if identityNumber := c.Query("identity_number"); identityNumber != "" {
			query = query.Where("identity_number LIKE ?", "%"+identityNumber+"%")
		}
		if firstName := c.Query("first_name"); firstName != "" {
			query = query.Where("first_name LIKE ?", "%"+firstName+"%")
		}
		if lastName := c.Query("last_name"); lastName != "" {
			query = query.Where("last_name LIKE ?", "%"+lastName+"%")
		}
		if gender := c.Query("gender"); gender != "" {
			query = query.Where("gender = ?", gender)
		}

		result := query.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
