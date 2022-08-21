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

func collection(env *models.AppEnv) *mongo.Collection {
	return env.DB.Collection("items")
}

// Controllers
func Read(env *models.AppEnv) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		pagination := models.PaginationQuery{
			Limit: models.DEFAULT_LIMIT,
		}

		if err := ctx.QueryParser(&pagination); err != nil {
			return err
		}

		var records []Item

		filter := bson.D{}
		count, err := collection(env).CountDocuments(context.TODO(), &filter)
		if err != nil {
			return err
		}
		opts := options.Find().SetLimit(pagination.Limit).SetSkip(pagination.Offset)
		cursor, err := collection(env).Find(context.TODO(), &filter, opts)

		if err != nil {
			return err
		}
		if err = cursor.All(context.TODO(), &records); err != nil {
			return err
		}
		defer cursor.Close(context.TODO())

		return ctx.Status(http.StatusOK).JSON(&models.PaginatedListResults[Item]{
			Count:   count,
			Limit:   pagination.Limit,
			Offset:  pagination.Offset,
			Cursor:  utils.GetNextOffset(pagination.Offset, count, int64(len(records))),
			Results: records,
		})
	}
}

func Create(env *models.AppEnv) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var record Item

		if err := ctx.BodyParser(&record); err != nil {
			return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": err.Error()})
		}

		errors := utils.Validate(&record)
		if errors != nil {
			return ctx.Status(http.StatusBadRequest).JSON(&errors)
		}

		inserted, err := collection(env).InsertOne(context.TODO(), &record)
		if err != nil {
			return err
		}

		var found Item
		collection(env).FindOne(context.TODO(), bson.M{"_id": inserted.InsertedID}).Decode(&found)
		return ctx.Status(http.StatusCreated).JSON(&found)
	}
}

func Find(env *models.AppEnv) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Invalid id"})
		}

		var found Item
		err = collection(env).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&found)
		if err == mongo.ErrNoDocuments {
			return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Not found"})
		} else if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": err.Error()})
		}

		return ctx.Status(http.StatusOK).JSON(&found)
	}
}
