package helper

import "gorm.io/gorm"

type PaginationParam struct {
	Page   int `validate:"min=1"`
	Limit  int `validate:"min=1"`
	Search string
}

type MultilingualPaginationParam struct {
	Page     int `validate:"min=1"`
	Limit    int `validate:"min=1"`
	SearchID string
	SearchEN string
}

func Paginate(limit int, page int, query *gorm.DB) (*gorm.DB, int64, error) {
	var count int64

	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	query = query.Offset(offset).Limit(limit)

	return query, count, nil
}
