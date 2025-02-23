package database

import (
	"math"

	"gorm.io/gorm"
)

// Pagination represents pagination information
type Pagination struct {
	CurrentPage  int         `json:"currentPage"`
	PageSize     int         `json:"pageSize"`
	TotalItems   int64      `json:"totalItems"`
	TotalPages   int         `json:"totalPages"`
	HasPrevious  bool        `json:"hasPrevious"`
	HasNext      bool        `json:"hasNext"`
	Items        interface{} `json:"items"`
}

// NewPagination creates a new Pagination instance
func NewPagination(page, pageSize int) *Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	return &Pagination{
		CurrentPage: page,
		PageSize:    pageSize,
	}
}

// Paginate executes a paginated query
func (p *Pagination) Paginate(db *gorm.DB, result interface{}) error {
	var totalItems int64
	if err := db.Model(result).Count(&totalItems).Error; err != nil {
		return err
	}

	p.TotalItems = totalItems
	p.TotalPages = int(math.Ceil(float64(totalItems) / float64(p.PageSize)))
	p.HasPrevious = p.CurrentPage > 1
	p.HasNext = p.CurrentPage < p.TotalPages

	offset := (p.CurrentPage - 1) * p.PageSize
	if err := db.Offset(offset).Limit(p.PageSize).Find(result).Error; err != nil {
		return err
	}

	p.Items = result
	return nil
}

// Example usage:
/*
func GetPaginatedTasks(page, pageSize int, projectID int) (*Pagination, error) {
    var tasks []models.Task
    
    pagination := NewPagination(page, pageSize)
    
    query := GetDB().
        Where("project_id = ?", projectID).
        Preload("Author").
        Preload("Assignee").
        Order("created_at DESC")
    
    if err := pagination.Paginate(query, &tasks); err != nil {
        return nil, err
    }
    
    return pagination, nil
}
*/
