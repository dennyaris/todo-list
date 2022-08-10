package routes

import (
	"github.com/labstack/echo/v4"

	ps "github.com/dennyaris/todo-list/deliveries/controllers/post"
	uc "github.com/dennyaris/todo-list/deliveries/controllers/user"
	"github.com/dennyaris/todo-list/deliveries/middlewares"
)

func UserPath(e *echo.Echo, uc uc.UserController) {
	api := e.Group("/api")
	api.POST("/register", uc.Register)
	api.POST("/login", uc.Login)
	api.GET("/profile", uc.Profile, middlewares.JWTSign())
}

func PostPath(e *echo.Echo, ps ps.PostController) {
	api := e.Group("/api")
	api.POST("/posts", ps.Register, middlewares.JWTSign())
	api.GET("/posts", ps.GetAll)
	api.GET("/posts/:id", ps.GetById)
	api.PUT("/posts/:id", ps.UpdateById, middlewares.JWTSign())
	api.PUT("/posts/:id/mark", ps.UpdateMarkById, middlewares.JWTSign())
	api.DELETE("/posts/:id", ps.DeleteById, middlewares.JWTSign())
}
