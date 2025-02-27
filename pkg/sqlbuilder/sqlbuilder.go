package sqlbuilder

import "strings"

type SQLBuilder struct {
	sb   *strings.Builder
	args []any
}

func New() *SQLBuilder {
	sqlBuilder := new(SQLBuilder)
	sqlBuilder.sb = new(strings.Builder)
	sqlBuilder.args = []any{}

	return sqlBuilder
}

func (sb *SQLBuilder) S(query string, args ...any) *SQLBuilder {
	sb.sb.WriteString(" ")
	sb.sb.WriteString(query)
	sb.args = append(sb.args, args...)

	return sb
}

func (sb *SQLBuilder) Build() (string, []any) {
	return sb.sb.String(), sb.args
}

func (sb *SQLBuilder) Reset() {
	sb.sb.Reset()
	sb.args = []any{}
}
