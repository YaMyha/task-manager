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
	// Make a query to database
	// Retrieve all tasks from the PostgreSQL database
	rows, err := p.db.Query(context.Background(),
		`SELECT id, title, done FROM tasks ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // defer closing connection after function execution

	// Iterate over the rows and build the tasks list

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	// Return the list of tasks and any error encountered during iteration
	// We don't rely only on errors from Query or Scan but also check for errors during rows iteration
	return tasks, rows.Err()
}

func (p *PostgresStorage) Save(tasks []task.Task) error {
	ctx := context.Background() // use a background context for DB operations

	tx, err := p.db.Begin(ctx) // start a transaction
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // ensure rollback if not committed

	_, err = tx.Exec(ctx, `TRUNCATE TABLE tasks`) // clear existing tasks
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

	return tx.Commit(ctx) // commit the transaction
}
