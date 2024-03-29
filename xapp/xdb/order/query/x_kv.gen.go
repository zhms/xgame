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

	"xapp/xdb/order/model"
)

func newXKv(db *gorm.DB, opts ...gen.DOOption) xKv {
	_xKv := xKv{}

	_xKv.xKvDo.UseDB(db, opts...)
	_xKv.xKvDo.UseModel(&model.XKv{})

	tableName := _xKv.xKvDo.TableName()
	_xKv.ALL = field.NewAsterisk(tableName)
	_xKv.SellerID = field.NewInt32(tableName, "seller_id")
	_xKv.K = field.NewString(tableName, "k")
	_xKv.V = field.NewString(tableName, "v")

	_xKv.fillFieldMap()

	return _xKv
}

type xKv struct {
	xKvDo

	ALL      field.Asterisk
	SellerID field.Int32
	K        field.String
	V        field.String

	fieldMap map[string]field.Expr
}

func (x xKv) Table(newTableName string) *xKv {
	x.xKvDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xKv) As(alias string) *xKv {
	x.xKvDo.DO = *(x.xKvDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xKv) updateTableName(table string) *xKv {
	x.ALL = field.NewAsterisk(table)
	x.SellerID = field.NewInt32(table, "seller_id")
	x.K = field.NewString(table, "k")
	x.V = field.NewString(table, "v")

	x.fillFieldMap()

	return x
}

func (x *xKv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xKv) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 3)
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["k"] = x.K
	x.fieldMap["v"] = x.V
}

func (x xKv) clone(db *gorm.DB) xKv {
	x.xKvDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xKv) replaceDB(db *gorm.DB) xKv {
	x.xKvDo.ReplaceDB(db)
	return x
}

type xKvDo struct{ gen.DO }

type IXKvDo interface {
	gen.SubQuery
	Debug() IXKvDo
	WithContext(ctx context.Context) IXKvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXKvDo
	WriteDB() IXKvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXKvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXKvDo
	Not(conds ...gen.Condition) IXKvDo
	Or(conds ...gen.Condition) IXKvDo
	Select(conds ...field.Expr) IXKvDo
	Where(conds ...gen.Condition) IXKvDo
	Order(conds ...field.Expr) IXKvDo
	Distinct(cols ...field.Expr) IXKvDo
	Omit(cols ...field.Expr) IXKvDo
	Join(table schema.Tabler, on ...field.Expr) IXKvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXKvDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXKvDo
	Group(cols ...field.Expr) IXKvDo
	Having(conds ...gen.Condition) IXKvDo
	Limit(limit int) IXKvDo
	Offset(offset int) IXKvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXKvDo
	Unscoped() IXKvDo
	Create(values ...*model.XKv) error
	CreateInBatches(values []*model.XKv, batchSize int) error
	Save(values ...*model.XKv) error
	First() (*model.XKv, error)
	Take() (*model.XKv, error)
	Last() (*model.XKv, error)
	Find() ([]*model.XKv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XKv, err error)
	FindInBatches(result *[]*model.XKv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XKv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXKvDo
	Assign(attrs ...field.AssignExpr) IXKvDo
	Joins(fields ...field.RelationField) IXKvDo
	Preload(fields ...field.RelationField) IXKvDo
	FirstOrInit() (*model.XKv, error)
	FirstOrCreate() (*model.XKv, error)
	FindByPage(offset int, limit int) (result []*model.XKv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXKvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xKvDo) Debug() IXKvDo {
	return x.withDO(x.DO.Debug())
}

func (x xKvDo) WithContext(ctx context.Context) IXKvDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xKvDo) ReadDB() IXKvDo {
	return x.Clauses(dbresolver.Read)
}

func (x xKvDo) WriteDB() IXKvDo {
	return x.Clauses(dbresolver.Write)
}

func (x xKvDo) Session(config *gorm.Session) IXKvDo {
	return x.withDO(x.DO.Session(config))
}

func (x xKvDo) Clauses(conds ...clause.Expression) IXKvDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xKvDo) Returning(value interface{}, columns ...string) IXKvDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xKvDo) Not(conds ...gen.Condition) IXKvDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xKvDo) Or(conds ...gen.Condition) IXKvDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xKvDo) Select(conds ...field.Expr) IXKvDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xKvDo) Where(conds ...gen.Condition) IXKvDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xKvDo) Order(conds ...field.Expr) IXKvDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xKvDo) Distinct(cols ...field.Expr) IXKvDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xKvDo) Omit(cols ...field.Expr) IXKvDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xKvDo) Join(table schema.Tabler, on ...field.Expr) IXKvDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xKvDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXKvDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xKvDo) RightJoin(table schema.Tabler, on ...field.Expr) IXKvDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xKvDo) Group(cols ...field.Expr) IXKvDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xKvDo) Having(conds ...gen.Condition) IXKvDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xKvDo) Limit(limit int) IXKvDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xKvDo) Offset(offset int) IXKvDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xKvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXKvDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xKvDo) Unscoped() IXKvDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xKvDo) Create(values ...*model.XKv) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xKvDo) CreateInBatches(values []*model.XKv, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xKvDo) Save(values ...*model.XKv) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xKvDo) First() (*model.XKv, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XKv), nil
	}
}

func (x xKvDo) Take() (*model.XKv, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XKv), nil
	}
}

func (x xKvDo) Last() (*model.XKv, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XKv), nil
	}
}

func (x xKvDo) Find() ([]*model.XKv, error) {
	result, err := x.DO.Find()
	return result.([]*model.XKv), err
}

func (x xKvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XKv, err error) {
	buf := make([]*model.XKv, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xKvDo) FindInBatches(result *[]*model.XKv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xKvDo) Attrs(attrs ...field.AssignExpr) IXKvDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xKvDo) Assign(attrs ...field.AssignExpr) IXKvDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xKvDo) Joins(fields ...field.RelationField) IXKvDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xKvDo) Preload(fields ...field.RelationField) IXKvDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xKvDo) FirstOrInit() (*model.XKv, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XKv), nil
	}
}

func (x xKvDo) FirstOrCreate() (*model.XKv, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XKv), nil
	}
}

func (x xKvDo) FindByPage(offset int, limit int) (result []*model.XKv, count int64, err error) {
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

func (x xKvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xKvDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xKvDo) Delete(models ...*model.XKv) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xKvDo) withDO(do gen.Dao) *xKvDo {
	x.DO = *do.(*gen.DO)
	return x
}
