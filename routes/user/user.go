package userRoutes
import(
	"github.com/MohamedBenIsmaiel/simple-ecommerce/controllers/user"
	"github.com/gin-gonic/gin"
)

func MapUri(router *gin.Engine){
	router.POST("/signup",userControllers.Signup)
	router.POST("/login",userControllers.Login)
}

