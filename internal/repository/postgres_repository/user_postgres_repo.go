package postgresrepository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhcodev/fake_store_api/internal/models"
)

type PostgresUserRepository struct {
	conn *pgxpool.Pool
}

func NewPostgresUserRepository(conn *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{conn: conn}
}

func (p *PostgresUserRepository) GetUsersByParams(ctx context.Context, params models.QueryParams) ([]models.User, error) {

	query := `SELECT
		id,
		user_type_id,
		name,
		email,
		password,
		avatar,
		phone,
		status,
		created_at,
		updated_at
	from tb_users
	LIMIT $1
	OFFSET $2`

	rows, err := p.conn.Query(ctx, query, params.Limit, params.Offset)

	if err != nil {
		return []models.User{}, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User

		err := rows.Scan(
			&u.ID,
			&u.UserTypeID,
			&u.Name,
			&u.Email,
			&u.Password,
			&u.Avatar,
			&u.Phone,
			&u.Status,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (p *PostgresUserRepository) GetUserByID(ctx context.Context, ID int) (models.User, error) {

	query := `SELECT
		id,
		user_type_id,
		name,
		email,
		password,
		avatar,
		phone,
		status,
		created_at,
		updated_at
	FROM tb_users
	WHERE id = $1`

	row := p.conn.QueryRow(ctx, query, ID)
	var user models.User

	err := row.Scan(
		&user.ID,
		&user.UserTypeID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Avatar,
		&user.Phone,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *PostgresUserRepository) UserEmailIsAvailable(ctx context.Context, email string) (bool, error) {

	query := `SELECT
		email
	FROM tb_users
	WHERE LOWER(email) = $1
	LIMIT 1`

	row := p.conn.QueryRow(ctx, query, email)
	var userEmail string

	err := row.Scan(&userEmail)

	fmt.Println("email: ", email)
	fmt.Println("userEmail: ", userEmail)
	fmt.Println("err: ", err)

	if err != nil {
		return true, nil
	}

	return false, nil
}
