package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pagination struct {
	Page int64
	Size int64
}

func NewPagination(page, size int64) Pagination {
	return Pagination{
		Page: page,
		Size: size,
	}
}

func ExtractPagination(ctx *gin.Context) Pagination {
	sizeQuery := ctx.Query("size")
	size, err := strconv.ParseInt(sizeQuery, 10, 64)
	if err != nil || size < 0 {
		size = 10
	}

	pageQuery := ctx.Query("page")
	page, err := strconv.ParseInt(pageQuery, 10, 64)
	if err != nil || page <= 0 {
		page = 1
	}

	return NewPagination(page, size)
}

func (p Pagination) Offset() int64 {
	return (p.Page - 1) * p.Size
}
