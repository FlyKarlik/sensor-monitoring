package dao

import (
	"database/sql"
	"sensor-monitoring/internal/domain"
)

type PaginationDAO struct {
	Limit sql.NullInt64
	Page  sql.NullInt64
}

func (p *PaginationDAO) FromPaginationInput(domain *domain.PaginationInput) *PaginationDAO {
	p.Limit = toNullInt64(domain.Limit)
	p.Page = toNullInt64(domain.Page)
	return p
}

type SortDAO struct {
	IsReverse sql.NullBool
	Field     sql.NullString
}

func (s *SortDAO) FromSortInput(domain *domain.SortInput) *SortDAO {
	s.IsReverse = toNullBool(domain.IsReverse)
	s.Field = toNullString(domain.Field)
	return s
}
