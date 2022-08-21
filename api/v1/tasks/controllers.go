package tasks

import (
	"net/http"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	Service *Service
}

func CreateController(env *models.AppEnv) *Controller {
	return &Controller{
		Service: CreateService(env),
	}
}

// Methods

func (c *Controller) Read(ctx *fiber.Ctx) error {
	// Read filters
	filters := TaskFilters{}
	if err := ctx.QueryParser(&filters); err != nil {
		return err
	}

	records, err := c.Service.Read(&filters)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(records)
}

func (c *Controller) ReadPaginated(ctx *fiber.Ctx) error {
	// Read pagination params
	pagination := models.PaginationQuery{
		Limit: models.DEFAULT_LIMIT,
	}
	if err := ctx.QueryParser(&pagination); err != nil {
		return err
	}

	// Read filters
	filters := TaskFilters{}
	if err := ctx.QueryParser(&filters); err != nil {
		return err
	}

	records, err := c.Service.ReadPaginated(&filters, &pagination)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(records)
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	// Read form
	var form Task
	if err := ctx.BodyParser(&form); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": err.Error()})
	}
	// Validate form
	errors, _ := utils.Validate(&form)
	if errors != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&errors)
	}

	record, err := c.Service.Create(&form)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(record)
}

func (c *Controller) FindById(ctx *fiber.Ctx) error {
	// Read URI param
	id := ctx.Params("id")
	// Convert to the ID type
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Invalid id"})
	}

	record, err := c.Service.FindByKey(objectId)
	if err == mongo.ErrNoDocuments {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Not found"})
	} else if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(record)
}
