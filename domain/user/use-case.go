package userDomain
import(
	"time"
	"context"
	"fmt"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/databases/mongodb"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/util/error"
	"github.com/MohamedBenIsmaiel/simple-ecommerce/util/password"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-playground/validator/v10"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func(user *User)Save() *handlError.HandleError{
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

	var getUser *User
	userCollection.FindOne(ctx, bson.M{"email": *user.Email}).Decode(&getUser)
	defer cancel()

	if getUser != nil{
		err := handlError.BadRequest(fmt.Sprintf("User with Email %s already exist", *user.Email))
		return err
	}


	user.Id = primitive.NewObjectID()
    user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))	
    user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))	
	user.Deposite = 10000 // default value for every user

	defaultUserType := "buyer"
	user.Type = &defaultUserType
	getHashedPassword := password.HashPassword(*user.Password)
	user.Password = &getHashedPassword


	_, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil{
		err := handlError.InternalServerError(insertErr.Error())
		return err
	}

	user.Password = nil // don't return password on the response

	defer cancel()
	return nil
}

func(user *User)Login() *handlError.HandleError{
	var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

	var foundUser *User
	userCollection.FindOne(ctx, bson.M{"email": *user.Email}).Decode(&foundUser)
	defer cancel()

	if  foundUser ==  nil || foundUser.Email == nil{
		err := handlError.BadRequest("User Email or Password doesn't match !")
		return err
	}


	isValidPassword := password.VerifyPassword(*user.Password, *foundUser.Password)
	if isValidPassword == false{
		err := handlError.BadRequest("User Email or Password doesn't match !")
		return err
	}

	user.Id = foundUser.Id
	user.Type = foundUser.Type
	
	return nil
}

func(user *User)Get(){

}

func (user User)Validate() *handlError.HandleError{
	validationErr := validate.Struct(user)
	if validationErr != nil{
		err := handlError.BadRequest(validationErr.Error())
		return err;
	}
	return nil
}


