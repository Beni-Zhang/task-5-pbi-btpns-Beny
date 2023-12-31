package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	"github.com/asaskevich/govalidator"
)

type PhotoController struct {
	DB *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{DB: db}
}

func (pc *PhotoController) AddPhoto(c *gin.Context) {
	var photo Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	claims, _ := c.Get("claims")
	userId := uint(claims.(jwt.MapClaims)["userId"].(float64))
	photo.UserID = userId

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Create(&photo)
	c.JSON(201, gin.H{"message": "Photo added successfully"})
}

func (pc *PhotoController) GetPhotos(c *gin.Context) {
	var photos []Photo
	pc.DB.Find(&photos)
	c.JSON(200, photos)
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {

	photoId := c.Param("photoId")

	var photo Photo
	if err := pc.DB.First(&photo, photoId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Photo not found"})
		return
	}

	claims, _ := c.Get("claims")
	authUserId := uint(claims.(jwt.MapClaims)["userId"].(float64))
	if authUserId != photo.UserID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Save(&photo)
	c.JSON(200, gin.H{"message": "Photo updated successfully"})
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {

	photoId := c.Param("photoId")

	var photo Photo
	if err := pc.DB.First(&photo, photoId).Error; err != nil {
		c.JSON(404, gin.H{"error": "Photo not found"})
		return
	}

	claims, _ := c.Get("claims")
	authUserId := uint(claims.(jwt.MapClaims)["userId"].(float64))
	if authUserId != photo.UserID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	pc.DB.Delete(&photo)
	c.JSON(200, gin.H{"message": "Photo deleted successfully"})
}