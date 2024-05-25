package routers



import(
	"github.com/gin-gonic/gin"

	"uk-hacktiv8-prakerja/controllers"
	"uk-hacktiv8-prakerja/middlewares"
)

var controller controllers.UserController

func UserRouters(engine *gin.Engine)*gin.Engine{

	UserRouter := engine.Group("users")

	UserRouter.POST("/register",controller.Register)
	UserRouter.POST("/login",controller.Login)
	UserRouter.PUT("/:userId",middlewares.UserAuthentication,controller.Update)
	UserRouter.DELETE("/",controller.Delete)

	return engine
}