package routers



import(
	"github.com/gin-gonic/gin"

	"uk-hacktiv8-prakerja/controllers"
)

var controller controllers.UserController

func UserRouter(engine *gin.Engine)*gin.Engine{

	router := engine.Group("users")

	router.POST("/register",controller.Register)
	router.POST("/login",controller.Login)

	return engine
}