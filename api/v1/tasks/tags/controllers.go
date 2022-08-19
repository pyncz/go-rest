package tags

import (
	"context"
	"net/http"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func collection(env *models.AppEnv) *mongo.Collection {
	return env.DB.Collection("tags")
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

		var records []Tag

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

		return ctx.Status(http.StatusOK).JSON(&models.PaginatedListResults[Tag]{
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
		var record Tag

		if err := ctx.BodyParser(&record); err != nil {
			return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": err.Error()})
		}

		// Validate slug
		var matched Tag
		err := collection(env).FindOne(context.TODO(), bson.M{"slug": record.Slug}).Decode(&matched)
		if err != mongo.ErrNoDocuments {
			if err != nil {
				return err
			}
			return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "'slug' is not unique"})
		}

		inserted, err := collection(env).InsertOne(context.TODO(), &record)
		if err != nil {
			return err
		}

		var found Tag
		collection(env).FindOne(context.TODO(), bson.M{"_id": inserted.InsertedID}).Decode(&found)
		return ctx.Status(http.StatusCreated).JSON(&found)
	}
}

func Find(env *models.AppEnv) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		slug := ctx.Params("slug")

		var found Tag
		err := collection(env).FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&found)
		if err == mongo.ErrNoDocuments {
			return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Not found"})
		} else if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": err.Error()})
		}

		return ctx.Status(http.StatusOK).JSON(&found)
	}
}
