package main

import (
	"blog/api/controller"
	"blog/api/repository"
	"blog/api/routes"
	"blog/api/service"
	"blog/models"

	"blog/infrastructure"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	// router := gin.Default()
	// router.GET("/", func(context *gin.Context) {
	// 	infrastructure.LoadEnv()
	// 	infrastructure.NewDatabase()
	// 	context.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	// })
	// router.Run(":8000")
	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	// defer infrastructure.CloseDB(db.DB)

	postRepository := repository.NewRepository(db)

	jwtService := service.NewJwtService()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, jwtService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router, jwtService)
	postRoute.Setup()
	db.DB.AutoMigrate(&models.Post{}, &models.User{})

	router.Gin.Run(":8000")
}
