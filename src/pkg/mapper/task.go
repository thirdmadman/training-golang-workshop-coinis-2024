package mapper

import "todo-service/pkg/model"

func MapToDto(task model.Task) model.MongoTask {
	return model.MongoTask{
		Title:     task.Title,
		Completed: task.Completed,
	}
}

func MapToModel(task model.MongoTask) model.Task {
	return model.Task{
		ID:        task.ID.Hex(),
		Title:     task.Title,
		Completed: task.Completed,
	}
}
