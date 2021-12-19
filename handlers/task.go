package handlers

import (
	"encoding/json"
	"net/http"
	"packages/auth"
	"packages/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTasks(rw http.ResponseWriter, r *http.Request) {
	if token := auth.ValidateToken(r); token.Valid {
		idUser := int(token.Claims.(*models.Claim).User.Id)

		if tasks, err := models.GetTasks(idUser); err != nil {
			models.SendNotFound(rw)
		} else {
			models.SendData(rw, tasks)
		}
	} else {
		models.SendUnauthorized(rw)
	}
}

func GetTask(rw http.ResponseWriter, r *http.Request) {
	if token := auth.ValidateToken(r); token.Valid {
		if task, err := getTaskById(r); err != nil {
			models.SendNotFound(rw)
		} else {
			models.SendData(rw, task)
		}
	} else {
		models.SendUnauthorized(rw)
	}
}

func NewTask(rw http.ResponseWriter, r *http.Request) {
	if token := auth.ValidateToken(r); token.Valid {
		task := models.Task{}
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			models.SendUnprocessableEntity(rw)
		} else {
			task.Save()
			models.SendData(rw, task)
		}
	} else {
		models.SendUnauthorized(rw)
	}
}

func UpdateTask(rw http.ResponseWriter, r *http.Request) {
	var idTask int64

	if token := auth.ValidateToken(r); token.Valid {
		taskById, err := getTaskById(r)
		if err != nil {
			models.SendNotFound(rw)
		} else {
			idTask = taskById.Id
		}

		task := models.Task{}
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			models.SendUnprocessableEntity(rw)
		} else {
			task.Id = idTask
			task.Completed = taskById.Completed
			task.CreatedAt = taskById.CreatedAt
			task.IdUser = taskById.IdUser
			task.Save()
			models.SendData(rw, task)
		}
	} else {
		models.SendUnauthorized(rw)
	}
}

func DeleteTask(rw http.ResponseWriter, r *http.Request) {
	if token := auth.ValidateToken(r); token.Valid {
		if task, err := getTaskById(r); err != nil {
			models.SendNotFound(rw)
		} else {
			task.Delete()
			models.SendData(rw, task)
		}
	} else {
		models.SendUnauthorized(rw)
	}
}

func getTaskById(r *http.Request) (*models.Task, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if task, err := models.GetTask(id); err != nil {
		return nil, err
	} else {
		return task, nil
	}
}
