package tags

import (
	"pyncz/go-rest/base"
	"pyncz/go-rest/models"
)

type Service = base.Service[Tag, TagFilters, Tag, string]

func CreateService(env *models.AppEnv) *Service {
	return base.CreateService[Tag, TagFilters, Tag, string](
		env,
		"tags",
		"slug",
	)
}
