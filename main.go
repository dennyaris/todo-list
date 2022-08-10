package main

import (
	"fmt"

	"github.com/dennyaris/todo-list/config"

	"github.com/dennyaris/todo-list/deliveries/helpers"
	"github.com/dennyaris/todo-list/deliveries/middlewares"
	"github.com/dennyaris/todo-list/deliveries/routes"

	cm "github.com/dennyaris/todo-list/repositories/client"
	pm "github.com/dennyaris/todo-list/repositories/post"
	um "github.com/dennyaris/todo-list/repositories/user"

	cs "github.com/dennyaris/todo-list/services/client"
	ps "github.com/dennyaris/todo-list/services/post"
	"github.com/dennyaris/todo-list/services/renderer"
	us "github.com/dennyaris/todo-list/services/user"
	"github.com/dennyaris/todo-list/services/validation"

	cc "github.com/dennyaris/todo-list/deliveries/controllers/client"
	pc "github.com/dennyaris/todo-list/deliveries/controllers/post"
	uc "github.com/dennyaris/todo-list/deliveries/controllers/user"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	middlewares.General(server)
	server.HTTPErrorHandler = helpers.CustomHTTPErrorHandler
	server.Use(session.Middleware(sessions.NewCookieStore([]byte(config.GetConfig().Secret_JWT))))

	// database connection
	db := config.InitDB(*config.GetConfig())
	config.AutoMigrate(db)

	// models
	userModel := um.NewUserModel(db)
	postModel := pm.NewPostModel(db)

	// services
	validate := validation.NewValidation()
	userService := us.NewUserService(userModel)
	postService := ps.NewPostService(postModel)

	// controllers
	userController := uc.NewUserController(userService, validate)
	postController := pc.NewPostController(postService, validate)

	// routes
	routes.UserPath(server, userController)
	routes.PostPath(server, postController)

	// Register Renderer
	server.Renderer = renderer.NewRenderer()

	// Client
	clientModel := cm.NewClientModel()
	clientService := cs.NewClientService(clientModel)
	clientController := cc.NewClientController(clientService)
	routes.ClientPath(server, clientController)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", config.GetConfig().Port)))
}
