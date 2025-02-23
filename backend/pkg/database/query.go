package database

import (
	"gorm.io/gorm"
)

// QueryBuilder helps build database queries with optional conditions
type QueryBuilder struct {
	db    *gorm.DB
	query *gorm.DB
}

// NewQueryBuilder creates a new QueryBuilder instance
func NewQueryBuilder(db *gorm.DB) *QueryBuilder {
	return &QueryBuilder{
		db:    db,
		query: db,
	}
}

// WhereIf adds a where condition if the condition is true
func (qb *QueryBuilder) WhereIf(condition bool, query interface{}, args ...interface{}) *QueryBuilder {
	if condition {
		qb.query = qb.query.Where(query, args...)
	}
	return qb
}

// OrWhereIf adds an OR where condition if the condition is true
func (qb *QueryBuilder) OrWhereIf(condition bool, query interface{}, args ...interface{}) *QueryBuilder {
	if condition {
		qb.query = qb.query.Or(query, args...)
	}
	return qb
}

// PreloadIf adds a preload if the condition is true
func (qb *QueryBuilder) PreloadIf(condition bool, query string, args ...interface{}) *QueryBuilder {
	if condition {
		qb.query = qb.query.Preload(query, args...)
	}
	return qb
}

// Paginate adds pagination to the query
func (qb *QueryBuilder) Paginate(page, pageSize int) *QueryBuilder {
	offset := (page - 1) * pageSize
	qb.query = qb.query.Offset(offset).Limit(pageSize)
	return qb
}

// OrderBy adds ordering to the query
func (qb *QueryBuilder) OrderBy(value string) *QueryBuilder {
	qb.query = qb.query.Order(value)
	return qb
}

// Build returns the built query
func (qb *QueryBuilder) Build() *gorm.DB {
	return qb.query
}

// Example usage:
/*
func GetUsersByFilters(teamID *int, search string, page, pageSize int) ([]models.User, error) {
    var users []models.User
    
    query := NewQueryBuilder(GetDB()).
        WhereIf(teamID != nil, "team_id = ?", teamID).
        WhereIf(search != "", "username LIKE ?", "%"+search+"%").
        PreloadIf(true, "Team").
        Paginate(page, pageSize).
        OrderBy("created_at DESC").
        Build()
    
    if err := query.Find(&users).Error; err != nil {
        return nil, err
    }
    
    return users, nil
}
*/
