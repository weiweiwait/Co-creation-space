package dao

import (
	"context"
	"my_project/project-project/internal/data/task"
	"my_project/project-project/internal/database/gorms"
)

type TaskStagesTemplateDao struct {
	conn *gorms.GormConn
}

func (t *TaskStagesTemplateDao) FindByProjectTemplateId(ctx context.Context, projectTemplateCode int) (list []*task.MsTaskStagesTemplate, err error) {
	session := t.conn.Session(ctx)
	err = session.
		Model(&task.MsTaskStagesTemplate{}).
		Where("project_template_code=?", projectTemplateCode).
		Order("sort desc,id asc").
		Find(&list).
		Error
	return
}

func (t *TaskStagesTemplateDao) FindInProTemIds(ctx context.Context, ids []int) ([]task.MsTaskStagesTemplate, error) {
	var tsts []task.MsTaskStagesTemplate
	session := t.conn.Session(ctx)
	err := session.Model(&task.MsTaskStagesTemplate{}).Where("project_template_code in ?", ids).Find(&tsts).Error
	return tsts, err
}

func NewTaskStagesTemplateDao() *TaskStagesTemplateDao {
	return &TaskStagesTemplateDao{
		conn: gorms.New(),
	}
}
