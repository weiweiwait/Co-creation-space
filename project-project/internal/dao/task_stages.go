package dao

import (
	"context"
	"my_project/project-project/internal/data/task"
	"my_project/project-project/internal/database"
	"my_project/project-project/internal/database/gorms"
)

type TaskStagesDao struct {
	conn *gorms.GormConn
}

func (t *TaskStagesDao) FindById(ctx context.Context, id int) (ts *task.TaskStages, err error) {
	err = t.conn.Session(ctx).Where("id=?", id).Find(&ts).Error
	return
}

func (t *TaskStagesDao) FindStagesByProjectId(
	ctx context.Context,
	projectCode int64,
	page int64,
	pageSize int64) (list []*task.TaskStages, total int64, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&task.TaskStages{}).
		Where("project_code=?", projectCode).
		Order("sort asc").
		Limit(int(pageSize)).Offset(int((page - 1) * pageSize)).
		Find(&list).Error
	err = session.Model(&task.TaskStages{}).
		Where("project_code=?", projectCode).
		Count(&total).Error
	return
}

func (t *TaskStagesDao) SaveTaskStages(ctx context.Context, conn database.DbConn, ts *task.TaskStages) error {
	t.conn = conn.(*gorms.GormConn)
	err := t.conn.Tx(ctx).Save(&ts).Error
	return err
}

func NewTaskStagesDao() *TaskStagesDao {
	return &TaskStagesDao{
		conn: gorms.New(),
	}
}
