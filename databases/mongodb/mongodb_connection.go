package database
import(
	"os"
	"time"
	"log"
	"fmt"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongodb() *mongo.Client{
	loadErr := godotenv.Load(".env")
	if loadErr != nil{
		log.Fatal(loadErr)
	}

	mongoUrl := os.Getenv("MONGO_URL")
	client, err :=  mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil{
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	defer cancel()
	fmt.Println("Mongodb connection successed ..")
	return client
}

var Client *mongo.Client = connectMongodb()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	mongodbName := os.Getenv("MONGO_DB_NAME")
	var collection *mongo.Collection = client.Database(mongodbName).Collection(collectionName)
	return collection
}