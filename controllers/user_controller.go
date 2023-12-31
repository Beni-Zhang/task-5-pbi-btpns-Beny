package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	uc.DB.Create(&user)
	c.JSON(201, gin.H{"message": "User created successfully"})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" valid:"email,required"`
		Password string `json:"password" valid:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(loginData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := uc.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

func (uc *UserController) UpdateUser(c *gin.Context) {

	userId := c.Param("userId")

	var user User
	if err := uc.DB.First(&user, userId).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	claims, _ := c.Get("claims")
	authUserId := uint(claims.(jwt.MapClaims)["userId"].(float64))
	if authUserId != user.ID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uc.DB.Save(&user)
	c.JSON(200, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {

	userId := c.Param("userId")

	var user User
	if err := uc.DB.First(&user, userId).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	claims, _ := c.Get("claims")
	authUserId := uint(claims.(jwt.MapClaims)["userId"].(float64))
	if authUserId != user.ID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	uc.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}