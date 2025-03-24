package repository

import "github.com/guregu/null/v6"

type FindAllOptions struct {
	Limit  null.Int
	Offset null.Int
	Where  map[string]any
}
