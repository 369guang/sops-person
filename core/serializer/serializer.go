package serializer

import (
	"math"
	"person/core/errno"
)

// Partition 分页器，return (起始值, 分页大小, 错误)
func Partition(total int64, page, pageSize int) (int, int, error) {
	if total == 0 {
		return 0, 0, nil
	}

	if pageSize == 0 {
		pageSize = 20
	}

	pageTotal := int(math.Ceil(float64(total) / float64(pageSize)))
	if page <= 0 || pageSize <= 0 {
		return 0, 0, errno.ErrMinPage
	}
	if pageTotal < page {
		return 0, 0, errno.ErrMaxPage
	}
	start := (page - 1) * pageSize

	return start, pageSize, nil
}
