package items

import (
	"context"
	"net/http"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Controllers
func Read(env *models.AppEnv) func(*gin.Context) {
	collection := env.DB.Collection("items")

	return func(ctx *gin.Context) {
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

		var records []Item = []Item{}

		filter := bson.D{}
		count, err := collection.CountDocuments(context.TODO(), filter)
		if err != nil {
			panic(err)
		}
		opts := options.Find().SetLimit(limit).SetSkip(offset)
		cursor, err := collection.Find(context.TODO(), filter, opts)

		if err != nil {
			panic(err)
		}
		if err = cursor.All(context.TODO(), &records); err != nil {
			panic(err)
		}
		defer cursor.Close(context.TODO())

		ctx.JSON(http.StatusOK, models.PaginatedResponse[Item]{
			Count:   count,
			Limit:   limit,
			Offset:  offset,
			Cursor:  utils.GetNextOffset(offset, count, int64(len(records))),
			Results: records,
		})
	}
}

func Create(env *models.AppEnv) func(*gin.Context) {
	collection := env.DB.Collection("items")

	return func(ctx *gin.Context) {
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
}

func Find(env *models.AppEnv) func(*gin.Context) {
	collection := env.DB.Collection("items")

	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
			return
		}

		var found Item
		err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&found)
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
			return
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, found)
	}
}
