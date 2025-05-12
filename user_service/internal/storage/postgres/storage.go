package postgres

import (
	"context"
	"github.com/Nikita-Mihailuk/simple_microservices_example/user_service/internal/domain/model"
)

type Storage struct {
	client Client
}

func NewStorage(client Client) *Storage {
	return &Storage{client: client}
}

func (s *Storage) SaveUser(ctx context.Context, name string, age int32) (int64, error) {
	var userId int64
	query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"

	err := s.client.QueryRow(ctx, query, name, age).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *Storage) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User

	rows, err := s.client.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
