package tags

import (
	"pyncz/go-rest/base"
	"pyncz/go-rest/models"
)

type Service = base.Service[Tag, TagFilters, TagCreateForm, string]

func CreateService(env *models.AppEnv) *Service {
	return base.CreateService[Tag, TagFilters, TagCreateForm, string](
		env,
		"tags",
		"slug",
	)
}
