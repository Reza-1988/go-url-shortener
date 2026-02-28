package postgres

import (
	"context"

	"github.com/Reza-1988/go-url-shorten/internal/domain"
	"gorm.io/gorm"
)

// URLRepo handles URL queries in PostgreSQL using GORM.
type URLRepo struct {
	db *gorm.DB // Database connection (GORM)
}

// NewURLRepo creates a new URL repository.
func NewURLRepo(db *gorm.DB) *URLRepo {
	return &URLRepo{db: db}
}

// Create inserts a new URL row into the database.
func (r *URLRepo) Create(ctx context.Context, u *domain.URL) error {
	// WithContext makes the query respect cancel/timeouts from ctx.
	return r.db.WithContext(ctx).Create(u).Error
}

// FindByShortCode returns the URL row for a given short code.
func (r *URLRepo) FindByShortCode(ctx context.Context, shortCode string) (*domain.URL, error) {
	var u domain.URL

	// Use a parameterized query to avoid SQL injection.
	err := r.db.WithContext(ctx).Where("short_code = ?", shortCode).First(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ListByOwner returns URLs for one user (newest first), with optional paging.
func (r *URLRepo) ListByOwner(ctx context.Context, ownerID int64, limit, offset int) ([]domain.URL, error) {
	var items []domain.URL

	// Build the base query (filter by owner and order by newest).
	q := r.db.WithContext(ctx).
		Where("owner_id = ?", ownerID).
		Order("id DESC")

	// Apply limit/offset only when provided (> 0).
	if limit > 0 {
		q = q.Limit(limit)
	}
	if offset > 0 {
		q = q.Offset(offset)
	}

	err := q.Find(&items).Error
	return items, err
}

// IncrementClickAtomic increases click_count by 1 in a single DB statement.
// "Atomic" means it is safe even if many requests run at the same time.
func (r *URLRepo) IncrementClickAtomic(ctx context.Context, shortCode string) (int64, error) {
	// UpdateColumn avoids overwriting other fields and uses SQL: click_count = click_count + 1
	res := r.db.WithContext(ctx).
		Model(&domain.URL{}).
		Where("short_code = ? AND is_disabled = false", shortCode).
		UpdateColumn("click_count", gorm.Expr("click_count + 1"))

	// RowsAffected == 0 usually means: not found, or the URL is disabled.
	return res.RowsAffected, res.Error
}
