package tasks

import (
	"context"
	"net/http"
	"pyncz/go-rest/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Controllers
func Read(ctx *gin.Context) {
	collection := utils.DB.Collection("tasks")

	var records []Task

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
	collection := utils.DB.Collection("tasks")

	var record Task

	if err := ctx.BindJSON(&record); err != nil {
		return
	}

	// TODO: Check if need to validate ID
	// var findRes bson.M
	// err := collection.FindOne(context.TODO(), bson.D{{"id", record.ID}}).Decode(&findRes)
	// if err != mongo.ErrNoDocuments {
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "'slug' is not unique"})
	// 	return
	// }

	inserted, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusCreated, inserted)
}

func Find(ctx *gin.Context) {
	collection := utils.DB.Collection("tasks")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect path param 'id'"})
		return
	}

	var found bson.M
	err = collection.FindOne(context.TODO(), id).Decode(&found)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	ctx.JSON(http.StatusOK, found)
}
