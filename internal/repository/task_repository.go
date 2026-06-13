package repository

import (
	"context"
	"time"

	"taskflow/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
	INSERT INTO tasks (
	user_id,
	title,
	description,
	status,
	created_at,
	updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, updated_at
	`

	now := time.Now()

	err := r.db.QueryRow(
		ctx,
		query,
		task.UserID,
		task.Title,
		task.Description,
		task.Status,
		now,
		now,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return err
}

func (r *TaskRepository) GetByUserID(ctx context.Context, userID int64) ([]model.Task, error) {
	query := `
	SELECT id, user_id, title, description, status, created_at, updated_at
	FROM tasks
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []model.Task{}

	for rows.Next() {
		var task model.Task

		if err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, taskID, userID int64) (*model.Task, error) {
	query := `
	SELECT id, user_id, title, description, status, created_at, updated_at
	FROM tasks
	WHERE id = $1 AND user_id = $2
	`

	var task model.Task

	err := r.db.QueryRow(ctx, query, taskID, userID).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *model.Task) error {
	query := `
	UPDATE tasks
	SET title = $1,
	 description = $2, 
	 status = $3, 
	 updated_at = $4
	WHERE id = $5 AND user_id = $6
	RETURNING id, user_id, title, description, status, created_at, updated_at
	`

	now := time.Now()

	err := r.db.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Status,
		now,
		task.ID,
		task.UserID,
	).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return err
}

func (r *TaskRepository) Delete(ctx context.Context, taskID, userID int64) error {
	query := `
	DELETE FROM tasks
	WHERE id = $1 AND user_id = $2
	`
	_, err := r.db.Exec(ctx, query, taskID, userID)
	return err
}
