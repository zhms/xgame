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

func newXChatBanIP(db *gorm.DB, opts ...gen.DOOption) xChatBanIP {
	_xChatBanIP := xChatBanIP{}

	_xChatBanIP.xChatBanIPDo.UseDB(db, opts...)
	_xChatBanIP.xChatBanIPDo.UseModel(&model.XChatBanIP{})

	tableName := _xChatBanIP.xChatBanIPDo.TableName()
	_xChatBanIP.ALL = field.NewAsterisk(tableName)
	_xChatBanIP.ID = field.NewInt32(tableName, "id")
	_xChatBanIP.SellerID = field.NewInt32(tableName, "seller_id")
	_xChatBanIP.IP = field.NewString(tableName, "ip")
	_xChatBanIP.AdminAccount = field.NewString(tableName, "admin_account")
	_xChatBanIP.CreateTime = field.NewTime(tableName, "create_time")

	_xChatBanIP.fillFieldMap()

	return _xChatBanIP
}

type xChatBanIP struct {
	xChatBanIPDo

	ALL          field.Asterisk
	ID           field.Int32
	SellerID     field.Int32
	IP           field.String
	AdminAccount field.String
	CreateTime   field.Time

	fieldMap map[string]field.Expr
}

func (x xChatBanIP) Table(newTableName string) *xChatBanIP {
	x.xChatBanIPDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xChatBanIP) As(alias string) *xChatBanIP {
	x.xChatBanIPDo.DO = *(x.xChatBanIPDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xChatBanIP) updateTableName(table string) *xChatBanIP {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt32(table, "id")
	x.SellerID = field.NewInt32(table, "seller_id")
	x.IP = field.NewString(table, "ip")
	x.AdminAccount = field.NewString(table, "admin_account")
	x.CreateTime = field.NewTime(table, "create_time")

	x.fillFieldMap()

	return x
}

func (x *xChatBanIP) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xChatBanIP) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 5)
	x.fieldMap["id"] = x.ID
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["ip"] = x.IP
	x.fieldMap["admin_account"] = x.AdminAccount
	x.fieldMap["create_time"] = x.CreateTime
}

func (x xChatBanIP) clone(db *gorm.DB) xChatBanIP {
	x.xChatBanIPDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xChatBanIP) replaceDB(db *gorm.DB) xChatBanIP {
	x.xChatBanIPDo.ReplaceDB(db)
	return x
}

type xChatBanIPDo struct{ gen.DO }

type IXChatBanIPDo interface {
	gen.SubQuery
	Debug() IXChatBanIPDo
	WithContext(ctx context.Context) IXChatBanIPDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXChatBanIPDo
	WriteDB() IXChatBanIPDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXChatBanIPDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXChatBanIPDo
	Not(conds ...gen.Condition) IXChatBanIPDo
	Or(conds ...gen.Condition) IXChatBanIPDo
	Select(conds ...field.Expr) IXChatBanIPDo
	Where(conds ...gen.Condition) IXChatBanIPDo
	Order(conds ...field.Expr) IXChatBanIPDo
	Distinct(cols ...field.Expr) IXChatBanIPDo
	Omit(cols ...field.Expr) IXChatBanIPDo
	Join(table schema.Tabler, on ...field.Expr) IXChatBanIPDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXChatBanIPDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXChatBanIPDo
	Group(cols ...field.Expr) IXChatBanIPDo
	Having(conds ...gen.Condition) IXChatBanIPDo
	Limit(limit int) IXChatBanIPDo
	Offset(offset int) IXChatBanIPDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXChatBanIPDo
	Unscoped() IXChatBanIPDo
	Create(values ...*model.XChatBanIP) error
	CreateInBatches(values []*model.XChatBanIP, batchSize int) error
	Save(values ...*model.XChatBanIP) error
	First() (*model.XChatBanIP, error)
	Take() (*model.XChatBanIP, error)
	Last() (*model.XChatBanIP, error)
	Find() ([]*model.XChatBanIP, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XChatBanIP, err error)
	FindInBatches(result *[]*model.XChatBanIP, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XChatBanIP) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXChatBanIPDo
	Assign(attrs ...field.AssignExpr) IXChatBanIPDo
	Joins(fields ...field.RelationField) IXChatBanIPDo
	Preload(fields ...field.RelationField) IXChatBanIPDo
	FirstOrInit() (*model.XChatBanIP, error)
	FirstOrCreate() (*model.XChatBanIP, error)
	FindByPage(offset int, limit int) (result []*model.XChatBanIP, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXChatBanIPDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xChatBanIPDo) Debug() IXChatBanIPDo {
	return x.withDO(x.DO.Debug())
}

func (x xChatBanIPDo) WithContext(ctx context.Context) IXChatBanIPDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xChatBanIPDo) ReadDB() IXChatBanIPDo {
	return x.Clauses(dbresolver.Read)
}

