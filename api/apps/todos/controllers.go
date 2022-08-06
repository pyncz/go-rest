package todos

import (
	"net/http"
	"pyncz/go-rest/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos = []Todo{
	{ID: 1, Title: "Yo wtf", Completed: true},
}

func Read(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, todos)
}

func Create(ctx *gin.Context) {
	var record Todo

	if err := ctx.BindJSON(&record); err != nil {
		return
	}

	for _, tag := range todos {
		if tag.ID == record.ID {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "'id' is not unique"})
		}
	}

	todos = append(todos, record)
	ctx.JSON(http.StatusCreated, record)
}

func Find(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect path param 'id'"})
		return
	}

	record, err := utils.FindByField(
		todos,
		func(t Todo) int { return t.ID },
		id,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	ctx.JSON(http.StatusOK, record)
}
