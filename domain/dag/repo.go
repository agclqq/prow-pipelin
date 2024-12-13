package flow

import (
	"github.com/agclqq/prow-framework/db/repo"
	"github.com/agclqq/prow-pipeline/boot"
	"gorm.io/gorm"
)

type RepoConfDag struct {
	db *gorm.DB
}

func NewConfFlowDag(db *gorm.DB) *RepoConfDag {
	cfd := &RepoConfDag{}
	if db != nil {
		cfd.db = db
	} else {
		cfd.db = boot.GetDbW()
	}
	return cfd
}

func (r *RepoConfDag) Table() string {
	return "repo_conf"
}

func (r *RepoConfDag) Select(columns string, where any, group string, having any, order string, page, pageSize int) ([]ConfDag, error) {
	var data []ConfDag
	tx := r.db.Table(r.Table())
	rs := repo.Select(tx, columns, where, group, having, order, page, pageSize).Find(&data)
	return data, rs.Error
}
