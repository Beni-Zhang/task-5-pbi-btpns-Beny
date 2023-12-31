package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Beni-Zhang/task-5-pbi-btpns-Beny/controllers"
)

func SetupRoutes(r *gin.Engine, uc *controllers.UserController, pc *controllers.PhotoController) {
	// User endpoints
	r.POST("/users/register", uc.RegisterUser)
	r.POST("/users/login", uc.LoginUser)
	r.PUT("/users/:userId", uc.UpdateUser)
	r.DELETE("/users/:userId", uc.DeleteUser)

	// Photo endpoints
	r.POST("/photos", pc.AddPhoto)
	r.GET("/photos", pc.GetPhotos)
	r.PUT("/photos/:photoId", pc.UpdatePhoto)
	r.DELETE("/photos/:photoId", pc.DeletePhoto)
}