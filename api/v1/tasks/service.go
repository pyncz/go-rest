package tasks

import (
	"pyncz/go-rest/base"
	"pyncz/go-rest/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service = base.Service[Task, TaskFilters, TaskCreateForm, primitive.ObjectID]

func CreateService(env *models.AppEnv) *Service {
	return base.CreateService[Task, TaskFilters, TaskCreateForm, primitive.ObjectID](
		env,
		"tasks",
		"_id",
	)
}
