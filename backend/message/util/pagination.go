package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page int
	Size int
}

func NewPagination(page, size int) Pagination {
	return Pagination{
		Page: page,
		Size: size,
	}
}

func ExtractPagination(ctx *gin.Context) Pagination {
	sizeQuery := ctx.Query("size")
	size, err := strconv.Atoi(sizeQuery)
	if err != nil || size <= 0 {
		size = 10
	}

	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil || page <= 0 {
		page = 1
	}

	return NewPagination(page, size)
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Size
}
