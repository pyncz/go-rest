package items

import (
	"pyncz/go-rest/base"
	"pyncz/go-rest/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service = base.Service[Item, ItemFilters, ItemCreateForm, primitive.ObjectID]

func CreateService(env *models.AppEnv) *Service {
	return base.CreateService[Item, ItemFilters, ItemCreateForm, primitive.ObjectID](
		env,
		"items",
		"_id",
	)
}
