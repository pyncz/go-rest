package tags

import (
	"net/http"
	"pyncz/go-rest/utils"

	"github.com/gin-gonic/gin"
)

var tags = []Tag{
	{Name: "Taggg", Slug: "taggg"},
}

func Read(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tags)
}

func Create(ctx *gin.Context) {
	var record Tag

	if err := ctx.BindJSON(&record); err != nil {
		return
	}

	for _, tag := range tags {
		if tag.Slug == record.Slug {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "'slug' is not unique"})
		}
	}

	tags = append(tags, record)
	ctx.JSON(http.StatusCreated, record)
}

func Find(ctx *gin.Context) {
	slug := ctx.Param("slug")
	record, err := utils.FindByField(
		tags,
		func(t Tag) string { return t.Slug },
		slug,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	ctx.JSON(http.StatusOK, record)
}
