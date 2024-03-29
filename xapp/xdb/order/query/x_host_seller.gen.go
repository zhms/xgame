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

func newXHostSeller(db *gorm.DB, opts ...gen.DOOption) xHostSeller {
	_xHostSeller := xHostSeller{}

	_xHostSeller.xHostSellerDo.UseDB(db, opts...)
	_xHostSeller.xHostSellerDo.UseModel(&model.XHostSeller{})

	tableName := _xHostSeller.xHostSellerDo.TableName()
	_xHostSeller.ALL = field.NewAsterisk(tableName)
	_xHostSeller.Host = field.NewString(tableName, "host")
	_xHostSeller.SellerID = field.NewInt32(tableName, "seller_id")

	_xHostSeller.fillFieldMap()

	return _xHostSeller
}

type xHostSeller struct {
	xHostSellerDo

	ALL      field.Asterisk
	Host     field.String
	SellerID field.Int32

	fieldMap map[string]field.Expr
}

func (x xHostSeller) Table(newTableName string) *xHostSeller {
	x.xHostSellerDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xHostSeller) As(alias string) *xHostSeller {
	x.xHostSellerDo.DO = *(x.xHostSellerDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xHostSeller) updateTableName(table string) *xHostSeller {
	x.ALL = field.NewAsterisk(table)
	x.Host = field.NewString(table, "host")
	x.SellerID = field.NewInt32(table, "seller_id")

	x.fillFieldMap()

	return x
}

func (x *xHostSeller) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xHostSeller) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 2)
	x.fieldMap["host"] = x.Host
	x.fieldMap["seller_id"] = x.SellerID
}

func (x xHostSeller) clone(db *gorm.DB) xHostSeller {
	x.xHostSellerDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xHostSeller) replaceDB(db *gorm.DB) xHostSeller {
	x.xHostSellerDo.ReplaceDB(db)
	return x
}

type xHostSellerDo struct{ gen.DO }

type IXHostSellerDo interface {
	gen.SubQuery
	Debug() IXHostSellerDo
	WithContext(ctx context.Context) IXHostSellerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXHostSellerDo
	WriteDB() IXHostSellerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXHostSellerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXHostSellerDo
	Not(conds ...gen.Condition) IXHostSellerDo
	Or(conds ...gen.Condition) IXHostSellerDo
	Select(conds ...field.Expr) IXHostSellerDo
	Where(conds ...gen.Condition) IXHostSellerDo
	Order(conds ...field.Expr) IXHostSellerDo
	Distinct(cols ...field.Expr) IXHostSellerDo
	Omit(cols ...field.Expr) IXHostSellerDo
	Join(table schema.Tabler, on ...field.Expr) IXHostSellerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXHostSellerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXHostSellerDo
	Group(cols ...field.Expr) IXHostSellerDo
	Having(conds ...gen.Condition) IXHostSellerDo
	Limit(limit int) IXHostSellerDo
	Offset(offset int) IXHostSellerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXHostSellerDo
	Unscoped() IXHostSellerDo
	Create(values ...*model.XHostSeller) error
	CreateInBatches(values []*model.XHostSeller, batchSize int) error
	Save(values ...*model.XHostSeller) error
	First() (*model.XHostSeller, error)
	Take() (*model.XHostSeller, error)
	Last() (*model.XHostSeller, error)
	Find() ([]*model.XHostSeller, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XHostSeller, err error)
	FindInBatches(result *[]*model.XHostSeller, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XHostSeller) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXHostSellerDo
	Assign(attrs ...field.AssignExpr) IXHostSellerDo
	Joins(fields ...field.RelationField) IXHostSellerDo
	Preload(fields ...field.RelationField) IXHostSellerDo
	FirstOrInit() (*model.XHostSeller, error)
	FirstOrCreate() (*model.XHostSeller, error)
	FindByPage(offset int, limit int) (result []*model.XHostSeller, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXHostSellerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xHostSellerDo) Debug() IXHostSellerDo {
	return x.withDO(x.DO.Debug())
}

func (x xHostSellerDo) WithContext(ctx context.Context) IXHostSellerDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xHostSellerDo) ReadDB() IXHostSellerDo {
	return x.Clauses(dbresolver.Read)
}

func (x xHostSellerDo) WriteDB() IXHostSellerDo {
	return x.Clauses(dbresolver.Write)
}

func (x xHostSellerDo) Session(config *gorm.Session) IXHostSellerDo {
	return x.withDO(x.DO.Session(config))
}

func (x xHostSellerDo) Clauses(conds ...clause.Expression) IXHostSellerDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xHostSellerDo) Returning(value interface{}, columns ...string) IXHostSellerDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xHostSellerDo) Not(conds ...gen.Condition) IXHostSellerDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xHostSellerDo) Or(conds ...gen.Condition) IXHostSellerDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xHostSellerDo) Select(conds ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xHostSellerDo) Where(conds ...gen.Condition) IXHostSellerDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xHostSellerDo) Order(conds ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xHostSellerDo) Distinct(cols ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xHostSellerDo) Omit(cols ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xHostSellerDo) Join(table schema.Tabler, on ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xHostSellerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xHostSellerDo) RightJoin(table schema.Tabler, on ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xHostSellerDo) Group(cols ...field.Expr) IXHostSellerDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xHostSellerDo) Having(conds ...gen.Condition) IXHostSellerDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xHostSellerDo) Limit(limit int) IXHostSellerDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xHostSellerDo) Offset(offset int) IXHostSellerDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xHostSellerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXHostSellerDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xHostSellerDo) Unscoped() IXHostSellerDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xHostSellerDo) Create(values ...*model.XHostSeller) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xHostSellerDo) CreateInBatches(values []*model.XHostSeller, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xHostSellerDo) Save(values ...*model.XHostSeller) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xHostSellerDo) First() (*model.XHostSeller, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XHostSeller), nil
	}
}

func (x xHostSellerDo) Take() (*model.XHostSeller, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XHostSeller), nil
	}
}

func (x xHostSellerDo) Last() (*model.XHostSeller, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XHostSeller), nil
	}
}

func (x xHostSellerDo) Find() ([]*model.XHostSeller, error) {
	result, err := x.DO.Find()
	return result.([]*model.XHostSeller), err
}

func (x xHostSellerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XHostSeller, err error) {
	buf := make([]*model.XHostSeller, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xHostSellerDo) FindInBatches(result *[]*model.XHostSeller, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xHostSellerDo) Attrs(attrs ...field.AssignExpr) IXHostSellerDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xHostSellerDo) Assign(attrs ...field.AssignExpr) IXHostSellerDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xHostSellerDo) Joins(fields ...field.RelationField) IXHostSellerDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xHostSellerDo) Preload(fields ...field.RelationField) IXHostSellerDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xHostSellerDo) FirstOrInit() (*model.XHostSeller, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XHostSeller), nil
	}
}

func (x xHostSellerDo) FirstOrCreate() (*model.XHostSeller, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XHostSeller), nil
	}
}

func (x xHostSellerDo) FindByPage(offset int, limit int) (result []*model.XHostSeller, count int64, err error) {
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

func (x xHostSellerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xHostSellerDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xHostSellerDo) Delete(models ...*model.XHostSeller) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xHostSellerDo) withDO(do gen.Dao) *xHostSellerDo {
	x.DO = *do.(*gen.DO)
	return x
}
