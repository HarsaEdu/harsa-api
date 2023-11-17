package conversion

import "github.com/HarsaEdu/harsa-api/internal/model/web"

func RecordToPaginationResponse(offset, limit int, total int64) *web.Pagination {
	return &web.Pagination{
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
}
