package postgres

import (
	"context"
	"fmt"
)

type Storage struct {
	client Client
}

func NewStorage(client Client) *Storage {
	return &Storage{client: client}
}

func (s *Storage) DeleteUserByID(ctx context.Context, userID int64) (bool, error) {

	query := `DELETE FROM users WHERE id = $1`
	result, err := s.client.Exec(ctx, query, userID)
	if err != nil {
		return false, err
	}

	resultRows := result.RowsAffected()
	if resultRows == 0 {
		return false, fmt.Errorf("user %d not found", userID)
	}
	
	return true, nil
}
