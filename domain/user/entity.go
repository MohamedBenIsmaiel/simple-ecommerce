package userDomain
import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct{
	Id 				  primitive.ObjectID		`bson:"_id"`
	FirstName 		  *string  					`json:"firstname" validate: "required, min=2, max=50"`
	LastName          *string  					`json:"lastname"  validate: "required, min=2, max=50"`
	Email             *string					`json:"email" validate: "email, required"`
	Password          *string  					`json:"password,omitempty" validate: "required, min=6, max=14`
	Deposite          int   				    `json:"deposite"`
	Type			  *string					`json: "type" validate:"required, eq=seller|eq=buyer"`
	CreatedAt		  time.Time  				`json: "createdAt"`
	UpdatedAt		  time.Time  				`json: "updatedAt"`
}
