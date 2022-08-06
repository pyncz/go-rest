package tags

import (
	"context"
	"net/http"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Controllers
func Read(ctx *gin.Context) {
	collection := utils.DB.Collection("tags")

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

	var records []Tag = []Tag{}

	filter := bson.D{}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	cursor, err := collection.Find(context.TODO(), filter, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	})

	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &records); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, models.PaginatedResponse[Tag]{
		Count:   count,
		Limit:   limit,
		Offset:  offset,
		Cursor:  utils.GetNextOffset(offset, count, int64(len(records))),
		Results: records,
	})
}

func Create(ctx *gin.Context) {
	collection := utils.DB.Collection("tags")

	var record Tag

	if err := ctx.BindJSON(&record); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validate slug
	var matched Tag
	err := collection.FindOne(context.TODO(), bson.M{"slug": record.Slug}).Decode(&matched)
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

	var found Tag
	collection.FindOne(context.TODO(), bson.M{"_id": inserted.InsertedID}).Decode(&found)
	ctx.JSON(http.StatusCreated, found)
}

func Find(ctx *gin.Context) {
	collection := utils.DB.Collection("tags")

	slug := ctx.Param("slug")

	var found Tag
	err := collection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&found)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	ctx.JSON(http.StatusOK, found)
}
