package tasks

import (
	"context"
	"log"
	"net/http"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Controllers
func Read(ctx *gin.Context) {
	collection := utils.DB.Collection("tasks")

	limit, err := utils.ExtractInt64Query(ctx, "limit", models.DEFAULT_LIMIT)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect query param 'limit'"})
		return
	}

	offset, err := utils.ExtractInt64Query(ctx, "offset", 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect query param 'offset'"})
		return
	}

	var records []Task = []Task{}

	cursor, err := collection.Find(context.TODO(), bson.D{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &records); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, records)
}

func Create(ctx *gin.Context) {
	collection := utils.DB.Collection("tasks")

	var record Task

	if err := ctx.BindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inserted, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		panic(err)
	}

	var found Task
	collection.FindOne(context.TODO(), bson.M{"_id": inserted.InsertedID}).Decode(&found)
	ctx.JSON(http.StatusCreated, found)
}

func Find(ctx *gin.Context) {
	collection := utils.DB.Collection("tasks")

	id := ctx.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	var found Task
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&found)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	ctx.JSON(http.StatusOK, found)
}
