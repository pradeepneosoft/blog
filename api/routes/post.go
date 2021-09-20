package routes

import (
	"blog/api/controller"
	"blog/api/service"
	"blog/infrastructure"
	"blog/middleware"
)

type PostRoute struct {
	Controller controller.PostController
	Handler    infrastructure.GinRouter
	jwtService service.JwtService
}

func NewPostRoute(
	controller controller.PostController,
	handler infrastructure.GinRouter,
	jwt service.JwtService) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler:    handler,
		jwtService: jwt,
	}
}
func (p PostRoute) Setup() {
	post := p.Handler.Gin.Group("/posts", middleware.AuthorizeJWT(p.jwtService))
	{
		post.GET("/", p.Controller.GetPosts)
		post.POST("/", p.Controller.AddPost)
		post.GET("/:id", p.Controller.GetPost)
		post.DELETE("/:id", p.Controller.DeletePost)
		post.PUT("/:id", p.Controller.UpdatePost)

	}
}
