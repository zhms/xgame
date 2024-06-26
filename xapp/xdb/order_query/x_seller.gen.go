// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package order_query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"xapp/xdb/order_model"
)

func newXSeller(db *gorm.DB, opts ...gen.DOOption) xSeller {
	_xSeller := xSeller{}

	_xSeller.xSellerDo.UseDB(db, opts...)
	_xSeller.xSellerDo.UseModel(&order_model.XSeller{})

	tableName := _xSeller.xSellerDo.TableName()
	_xSeller.ALL = field.NewAsterisk(tableName)
	_xSeller.ID = field.NewInt64(tableName, "id")
	_xSeller.SellerID = field.NewInt32(tableName, "seller_id")
	_xSeller.State = field.NewInt32(tableName, "state")
	_xSeller.SellerName = field.NewString(tableName, "seller_name")
	_xSeller.Memo = field.NewString(tableName, "memo")
	_xSeller.CreateTime = field.NewTime(tableName, "create_time")

	_xSeller.fillFieldMap()

	return _xSeller
}

type xSeller struct {
	xSellerDo

	ALL        field.Asterisk
	ID         field.Int64  // 自增Id
	SellerID   field.Int32  // 运营商
	State      field.Int32  // 状态 1开启,2关闭
	SellerName field.String // 运营商名称
	Memo       field.String // 备注
	CreateTime field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (x xSeller) Table(newTableName string) *xSeller {
	x.xSellerDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xSeller) As(alias string) *xSeller {
	x.xSellerDo.DO = *(x.xSellerDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xSeller) updateTableName(table string) *xSeller {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.SellerID = field.NewInt32(table, "seller_id")
	x.State = field.NewInt32(table, "state")
	x.SellerName = field.NewString(table, "seller_name")
	x.Memo = field.NewString(table, "memo")
	x.CreateTime = field.NewTime(table, "create_time")

	x.fillFieldMap()

	return x
}

func (x *xSeller) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xSeller) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 6)
	x.fieldMap["id"] = x.ID
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["state"] = x.State
	x.fieldMap["seller_name"] = x.SellerName
	x.fieldMap["memo"] = x.Memo
	x.fieldMap["create_time"] = x.CreateTime
}

func (x xSeller) clone(db *gorm.DB) xSeller {
	x.xSellerDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xSeller) replaceDB(db *gorm.DB) xSeller {
	x.xSellerDo.ReplaceDB(db)
	return x
}

type xSellerDo struct{ gen.DO }

type IXSellerDo interface {
	gen.SubQuery
	Debug() IXSellerDo
	WithContext(ctx context.Context) IXSellerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXSellerDo
	WriteDB() IXSellerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXSellerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXSellerDo
	Not(conds ...gen.Condition) IXSellerDo
	Or(conds ...gen.Condition) IXSellerDo
	Select(conds ...field.Expr) IXSellerDo
	Where(conds ...gen.Condition) IXSellerDo
	Order(conds ...field.Expr) IXSellerDo
	Distinct(cols ...field.Expr) IXSellerDo
	Omit(cols ...field.Expr) IXSellerDo
	Join(table schema.Tabler, on ...field.Expr) IXSellerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXSellerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXSellerDo
	Group(cols ...field.Expr) IXSellerDo
	Having(conds ...gen.Condition) IXSellerDo
	Limit(limit int) IXSellerDo
	Offset(offset int) IXSellerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXSellerDo
	Unscoped() IXSellerDo
	Create(values ...*order_model.XSeller) error
	CreateInBatches(values []*order_model.XSeller, batchSize int) error
	Save(values ...*order_model.XSeller) error
	First() (*order_model.XSeller, error)
	Take() (*order_model.XSeller, error)
	Last() (*order_model.XSeller, error)
	Find() ([]*order_model.XSeller, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*order_model.XSeller, err error)
	FindInBatches(result *[]*order_model.XSeller, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*order_model.XSeller) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXSellerDo
	Assign(attrs ...field.AssignExpr) IXSellerDo
	Joins(fields ...field.RelationField) IXSellerDo
	Preload(fields ...field.RelationField) IXSellerDo
	FirstOrInit() (*order_model.XSeller, error)
	FirstOrCreate() (*order_model.XSeller, error)
	FindByPage(offset int, limit int) (result []*order_model.XSeller, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXSellerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xSellerDo) Debug() IXSellerDo {
	return x.withDO(x.DO.Debug())
}

func (x xSellerDo) WithContext(ctx context.Context) IXSellerDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xSellerDo) ReadDB() IXSellerDo {
	return x.Clauses(dbresolver.Read)
}

func (x xSellerDo) WriteDB() IXSellerDo {
	return x.Clauses(dbresolver.Write)
}

func (x xSellerDo) Session(config *gorm.Session) IXSellerDo {
	return x.withDO(x.DO.Session(config))
}

func (x xSellerDo) Clauses(conds ...clause.Expression) IXSellerDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xSellerDo) Returning(value interface{}, columns ...string) IXSellerDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xSellerDo) Not(conds ...gen.Condition) IXSellerDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xSellerDo) Or(conds ...gen.Condition) IXSellerDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xSellerDo) Select(conds ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xSellerDo) Where(conds ...gen.Condition) IXSellerDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xSellerDo) Order(conds ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xSellerDo) Distinct(cols ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xSellerDo) Omit(cols ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xSellerDo) Join(table schema.Tabler, on ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xSellerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xSellerDo) RightJoin(table schema.Tabler, on ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xSellerDo) Group(cols ...field.Expr) IXSellerDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xSellerDo) Having(conds ...gen.Condition) IXSellerDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xSellerDo) Limit(limit int) IXSellerDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xSellerDo) Offset(offset int) IXSellerDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xSellerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXSellerDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xSellerDo) Unscoped() IXSellerDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xSellerDo) Create(values ...*order_model.XSeller) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xSellerDo) CreateInBatches(values []*order_model.XSeller, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xSellerDo) Save(values ...*order_model.XSeller) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xSellerDo) First() (*order_model.XSeller, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XSeller), nil
	}
}

func (x xSellerDo) Take() (*order_model.XSeller, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XSeller), nil
	}
}

func (x xSellerDo) Last() (*order_model.XSeller, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XSeller), nil
	}
}

func (x xSellerDo) Find() ([]*order_model.XSeller, error) {
	result, err := x.DO.Find()
	return result.([]*order_model.XSeller), err
}

func (x xSellerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*order_model.XSeller, err error) {
	buf := make([]*order_model.XSeller, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xSellerDo) FindInBatches(result *[]*order_model.XSeller, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xSellerDo) Attrs(attrs ...field.AssignExpr) IXSellerDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xSellerDo) Assign(attrs ...field.AssignExpr) IXSellerDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xSellerDo) Joins(fields ...field.RelationField) IXSellerDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xSellerDo) Preload(fields ...field.RelationField) IXSellerDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xSellerDo) FirstOrInit() (*order_model.XSeller, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XSeller), nil
	}
}

func (x xSellerDo) FirstOrCreate() (*order_model.XSeller, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XSeller), nil
	}
}

func (x xSellerDo) FindByPage(offset int, limit int) (result []*order_model.XSeller, count int64, err error) {
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

func (x xSellerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xSellerDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xSellerDo) Delete(models ...*order_model.XSeller) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xSellerDo) withDO(do gen.Dao) *xSellerDo {
	x.DO = *do.(*gen.DO)
	return x
}
