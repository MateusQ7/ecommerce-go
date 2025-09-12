package db

import (
	"database/sql"
	"time"

	"github.com/MateusQ7/ecommerce-go/auth-service/internal/domain"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	query := `
	INSERT INTO users (name, email, password, role, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at
	`

	return r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		time.Now(),
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
	query := `SELECT id, name, email, password, role, created_at FROM users`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var u domain.User

		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.Role,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
