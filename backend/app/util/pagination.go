package util

import (
	common "backend/app/common/models"
	"math"
)

func CreateMeta(itemCount, currentPage int, pageSize int, totalItemCount int) common.Meta {
	var meta common.Meta

	itemCounts := pageSize
	if currentPage*pageSize > totalItemCount {
		itemCounts = totalItemCount % pageSize
		if itemCounts == 0 && totalItemCount > 0 {
			itemCounts = pageSize
		}
	}

	previousPage := currentPage - 1
	if previousPage < 1 {
		previousPage = 0
	}

	totalPages := int(math.Ceil(float64(totalItemCount) / float64(pageSize)))
	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = 0
	}

	if currentPage > 0 && pageSize > 1 {
		meta = common.Meta{
			ItemCount: itemCounts,
			ItemTotal: totalItemCount,
			Page: &common.Page{
				IsCursor: false,
				Current:  currentPage,
				Previous: previousPage,
				Next:     nextPage,
				Limit:    pageSize,
				Total:    totalPages,
			},
		}
	} else {
		meta = common.Meta{
			ItemCount: itemCount,
		}
	}
	return meta
}
