package items

import (
	"context"
	"log"
	"net/http"
	"pyncz/go-rest/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controllers
func Read(ctx *gin.Context) {
	collection := utils.DB.Collection("items")

	var records []Item

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &records); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, records)
}

func Create(ctx *gin.Context) {
	collection := utils.DB.Collection("items")

	var record Item

	if err := ctx.BindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inserted, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		panic(err)
	}

	var found Item
	collection.FindOne(context.TODO(), bson.M{"_id": inserted.InsertedID}).Decode(&found)
	ctx.JSON(http.StatusCreated, found)
}

func Find(ctx *gin.Context) {
	collection := utils.DB.Collection("items")

	id := ctx.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	var found Item
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&found)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	ctx.JSON(http.StatusOK, found)
}
