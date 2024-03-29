// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"xapp/xdb/game/model"
)

func newXLiveProvider(db *gorm.DB, opts ...gen.DOOption) xLiveProvider {
	_xLiveProvider := xLiveProvider{}

	_xLiveProvider.xLiveProviderDo.UseDB(db, opts...)
	_xLiveProvider.xLiveProviderDo.UseModel(&model.XLiveProvider{})

	tableName := _xLiveProvider.xLiveProviderDo.TableName()
	_xLiveProvider.ALL = field.NewAsterisk(tableName)
	_xLiveProvider.SellerID = field.NewInt32(tableName, "seller_id")
	_xLiveProvider.Provider = field.NewString(tableName, "provider")
	_xLiveProvider.PushURL = field.NewString(tableName, "push_url")
	_xLiveProvider.PushKey = field.NewString(tableName, "push_key")
	_xLiveProvider.PullURL = field.NewString(tableName, "pull_url")
	_xLiveProvider.PullKey = field.NewString(tableName, "pull_key")
	_xLiveProvider.CreateTime = field.NewTime(tableName, "create_time")

	_xLiveProvider.fillFieldMap()

	return _xLiveProvider
}

type xLiveProvider struct {
	xLiveProviderDo

	ALL        field.Asterisk
	SellerID   field.Int32  // 运营商
	Provider   field.String // 提供商
	PushURL    field.String // 推流地址
	PushKey    field.String // 推流鉴权key
	PullURL    field.String // 拉流地址
	PullKey    field.String // 拉流鉴权key
	CreateTime field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (x xLiveProvider) Table(newTableName string) *xLiveProvider {
	x.xLiveProviderDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xLiveProvider) As(alias string) *xLiveProvider {
	x.xLiveProviderDo.DO = *(x.xLiveProviderDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xLiveProvider) updateTableName(table string) *xLiveProvider {
	x.ALL = field.NewAsterisk(table)
	x.SellerID = field.NewInt32(table, "seller_id")
	x.Provider = field.NewString(table, "provider")
	x.PushURL = field.NewString(table, "push_url")
	x.PushKey = field.NewString(table, "push_key")
	x.PullURL = field.NewString(table, "pull_url")
	x.PullKey = field.NewString(table, "pull_key")
	x.CreateTime = field.NewTime(table, "create_time")

	x.fillFieldMap()

	return x
}

func (x *xLiveProvider) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xLiveProvider) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 7)
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["provider"] = x.Provider
	x.fieldMap["push_url"] = x.PushURL
	x.fieldMap["push_key"] = x.PushKey
	x.fieldMap["pull_url"] = x.PullURL
	x.fieldMap["pull_key"] = x.PullKey
	x.fieldMap["create_time"] = x.CreateTime
}

func (x xLiveProvider) clone(db *gorm.DB) xLiveProvider {
	x.xLiveProviderDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xLiveProvider) replaceDB(db *gorm.DB) xLiveProvider {
	x.xLiveProviderDo.ReplaceDB(db)
	return x
}

type xLiveProviderDo struct{ gen.DO }

type IXLiveProviderDo interface {
	gen.SubQuery
	Debug() IXLiveProviderDo
	WithContext(ctx context.Context) IXLiveProviderDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXLiveProviderDo
	WriteDB() IXLiveProviderDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXLiveProviderDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXLiveProviderDo
	Not(conds ...gen.Condition) IXLiveProviderDo
	Or(conds ...gen.Condition) IXLiveProviderDo
	Select(conds ...field.Expr) IXLiveProviderDo
	Where(conds ...gen.Condition) IXLiveProviderDo
	Order(conds ...field.Expr) IXLiveProviderDo
	Distinct(cols ...field.Expr) IXLiveProviderDo
	Omit(cols ...field.Expr) IXLiveProviderDo
	Join(table schema.Tabler, on ...field.Expr) IXLiveProviderDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXLiveProviderDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXLiveProviderDo
	Group(cols ...field.Expr) IXLiveProviderDo
	Having(conds ...gen.Condition) IXLiveProviderDo
	Limit(limit int) IXLiveProviderDo
	Offset(offset int) IXLiveProviderDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXLiveProviderDo
	Unscoped() IXLiveProviderDo
	Create(values ...*model.XLiveProvider) error
	CreateInBatches(values []*model.XLiveProvider, batchSize int) error
	Save(values ...*model.XLiveProvider) error
	First() (*model.XLiveProvider, error)
	Take() (*model.XLiveProvider, error)
	Last() (*model.XLiveProvider, error)
	Find() ([]*model.XLiveProvider, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XLiveProvider, err error)
	FindInBatches(result *[]*model.XLiveProvider, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XLiveProvider) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXLiveProviderDo
	Assign(attrs ...field.AssignExpr) IXLiveProviderDo
	Joins(fields ...field.RelationField) IXLiveProviderDo
	Preload(fields ...field.RelationField) IXLiveProviderDo
	FirstOrInit() (*model.XLiveProvider, error)
	FirstOrCreate() (*model.XLiveProvider, error)
	FindByPage(offset int, limit int) (result []*model.XLiveProvider, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXLiveProviderDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xLiveProviderDo) Debug() IXLiveProviderDo {
	return x.withDO(x.DO.Debug())
}

func (x xLiveProviderDo) WithContext(ctx context.Context) IXLiveProviderDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xLiveProviderDo) ReadDB() IXLiveProviderDo {
	return x.Clauses(dbresolver.Read)
}

