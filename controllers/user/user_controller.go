package userControllers
import(
	"net/http"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/services/user"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/domain/user"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/util/error"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/util/token"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context){
	var user userDomain.User
	err := c.ShouldBindJSON(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, handlError.BadRequest("Invalid Json"))
		return
	}

	result, saveError := userServices.Signup(&user)
	if saveError != nil{
		c.JSON(saveError.Code,saveError)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func Login(c *gin.Context){
	var user userDomain.User
	err := c.ShouldBindJSON(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, handlError.BadRequest("Invalid Json"))
		return
	}

	result, loginError := userServices.Login(&user)
	if loginError != nil{
		c.JSON(loginError.Code,loginError)
		return
	}

	isGenerated, getToken := token.GenerateToken(result.Id, *result.Type)
	if isGenerated == false{
	   c.JSON(http.StatusInternalServerError, handlError.InternalServerError("Something wrong happened try later"))
	}
	
	c.JSON(http.StatusCreated, map[string]string{"message": "success", "token": getToken})

}