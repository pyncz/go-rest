package contexts

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

type Service[
	T any,
	TFilters map[string]any,
	TCreateForm any,
] struct {
	CollectionName string
	Env            *models.AppEnv
	Paginated      bool
	KeyParam       string
	Lookup         string
}

func NewService[
	T any,
	TFilters map[string]any,
	TCreateForm any,
](
	env *models.AppEnv,
	collectionName string,
	paginated bool,
	keyParam string,
	lookup string,
) *Service[T, TFilters, TCreateForm] {
	// Set default values
	if keyParam == "" {
		keyParam = "id"
	}
	if lookup == "" {
		lookup = "_id"
	}

	return &Service[T, TFilters, TCreateForm]{
		CollectionName: collectionName,
		Env:            env,
		Paginated:      paginated,
		KeyParam:       keyParam,
		Lookup:         lookup,
	}
}

func (s *Service[_, _, _]) collection() *mongo.Collection {
	return s.Env.DB.Collection(s.CollectionName)
}

// Controllers
func (s *Service[T, TFilters, _]) Read(ctx *fiber.Ctx) error {
	c := s.collection()

	var records []T

	// Parse filters
	filters := *new(TFilters)
	if err := ctx.QueryParser(&filters); err != nil {
		return err
	}

	// Get count
	count, err := c.CountDocuments(context.TODO(), &filters)
	if err != nil {
		return err
	}

	// Build options
	opts := options.Find()
	var pagination *models.PaginationQuery
	if s.Paginated {
		pagination = &models.PaginationQuery{
			Limit: models.DEFAULT_LIMIT,
		}

		if err := ctx.QueryParser(pagination); err != nil {
			return err
		}

		opts = opts.SetLimit(pagination.Limit).SetSkip(pagination.Offset)
	}

	// Extract results
	cursor, err := c.Find(context.TODO(), &filters, opts)
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &records); err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	// Form results
	if pagination != nil {
		return ctx.Status(http.StatusOK).JSON(&models.PaginatedListResults[T]{
			Count:   count,
			Limit:   pagination.Limit,
			Offset:  pagination.Offset,
			Cursor:  utils.GetNextOffset(pagination.Offset, count, int64(len(records))),
			Results: records,
		})
	} else {
		return ctx.Status(http.StatusOK).JSON(&models.ListResults[T]{
			Count:   count,
			Results: records,
		})
	}
}

func (s *Service[T, _, TCreateForm]) Create(ctx *fiber.Ctx) error {
	c := s.collection()

	var record T

	if err := ctx.BodyParser(&record); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": err.Error()})
	}

	errors, _ := utils.Validate(&record)
	if errors != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&errors)
	}

	inserted, err := c.InsertOne(context.TODO(), &record)
	if err != nil {
		return err
	}

	var found T
	c.FindOne(context.TODO(), bson.M{"_id": inserted.InsertedID}).Decode(&found)
	return ctx.Status(http.StatusCreated).JSON(&found)
}

func (s *Service[T, TFilters, _]) Find(ctx *fiber.Ctx) error {
	c := s.collection()

	// Parse filters
	filters := *new(TFilters)
	if err := ctx.QueryParser(&filters); err != nil {
		return err
	}

	id := ctx.Params(s.KeyParam)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Invalid id"})
	}

	filters[s.Lookup] = objectId

	var found T
	err = c.FindOne(context.TODO(), filters).Decode(&found)
	if err == mongo.ErrNoDocuments {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Not found"})
	} else if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(&found)
}
