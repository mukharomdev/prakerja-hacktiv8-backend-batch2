package routers



import(
	"github.com/gin-gonic/gin"

	"uk-hacktiv8-prakerja/controllers"
	"uk-hacktiv8-prakerja/middlewares"
)

var controller controllers.UserController
var Photocontroller controllers.PhotoController

func Routers(engine *gin.Engine)*gin.Engine{
	// user
	UserRouter := engine.Group("users")

	UserRouter.POST("/register",controller.Register)
	UserRouter.POST("/login",controller.Login)
	UserRouter.PUT("/:userId",middlewares.UserAuthentication,controller.Update)
	UserRouter.DELETE("/",middlewares.UserAuthentication,controller.Delete)

	// photo
	PhotoRouter := engine.Group("photos")

	PhotoRouter.POST("/",middlewares.UserAuthentication,Photocontroller.Create)
	PhotoRouter.GET("/",middlewares.UserAuthentication,Photocontroller.GetAllPhotos)

	return engine
}