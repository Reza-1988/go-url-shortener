package postgres

import (
	"context"

	"github.com/Reza-1988/go-url-shorten/internal/domain"
	"gorm.io/gorm"
)

// UserRepo handles user queries in PostgreSQL using GORM.
type UserRepo struct {
	db *gorm.DB // Database connection (GORM)
}

// NewUserRepo creates a new repository with a DB dependency.
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create inserts a new user row into the database.
func (r *UserRepo) Create(ctx context.Context, u *domain.User) error {
	// WithContext makes the query respect cancel/timeouts from ctx.
	return r.db.WithContext(ctx).Create(u).Error
}

// FindByEmail returns a user by email (or an error if not found).
func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User

	// Use a parameterized query to avoid SQL injection.
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// FindByID returns a user by primary key id.
func (r *UserRepo) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	var u domain.User

	// First(&u, id) queries by primary key.
	err := r.db.WithContext(ctx).First(&u, id).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}
