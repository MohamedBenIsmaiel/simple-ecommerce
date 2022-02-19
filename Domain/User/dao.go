package User
import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct{
	Id 				  primitive.ObjectID		`bson:"_id"`
	FirstName 		  *string  					`json:"firstName" validate: "required, min=2, max=50"`
	LastName          *string  					`json:"lastName"  validate: "required, min=2, max=50"`
	Email             *string					`json:"email" validate: "email, required"`
	Password          *string  					`json:"password" validate: "required, min=6, max=14`
	Deposite          *uint   				    `json:"deposite"`
	Type			  *string					`json: "type" validate:"required, eq=SELLER|eq=BUYER"`
	CreatedAt		  time.Time  				`json: "createdAt"`
	UpdatedAt		  time.Time  				`json: "updatedAt"`
}
