package repo

import (
	"context"
	"my_project/project-project/internal/data/task"
	"my_project/project-project/internal/database"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, ids []int) ([]task.MsTaskStagesTemplate, error)
	FindByProjectTemplateId(ctx context.Context, projectTemplateCode int) (list []*task.MsTaskStagesTemplate, err error)
}
type TaskStagesRepo interface {
	SaveTaskStages(ctx context.Context, conn database.DbConn, ts *task.TaskStages) error
	FindStagesByProjectId(ctx context.Context, projectCode int64, page int64, pageSize int64) (list []*task.TaskStages, total int64, err error)
	FindById(ctx context.Context, id int) (ts *task.TaskStages, err error)
}

type TaskRepo interface {
	FindTaskByStageCode(ctx context.Context, stageCode int) (list []*task.Task, err error)
	FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *task.TaskMember, err error)
	FindTaskMaxIdNum(ctx context.Context, projectCode int64) (v *int, err error)
	FindTaskSort(ctx context.Context, projectCode int64, stageCode int64) (v *int, err error)
	SaveTask(ctx context.Context, conn database.DbConn, ts *task.Task) error
	SaveTaskMember(ctx context.Context, conn database.DbConn, tm *task.TaskMember) error
	FindTaskById(ctx context.Context, taskCode int64) (ts *task.Task, err error)
	UpdateTaskSort(ctx context.Context, conn database.DbConn, ts *task.Task) error
	FindTaskByStageCodeLtSort(ctx context.Context, stageCode int, sort int) (ts *task.Task, err error)
	FindTaskByAssignTo(ctx context.Context, memberId int64, done int, page int64, size int64) ([]*task.Task, int64, error)
	FindTaskByMemberCode(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*task.Task, total int64, err error)
	FindTaskByCreateBy(ctx context.Context, memberId int64, done int, page int64, size int64) (tList []*task.Task, total int64, err error)
}