func (x xLiveProviderDo) WriteDB() IXLiveProviderDo {
	return x.Clauses(dbresolver.Write)
}

func (x xLiveProviderDo) Session(config *gorm.Session) IXLiveProviderDo {
	return x.withDO(x.DO.Session(config))
}

func (x xLiveProviderDo) Clauses(conds ...clause.Expression) IXLiveProviderDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xLiveProviderDo) Returning(value interface{}, columns ...string) IXLiveProviderDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xLiveProviderDo) Not(conds ...gen.Condition) IXLiveProviderDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xLiveProviderDo) Or(conds ...gen.Condition) IXLiveProviderDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xLiveProviderDo) Select(conds ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xLiveProviderDo) Where(conds ...gen.Condition) IXLiveProviderDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xLiveProviderDo) Order(conds ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xLiveProviderDo) Distinct(cols ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xLiveProviderDo) Omit(cols ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xLiveProviderDo) Join(table schema.Tabler, on ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xLiveProviderDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xLiveProviderDo) RightJoin(table schema.Tabler, on ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xLiveProviderDo) Group(cols ...field.Expr) IXLiveProviderDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xLiveProviderDo) Having(conds ...gen.Condition) IXLiveProviderDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xLiveProviderDo) Limit(limit int) IXLiveProviderDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xLiveProviderDo) Offset(offset int) IXLiveProviderDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xLiveProviderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXLiveProviderDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xLiveProviderDo) Unscoped() IXLiveProviderDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xLiveProviderDo) Create(values ...*model.XLiveProvider) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xLiveProviderDo) CreateInBatches(values []*model.XLiveProvider, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xLiveProviderDo) Save(values ...*model.XLiveProvider) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xLiveProviderDo) First() (*model.XLiveProvider, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XLiveProvider), nil
	}
}

func (x xLiveProviderDo) Take() (*model.XLiveProvider, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XLiveProvider), nil
	}
}

func (x xLiveProviderDo) Last() (*model.XLiveProvider, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XLiveProvider), nil
	}
}

func (x xLiveProviderDo) Find() ([]*model.XLiveProvider, error) {
	result, err := x.DO.Find()
	return result.([]*model.XLiveProvider), err
}

func (x xLiveProviderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XLiveProvider, err error) {
	buf := make([]*model.XLiveProvider, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xLiveProviderDo) FindInBatches(result *[]*model.XLiveProvider, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xLiveProviderDo) Attrs(attrs ...field.AssignExpr) IXLiveProviderDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xLiveProviderDo) Assign(attrs ...field.AssignExpr) IXLiveProviderDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xLiveProviderDo) Joins(fields ...field.RelationField) IXLiveProviderDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xLiveProviderDo) Preload(fields ...field.RelationField) IXLiveProviderDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xLiveProviderDo) FirstOrInit() (*model.XLiveProvider, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XLiveProvider), nil
	}
}

func (x xLiveProviderDo) FirstOrCreate() (*model.XLiveProvider, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XLiveProvider), nil
	}
}

func (x xLiveProviderDo) FindByPage(offset int, limit int) (result []*model.XLiveProvider, count int64, err error) {
	result, err = x.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = x.Offset(-1).Limit(-1).Count()
	return
}

func (x xLiveProviderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xLiveProviderDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xLiveProviderDo) Delete(models ...*model.XLiveProvider) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xLiveProviderDo) withDO(do gen.Dao) *xLiveProviderDo {
	x.DO = *do.(*gen.DO)
	return x
}
