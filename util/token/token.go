package token
import(
	"os"
	"log"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv" 
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type signedDetails struct{
	Id  primitive.ObjectID
	Type string
	jwt.StandardClaims
}

func GenerateToken(userId  primitive.ObjectID, userType string) (bool, string){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	tokenSecret := os.Getenv("TOKEN_SECRET") // get secret from .env

	claim := &signedDetails{
		Id: userId,
		Type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(tokenSecret))
	if err != nil{
		log.Panic(err)
		return false, ""
	}
	return true, token
}