package api

import "github.com/guregu/null/v5"

type PaginationQuery struct {
	Page    null.Int `query:"page"`
	PerPage null.Int `query:"per_page"`
}
