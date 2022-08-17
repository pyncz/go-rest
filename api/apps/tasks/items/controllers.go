package items

import (
	"context"
	"net/http"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Controllers
func Read(env *models.AppEnv) func(*fiber.Ctx) error {
	collection := env.DB.Collection("items")

	return func(ctx *fiber.Ctx) error {
		limit, err := utils.ExtractInt64Query(ctx.Query("limit"), models.DEFAULT_LIMIT)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"message": "Incorrect query param 'limit'",
			})
		}

		offset, err := utils.ExtractInt64Query(ctx.Query("offset"), 0)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"message": "Incorrect query param 'offset'",
			})
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

		return ctx.Status(http.StatusOK).JSON(models.PaginatedResponse[Item]{
			Count:   count,
			Limit:   limit,
			Offset:  offset,
			Cursor:  utils.GetNextOffset(offset, count, int64(len(records))),
			Results: records,
		})
	}
}

func Create(env *models.AppEnv) func(*fiber.Ctx) error {
	collection := env.DB.Collection("items")

	return func(ctx *fiber.Ctx) error {
		var record Item

		if err := ctx.BodyParser(&record); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": err.Error()})
		}

		inserted, err := collection.InsertOne(context.TODO(), record)
		if err != nil {
			panic(err)
		}

		var found Item
		collection.FindOne(context.TODO(), bson.M{"_id": inserted.InsertedID}).Decode(&found)
		return ctx.Status(http.StatusCreated).JSON(found)
	}
}

func Find(env *models.AppEnv) func(*fiber.Ctx) error {
	collection := env.DB.Collection("items")

	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Invalid id"})
		}

		var found Item
		err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&found)
		if err == mongo.ErrNoDocuments {
			return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Not found"})
		} else if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": err.Error()})
		}

		return ctx.Status(http.StatusOK).JSON(found)
	}
}
