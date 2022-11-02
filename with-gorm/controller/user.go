package controller

import (
	"errors"
	"mod-go-gorm/database"
	"mod-go-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Users{})

	return &UserRepo{
		Db: db,
	}
}

// create user
func (respositoryUser *UserRepo) CreateUser(c *gin.Context) {
	var user models.Users
	c.BindJSON(&user)
	err := models.CreateUser(respositoryUser.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// get users
func (respositoryUser *UserRepo) GetUsers(c *gin.Context) {
	var user []models.Users
	err := models.GetUsers(respositoryUser.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// get user by email
func (respositoryUser *UserRepo) GetUser(c *gin.Context) {
	email := c.Param("email")
	var user models.Users
	err := models.GetUser(respositoryUser.Db, &user, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// update user
func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user models.Users
	email := c.Param("email")
	err := models.GetUser(repository.Db, &user, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = models.UpdateUser(repository.Db, &user, email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// delete user
func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user models.Users
	email := c.Param("email")
	_, err := models.DeleteUser(repository.Db, &user, email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
