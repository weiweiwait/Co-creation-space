package domain

import (
	"context"
	"my_project/project-common/errs"
	"my_project/project-common/kk"
	"my_project/project-project/config"
	"my_project/project-project/internal/dao"
	"my_project/project-project/internal/repo"
	"my_project/project-project/pkg/model"
)

type TaskDomain struct {
	taskRepo repo.TaskRepo
}

func NewTaskDomain() *TaskDomain {
	return &TaskDomain{
		taskRepo: dao.NewTaskDao(),
	}
}

func (d *TaskDomain) FindProjectIdByTaskId(taskId int64) (int64, bool, *errs.BError) {
	config.SendLog(kk.Info("Find", "TaskDomain.FindProjectIdByTaskId", kk.FieldMap{
		"taskId": taskId,
	}))
	task, err := d.taskRepo.FindTaskById(context.Background(), taskId)
	if err != nil {
		config.SendLog(kk.Error(err, "TaskDomain.FindProjectIdByTaskId.taskRepo.FindTaskById", kk.FieldMap{
			"taskId": taskId,
		}))
		return 0, false, model.DBError
	}
	if task == nil {
		return 0, false, nil
	}
	return task.ProjectCode, true, nil
}
