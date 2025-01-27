package db

import (
	"context"
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func (pg *postgres) GetUsers(ctx context.Context) ([]User, error) {
	var err error

	query := `SELECT id, name, email FROM users LIMIT 10`

	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
