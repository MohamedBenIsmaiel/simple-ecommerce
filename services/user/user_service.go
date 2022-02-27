package userServices
import(
	"github.com/MohamedBenIsmaiel/simple-ecommerce/util/error"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/domain/user"
)

func Signup(user *userDomain.User)(*userDomain.User, *handlError.HandleError){
	if err:= user.Validate(); err != nil{
		return nil, err
	}

	if err:= user.Save();err != nil{
		return nil, err
	}

	return user, nil
}

func Login(user *userDomain.User)(*userDomain.User, *handlError.HandleError){
	if err:= user.Validate(); err != nil{
		return nil, err
	}

	if err:= user.Login();err != nil{
		return nil, err
	}

	return user, nil
}