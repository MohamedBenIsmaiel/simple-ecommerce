package password
import(
	"golang.org/x/crypto/bcrypt"
	"log"
)

func VerifyPassword(userPassword, providedPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil{
		return false
	}
	return true
}

func HashPassword(password string) string{
	passwrodBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err!= nil{
		log.Panic(err)
	}
	return string(passwrodBytes)

}