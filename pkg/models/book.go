package models

//primitive.ObjectIDFromHex("5eb3d668b31de5d588f42a7a")
import (
	"context"
	"fmt"
	"log"

	"github.com/Shashank-Panda/crud-mongoDB/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var BookCollection = &mongo.Collection{}

type Book struct {
	Id     int    `json:"id" bson:"_id"`
	Title  string `json:"title" bson:"_title"`
	Author string `json:"author" bson:"_author"`
	Isbn   string `json:"isbn" bson:"_isbn"`
}

// var Books []Book

func init() {
	config.Connect()
	db = config.GetDB()
	BookCollection = db.Collection("books")
}

func GetAllBooks() []Book {
	var results []Book
	findOptions := options.Find()
	cur, err := BookCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Here cur is a cursor which provides a stream of documents
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Book
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}

func GetBookById(id int64) (Book, *mongo.Database) {
	var result Book
	err := BookCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	return result, db
}

func (b *Book) CreateBook() *Book {
	insertResult, err := BookCollection.InsertOne(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return b
}

func DeleteBook(id int64) {
	deleteResult, err := BookCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func UpdateBook(id int64, book Book) Book {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"_title", book.Title}}}, {"$set", bson.D{{"_author", book.Author}}}, {"$set", bson.D{{"_isbn", book.Isbn}}}}
	result, err := BookCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", result)
	var ret Book
	err = BookCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&ret)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
