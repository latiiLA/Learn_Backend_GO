package usecases

import (
	"context"
	domain "task_management_clean_architecture/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase{
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) AddTask(c context.Context, task *domain.Task)error{
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) GetTasksByUserID(c context.Context, userID string)([]domain.Task, error){
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByUserID(ctx, userID)
}

func (tu *taskUsecase) GetAllTasks(c context.Context)([]domain.Task, error){
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchAll(ctx)
}

func (tu *taskUsecase) UpdateTask(c context.Context, taskID primitive.ObjectID, updatedData *domain.Task)(int64, int64, error){
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.UpdateByTaskID(ctx, taskID, updatedData)
}

func (tu *taskUsecase) DeleteTask(c context.Context, taskID primitive.ObjectID)(int64, error){
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.DeleteByTaskID(ctx, taskID)
}

func (tu *taskUsecase) GetSpecificTask(c context.Context, taskID primitive.ObjectID)(domain.Task, error){
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByTaskID(ctx, taskID)
}