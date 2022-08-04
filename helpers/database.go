package helpers

import (
	"fmt"

	"crm-sebagian-team/domain"

	"gorm.io/gorm"
)

func PaginateQuery(params *domain.Request) func(db *gorm.DB) *gorm.DB {
	defaultSort := "id"

	offset := (params.Page - 1) * params.PerPage
	if params.SortBy == "" {
		fmt.Println("empty sort by")
		params.SortBy = defaultSort
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(offset)).Limit(int(params.PerPage)).Order(params.SortBy + " " + params.SortOrder)
	}
}
