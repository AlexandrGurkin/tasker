package internal

import "github.com/AlexandrGurkin/tasker/models"

type TaskManager struct {
	Tasks   map[string]chan models.ResponseTask
	VMQueue map[string]chan models.RequestTask
}
