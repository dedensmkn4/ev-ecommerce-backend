package sqkit

import (
	sq "github.com/Masterminds/squirrel"
)

type (
	// Eq equal
	Eq map[string]interface{}
)

var _ SelectOption = (Eq)(nil)

// CompileSelect to compile select query for filtering
func (e Eq) CompileSelect(base sq.SelectBuilder) sq.SelectBuilder {
	if len(e) > 0 {
		return base.Where(sq.Eq(e))
	}
	return base
}