func (x xChatBanIPDo) WriteDB() IXChatBanIPDo {
	return x.Clauses(dbresolver.Write)
}

func (x xChatBanIPDo) Session(config *gorm.Session) IXChatBanIPDo {
	return x.withDO(x.DO.Session(config))
}

func (x xChatBanIPDo) Clauses(conds ...clause.Expression) IXChatBanIPDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xChatBanIPDo) Returning(value interface{}, columns ...string) IXChatBanIPDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xChatBanIPDo) Not(conds ...gen.Condition) IXChatBanIPDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xChatBanIPDo) Or(conds ...gen.Condition) IXChatBanIPDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xChatBanIPDo) Select(conds ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xChatBanIPDo) Where(conds ...gen.Condition) IXChatBanIPDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xChatBanIPDo) Order(conds ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xChatBanIPDo) Distinct(cols ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xChatBanIPDo) Omit(cols ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xChatBanIPDo) Join(table schema.Tabler, on ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xChatBanIPDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xChatBanIPDo) RightJoin(table schema.Tabler, on ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xChatBanIPDo) Group(cols ...field.Expr) IXChatBanIPDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xChatBanIPDo) Having(conds ...gen.Condition) IXChatBanIPDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xChatBanIPDo) Limit(limit int) IXChatBanIPDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xChatBanIPDo) Offset(offset int) IXChatBanIPDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xChatBanIPDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXChatBanIPDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xChatBanIPDo) Unscoped() IXChatBanIPDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xChatBanIPDo) Create(values ...*model.XChatBanIP) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xChatBanIPDo) CreateInBatches(values []*model.XChatBanIP, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xChatBanIPDo) Save(values ...*model.XChatBanIP) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xChatBanIPDo) First() (*model.XChatBanIP, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XChatBanIP), nil
	}
}

func (x xChatBanIPDo) Take() (*model.XChatBanIP, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XChatBanIP), nil
	}
}

func (x xChatBanIPDo) Last() (*model.XChatBanIP, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XChatBanIP), nil
	}
}

func (x xChatBanIPDo) Find() ([]*model.XChatBanIP, error) {
	result, err := x.DO.Find()
	return result.([]*model.XChatBanIP), err
}

func (x xChatBanIPDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XChatBanIP, err error) {
	buf := make([]*model.XChatBanIP, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xChatBanIPDo) FindInBatches(result *[]*model.XChatBanIP, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xChatBanIPDo) Attrs(attrs ...field.AssignExpr) IXChatBanIPDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xChatBanIPDo) Assign(attrs ...field.AssignExpr) IXChatBanIPDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xChatBanIPDo) Joins(fields ...field.RelationField) IXChatBanIPDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xChatBanIPDo) Preload(fields ...field.RelationField) IXChatBanIPDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xChatBanIPDo) FirstOrInit() (*model.XChatBanIP, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XChatBanIP), nil
	}
}

func (x xChatBanIPDo) FirstOrCreate() (*model.XChatBanIP, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XChatBanIP), nil
	}
}

func (x xChatBanIPDo) FindByPage(offset int, limit int) (result []*model.XChatBanIP, count int64, err error) {
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

func (x xChatBanIPDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xChatBanIPDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xChatBanIPDo) Delete(models ...*model.XChatBanIP) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xChatBanIPDo) withDO(do gen.Dao) *xChatBanIPDo {
	x.DO = *do.(*gen.DO)
	return x
}
