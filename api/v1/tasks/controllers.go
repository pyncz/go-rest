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

// Read tasks.
//
// @Summary read tasks
// @Description Reads a paginated list of the tasks
// @Tags Task
// @Accept json
// @Produce json
// @Param        offset    query     int  false  "Pagination offset" default(0)
// @Param        limit    query     int  false  "Pagination limit" default(12)
// @Success 200 {object} models.PaginatedListResults[tasks.Task]
// @Failure      422  {object}  utils.HttpError
// @Failure      500  {object}  utils.HttpError
// @Router /tasks [get]
func (c *Controller) ReadPaginated(ctx *fiber.Ctx) error {
	// Read pagination params
	pagination := models.PaginationQuery{
		Limit: models.DEFAULT_LIMIT,
	}
	if err := ctx.QueryParser(&pagination); err != nil {
		return err
	}

	// Parse filters
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

// Get task.
//
// @Summary get task by ID
// @Description Finds a task by the provided ID
// @Tags Task
// @Accept json
// @Produce json
// @Param        id   path      string true  "Task ID"
// @Success 200 {object} tasks.Task
// @Failure      400  {object}  utils.HttpBadRequestError
// @Failure      404  {object}  utils.HttpError
// @Failure      422  {object}  utils.HttpError
// @Failure      500  {object}  utils.HttpError
// @Router       /tasks/{id} [get]
func (c *Controller) FindById(ctx *fiber.Ctx) error {
	// Read URI param
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(&utils.HttpBadRequestError{"id": "Id is not provided"})
	}
	// Convert to the ID type
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&utils.HttpBadRequestError{"id": "Invalid id"})
	}

	record, err := c.Service.FindByKey(objectId)
	if err == mongo.ErrNoDocuments {
		return ctx.Status(http.StatusNotFound).JSON(utils.NewError("Not found"))
	} else if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(record)
}

// Create task.
//
// @Summary create task
// @Description Creates a new task from the provided form
// @Tags Task
// @Produce json
// @Param form body tasks.TaskCreateForm true "Creation form"
// @Success 201 {object} tasks.Task
// @Failure      400  {object}  utils.HttpError
// @Failure      422  {object}  utils.HttpError
// @Failure      500  {object}  utils.HttpError
// @Router       /tasks [post]
func (c *Controller) Create(ctx *fiber.Ctx) error {
	// Parse form
	var form TaskCreateForm
	if err := ctx.BodyParser(&form); err != nil {
		c.Service.Env.Log.Printf("%#v", form)
		c.Service.Env.Log.Println("sdafd", err, err.Error())
		return ctx.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err.Error()))
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
