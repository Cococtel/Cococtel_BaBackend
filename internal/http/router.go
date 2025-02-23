package http

import (
	"database/sql"
	"github.com/Cococtel/Cococtel_BaBackend/internal/controllers"
	"github.com/Cococtel/Cococtel_BaBackend/internal/controllers/usercontroller"
	"github.com/Cococtel/Cococtel_BaBackend/internal/defines"
	"github.com/Cococtel/Cococtel_BaBackend/internal/middleware"
	"github.com/Cococtel/Cococtel_BaBackend/internal/repository/userrepository"
	"github.com/Cococtel/Cococtel_BaBackend/internal/services/userservice"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.addSystemPaths()
	r.buildRoutes()
}

func (r *router) setGroup() {
	r.eng.Use(middleware.CORS())
	r.rg = r.eng.Group("/v1", middleware.ProtectedHandler())
}

func (r *router) buildRoutes() {
	userRepository := userrepository.NewUserRepository(r.db)

	userService := userservice.NewUser(userRepository)

	userController := usercontroller.NewUser(userService)

	r.rg.POST(defines.VerifyPath, userController.VerifyUser())
	r.eng.POST(defines.RegisterPath, userController.RegisterUser())
	r.eng.POST(defines.LoginPath, userController.LoginUser())
	r.eng.POST(defines.ValidateLoginPath, userController.ValidateLogin())
	r.rg.POST(defines.GetQRDoubleAuthPath, userController.GetQRDoubleAuth())
	r.rg.POST(defines.NotifyQRReadPath, userController.NotifyQRRead())
	r.rg.GET(defines.ProfilePath+defines.IDPath, userController.GetUser())
	r.rg.PUT(defines.ProfilePath, userController.EditProfile())
}
func (r *router) addSystemPaths() {
	r.eng.GET(defines.PingPath, controllers.Ping())
}
