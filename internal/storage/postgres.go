package storage

import (
	"context"
	"task-manager/internal/task"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewPostgresStorage(db *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (p *PostgresStorage) GetAll() ([]task.Task, error) {
	rows, err := p.db.Query(context.Background(),
		`SELECT id, title, done FROM tasks ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func (p *PostgresStorage) Save(tasks []task.Task) error {
	ctx := context.Background()

	tx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `TRUNCATE TABLE tasks`)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		_, err := tx.Exec(ctx,
			`INSERT INTO tasks (id, title, done) VALUES ($1, $2, $3)`,
			t.ID, t.Title, t.Done,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
