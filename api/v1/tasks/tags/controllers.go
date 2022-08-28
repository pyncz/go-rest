package tags

import (
	"net/http"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"

	"github.com/gofiber/fiber/v2"
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

// Read tags.
//
// @Summary read tags
// @Description Reads a paginated list of the tags
// @Tags Tag
// @Accept json
// @Produce json
// @Param        offset    query     int  false  "Pagination offset" default(0)
// @Param        limit    query     int  false  "Pagination limit" default(12)
// @Success 200 {object} models.PaginatedListResults[tags.Tag]
// @Failure      422  {object}  utils.HttpError
// @Failure      500  {object}  utils.HttpError
// @Router /tasks/tags [get]
func (c *Controller) ReadPaginated(ctx *fiber.Ctx) error {
	// Read pagination params
	pagination := models.PaginationQuery{
		Limit: models.DEFAULT_LIMIT,
	}
	if err := ctx.QueryParser(&pagination); err != nil {
		return err
	}

	// Parse filters
	filters := TagFilters{}
	if err := ctx.QueryParser(&filters); err != nil {
		return err
	}

	records, err := c.Service.ReadPaginated(&filters, &pagination)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(records)
}

// Get tag.
//
// @Summary get tag by slug
// @Description Finds a tag by the provided slug
// @Tags Tag
// @Accept json
// @Produce json
// @Param        slug   path      string true  "Tag slug"
// @Success 200 {object} tags.Tag
// @Failure      400  {object}  utils.HttpBadRequestError
// @Failure      404  {object}  utils.HttpError
// @Failure      422  {object}  utils.HttpError
// @Failure      500  {object}  utils.HttpError
// @Router       /tasks/tags/{slug} [get]
func (c *Controller) FindBySlug(ctx *fiber.Ctx) error {
	// Read URI param
	slug := ctx.Params("slug")
	if slug == "" {
		return ctx.Status(http.StatusBadRequest).JSON(&utils.HttpBadRequestError{"slug": "Slug is not provided"})
	}

	record, err := c.Service.FindByKey(slug)
	if err == mongo.ErrNoDocuments {
		return ctx.Status(http.StatusNotFound).JSON(utils.NewError("Not found"))
	} else if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.NewError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(record)
}

// Create tag.
//
// @Summary create tag
// @Description Creates a new tag from the provided form
// @Tags Tag
// @Produce json
// @Param form body tags.TagCreateForm true "Creation form"
// @Success 201 {object} tags.Tag
// @Failure      400  {object}  utils.HttpError
// @Failure      422  {object}  utils.HttpError
// @Failure      500  {object}  utils.HttpError
// @Router       /tasks/tags [post]
func (c *Controller) Create(ctx *fiber.Ctx) error {
	// Parse form
	var form TagCreateForm
	if err := ctx.BodyParser(&form); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err.Error()))
	}
	// Validate form
	errors, _ := utils.Validate(&form)
	if errors != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&errors)
	}

	// Validate slug: try to find a record with the same slug
	// TODO: Slugify in order to make the slug unique
	_, err := c.Service.FindByKey(form.Slug)
	if err != mongo.ErrNoDocuments {
		if err != nil {
			return err
		}
		return ctx.Status(http.StatusBadRequest).JSON(&utils.HttpBadRequestError{"slug": "'slug' is not unique"})
	}

	record, err := c.Service.Create(&form)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(record)
}
