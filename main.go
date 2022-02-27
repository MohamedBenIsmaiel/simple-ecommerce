package main
import(
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/routes/user"

)

func main(){
	err  := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == ""{
		port = "5000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"success": "PONG !..."})
	})
	userRoutes.MapUri(router)
	router.Run(":" + port)
}