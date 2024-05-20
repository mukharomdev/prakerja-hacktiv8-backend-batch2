package routers


import(

	"prakerja-session-6/controllers"
	"github.com/gin-gonic/gin"
)


func StartServer() *gin.Engine{
	router := gin.Default()

     router.GET("/",controllers.Halo)
     router.GET("/products",controllers.GetProducts)

	 router.POST("/products",controllers.CreateProducts)

	// router.PUT("/products/:id",controllers.UpdateProducts)

	// router.DELETE("/products/:id",controllers.DeleteProducts)

	return router
}