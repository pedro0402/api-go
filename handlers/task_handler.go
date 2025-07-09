package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pedro0402/go-mod/models"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (taskHandler *TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request) {
	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (taskHandler *TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task models.Task
	err := json.NewDecoder(request.Body).Decode(&task)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = taskHandler.DB.Exec("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)", task.Title, task.Description, task.Status)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(task)
}

func (taskHandler *TaskHandler) UpdateTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(writer, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	err = json.NewDecoder(request.Body).Decode(&task)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := taskHandler.DB.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", task.Title, task.Description, task.Status, id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(writer, "no task found with this ID", http.StatusInternalServerError)
		return
	}

	task.ID = id
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(task)
}

func (taskHandler *TaskHandler) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := taskHandler.DB.Exec("DELETE FROM tasks WHERE id=$1", id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(writer, "no task found with this ID", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
