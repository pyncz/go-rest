package tags

import (
	"context"
	"net/http"
	"pyncz/go-rest/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Controllers
func Read(ctx *gin.Context) {
	collection := utils.DB.Collection("tags")

	var records []Tag

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
	collection := utils.DB.Collection("tags")

	var record Tag

	if err := ctx.BindJSON(&record); err != nil {
		return
	}

	// Validate slug
	var found bson.M
	err := collection.FindOne(context.TODO(), record.ID).Decode(&found)
	if err != mongo.ErrNoDocuments {
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "'slug' is not unique"})
		return
	}

	inserted, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusCreated, inserted)
}

func Find(ctx *gin.Context) {
	collection := utils.DB.Collection("tags")

	slug := ctx.Param("slug")

	var found bson.M
	err := collection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&found)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	ctx.JSON(http.StatusOK, found)
}
