package persistence

import (
	"context"

	. "github.com/Masterminds/squirrel"
	"github.com/beego/beego/v2/client/orm"
	"github.com/deluan/rest"
	"github.com/navidrome/navidrome/model"
)

type transcodingRepository struct {
	sqlRepository
	sqlRestful
}

func NewTranscodingRepository(ctx context.Context, o orm.QueryExecutor) model.TranscodingRepository {
	r := &transcodingRepository{}
	r.ctx = ctx
	r.ormer = o
	r.tableName = "transcoding"
	return r
}

func (r *transcodingRepository) Get(id string) (*model.Transcoding, error) {
	sel := r.newSelect().Columns("*").Where(Eq{"id": id})
	var res model.Transcoding
	err := r.queryOne(sel, &res)
	return &res, err
}

func (r *transcodingRepository) CountAll(qo ...model.QueryOptions) (int64, error) {
	return r.count(Select(), qo...)
}

func (r *transcodingRepository) FindByFormat(format string) (*model.Transcoding, error) {
	sel := r.newSelect().Columns("*").Where(Eq{"target_format": format})
	var res model.Transcoding
	err := r.queryOne(sel, &res)
	return &res, err
}

func (r *transcodingRepository) Put(t *model.Transcoding) error {
	_, err := r.put(t.ID, t)
	return err
}

func (r *transcodingRepository) Count(options ...rest.QueryOptions) (int64, error) {
	return r.count(Select(), r.parseRestOptions(options...))
}

func (r *transcodingRepository) Read(id string) (interface{}, error) {
	return r.Get(id)
}

func (r *transcodingRepository) ReadAll(options ...rest.QueryOptions) (interface{}, error) {
	sel := r.newSelect(r.parseRestOptions(options...)).Columns("*")
	res := model.Transcodings{}
	err := r.queryAll(sel, &res)
	return res, err
}

func (r *transcodingRepository) EntityName() string {
	return "transcoding"
}

func (r *transcodingRepository) NewInstance() interface{} {
	return &model.Transcoding{}
}

func (r *transcodingRepository) Save(entity interface{}) (string, error) {
	t := entity.(*model.Transcoding)
	id, err := r.put(t.ID, t)
	if err == model.ErrNotFound {
		return "", rest.ErrNotFound
	}
	return id, err
}

func (r *transcodingRepository) Update(id string, entity interface{}, cols ...string) error {
	t := entity.(*model.Transcoding)
	t.ID = id
	_, err := r.put(id, t)
	if err == model.ErrNotFound {
		return rest.ErrNotFound
	}
	return err
}

func (r *transcodingRepository) Delete(id string) error {
	err := r.delete(Eq{"id": id})
	if err == model.ErrNotFound {
		return rest.ErrNotFound
	}
	return err
}

var _ model.TranscodingRepository = (*transcodingRepository)(nil)
var _ rest.Repository = (*transcodingRepository)(nil)
var _ rest.Persistable = (*transcodingRepository)(nil)
