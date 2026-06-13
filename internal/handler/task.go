package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"taskflow/internal/middleware"
	"taskflow/internal/model"
	"taskflow/internal/service"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	task, err := h.taskService.CreateTask(r.Context(), userID, req.Title, req.Description)
	if err != nil {
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(task)
}

type GetTasksResponse struct {
	Tasks []model.Task `json:"tasks"`
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	tasks, err := h.taskService.GetUserTasks(r.Context(), userID)
	if err != nil {
		http.Error(w, "failed to get tasks", http.StatusInternalServerError)
		return
	}

	resp := GetTasksResponse{
		Tasks: tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	taskID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	task, err := h.taskService.GetTaskByID(r.Context(), taskID, userID)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(task)
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	taskID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	var req UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	task, err := h.taskService.UpdateTask(
		r.Context(),
		taskID,
		userID,
		req.Title,
		req.Description,
		req.Status,
	)
	if err != nil {
		http.Error(w, "failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	taskID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	if err := h.taskService.DeleteTask(r.Context(), taskID, userID); err != nil {
		http.Error(w, "failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
