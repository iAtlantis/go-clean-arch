package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"go-clean-arch/constant"
	"go-clean-arch/domain"
)

type UserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) *UserRepository {
	return &UserRepository{Conn}
}

func (m *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			fmt.Errorf("err", err)
		}
	}()

	result = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Age,
			&t.Address,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *UserRepository) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.User{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, constant.ErrNotFound
	}

	return
}
