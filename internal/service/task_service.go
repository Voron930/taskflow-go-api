package service

import (
	"context"

	"taskflow/internal/model"
	"taskflow/internal/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, userID int64, title, description string) (*model.Task, error) {
	task := &model.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      "todo",
	}

	if err := s.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil

}

func (s *TaskService) GetUserTasks(ctx context.Context, userID int64) ([]model.Task, error) {
	return s.taskRepo.GetByUserID(ctx, userID)
}

func (s *TaskService) GetTaskByID(ctx context.Context, taskID int64, userID int64) (*model.Task, error) {
	return s.taskRepo.GetByID(ctx, taskID, userID)
}

func (s *TaskService) UpdateTask(ctx context.Context, taskID, userID int64, title, description, status string) (*model.Task, error) {
	task := &model.Task{
		ID:          taskID,
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      status,
	}

	if err := s.taskRepo.Update(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, taskID, userID int64) error {
	return s.taskRepo.Delete(ctx, taskID, userID)
}
