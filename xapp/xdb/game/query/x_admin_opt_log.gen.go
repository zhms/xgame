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

func newXAdminOptLog(db *gorm.DB, opts ...gen.DOOption) xAdminOptLog {
	_xAdminOptLog := xAdminOptLog{}

	_xAdminOptLog.xAdminOptLogDo.UseDB(db, opts...)
	_xAdminOptLog.xAdminOptLogDo.UseModel(&model.XAdminOptLog{})

	tableName := _xAdminOptLog.xAdminOptLogDo.TableName()
	_xAdminOptLog.ALL = field.NewAsterisk(tableName)
	_xAdminOptLog.ID = field.NewInt64(tableName, "id")
	_xAdminOptLog.SellerID = field.NewInt32(tableName, "seller_id")
	_xAdminOptLog.Account = field.NewString(tableName, "account")
	_xAdminOptLog.OptName = field.NewString(tableName, "opt_name")
	_xAdminOptLog.ReqPath = field.NewString(tableName, "req_path")
	_xAdminOptLog.ReqData = field.NewString(tableName, "req_data")
	_xAdminOptLog.ReqIP = field.NewString(tableName, "req_ip")
	_xAdminOptLog.ReqIPLocation = field.NewString(tableName, "req_ip_location")
	_xAdminOptLog.CreateTime = field.NewTime(tableName, "create_time")

	_xAdminOptLog.fillFieldMap()

	return _xAdminOptLog
}

type xAdminOptLog struct {
	xAdminOptLogDo

	ALL           field.Asterisk
	ID            field.Int64  // 自增Id
	SellerID      field.Int32  // 运营商
	Account       field.String // 账号
	OptName       field.String // 操作名称
	ReqPath       field.String // 请求路径
	ReqData       field.String // 请求参数
	ReqIP         field.String // 请求Ip
	ReqIPLocation field.String // 请求Ip所在地区
	CreateTime    field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (x xAdminOptLog) Table(newTableName string) *xAdminOptLog {
	x.xAdminOptLogDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xAdminOptLog) As(alias string) *xAdminOptLog {
	x.xAdminOptLogDo.DO = *(x.xAdminOptLogDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xAdminOptLog) updateTableName(table string) *xAdminOptLog {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.SellerID = field.NewInt32(table, "seller_id")
	x.Account = field.NewString(table, "account")
	x.OptName = field.NewString(table, "opt_name")
	x.ReqPath = field.NewString(table, "req_path")
	x.ReqData = field.NewString(table, "req_data")
	x.ReqIP = field.NewString(table, "req_ip")
	x.ReqIPLocation = field.NewString(table, "req_ip_location")
	x.CreateTime = field.NewTime(table, "create_time")

	x.fillFieldMap()

	return x
}

func (x *xAdminOptLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xAdminOptLog) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 9)
	x.fieldMap["id"] = x.ID
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["account"] = x.Account
	x.fieldMap["opt_name"] = x.OptName
	x.fieldMap["req_path"] = x.ReqPath
	x.fieldMap["req_data"] = x.ReqData
	x.fieldMap["req_ip"] = x.ReqIP
	x.fieldMap["req_ip_location"] = x.ReqIPLocation
	x.fieldMap["create_time"] = x.CreateTime
}

func (x xAdminOptLog) clone(db *gorm.DB) xAdminOptLog {
	x.xAdminOptLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xAdminOptLog) replaceDB(db *gorm.DB) xAdminOptLog {
	x.xAdminOptLogDo.ReplaceDB(db)
	return x
}

type xAdminOptLogDo struct{ gen.DO }

type IXAdminOptLogDo interface {
	gen.SubQuery
	Debug() IXAdminOptLogDo
	WithContext(ctx context.Context) IXAdminOptLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXAdminOptLogDo
	WriteDB() IXAdminOptLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXAdminOptLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXAdminOptLogDo
	Not(conds ...gen.Condition) IXAdminOptLogDo
	Or(conds ...gen.Condition) IXAdminOptLogDo
	Select(conds ...field.Expr) IXAdminOptLogDo
	Where(conds ...gen.Condition) IXAdminOptLogDo
	Order(conds ...field.Expr) IXAdminOptLogDo
	Distinct(cols ...field.Expr) IXAdminOptLogDo
	Omit(cols ...field.Expr) IXAdminOptLogDo
	Join(table schema.Tabler, on ...field.Expr) IXAdminOptLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXAdminOptLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXAdminOptLogDo
	Group(cols ...field.Expr) IXAdminOptLogDo
	Having(conds ...gen.Condition) IXAdminOptLogDo
	Limit(limit int) IXAdminOptLogDo
	Offset(offset int) IXAdminOptLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXAdminOptLogDo
	Unscoped() IXAdminOptLogDo
	Create(values ...*model.XAdminOptLog) error
	CreateInBatches(values []*model.XAdminOptLog, batchSize int) error
	Save(values ...*model.XAdminOptLog) error
	First() (*model.XAdminOptLog, error)
	Take() (*model.XAdminOptLog, error)
	Last() (*model.XAdminOptLog, error)
	Find() ([]*model.XAdminOptLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XAdminOptLog, err error)
	FindInBatches(result *[]*model.XAdminOptLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XAdminOptLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXAdminOptLogDo
	Assign(attrs ...field.AssignExpr) IXAdminOptLogDo
	Joins(fields ...field.RelationField) IXAdminOptLogDo
	Preload(fields ...field.RelationField) IXAdminOptLogDo
	FirstOrInit() (*model.XAdminOptLog, error)
	FirstOrCreate() (*model.XAdminOptLog, error)
	FindByPage(offset int, limit int) (result []*model.XAdminOptLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXAdminOptLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xAdminOptLogDo) Debug() IXAdminOptLogDo {
	return x.withDO(x.DO.Debug())
}

func (x xAdminOptLogDo) WithContext(ctx context.Context) IXAdminOptLogDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xAdminOptLogDo) ReadDB() IXAdminOptLogDo {
	return x.Clauses(dbresolver.Read)
}

func (x xAdminOptLogDo) WriteDB() IXAdminOptLogDo {
	return x.Clauses(dbresolver.Write)
}

func (x xAdminOptLogDo) Session(config *gorm.Session) IXAdminOptLogDo {
	return x.withDO(x.DO.Session(config))
}

func (x xAdminOptLogDo) Clauses(conds ...clause.Expression) IXAdminOptLogDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xAdminOptLogDo) Returning(value interface{}, columns ...string) IXAdminOptLogDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xAdminOptLogDo) Not(conds ...gen.Condition) IXAdminOptLogDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xAdminOptLogDo) Or(conds ...gen.Condition) IXAdminOptLogDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xAdminOptLogDo) Select(conds ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xAdminOptLogDo) Where(conds ...gen.Condition) IXAdminOptLogDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xAdminOptLogDo) Order(conds ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xAdminOptLogDo) Distinct(cols ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xAdminOptLogDo) Omit(cols ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xAdminOptLogDo) Join(table schema.Tabler, on ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xAdminOptLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xAdminOptLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xAdminOptLogDo) Group(cols ...field.Expr) IXAdminOptLogDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xAdminOptLogDo) Having(conds ...gen.Condition) IXAdminOptLogDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xAdminOptLogDo) Limit(limit int) IXAdminOptLogDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xAdminOptLogDo) Offset(offset int) IXAdminOptLogDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xAdminOptLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXAdminOptLogDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xAdminOptLogDo) Unscoped() IXAdminOptLogDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xAdminOptLogDo) Create(values ...*model.XAdminOptLog) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xAdminOptLogDo) CreateInBatches(values []*model.XAdminOptLog, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xAdminOptLogDo) Save(values ...*model.XAdminOptLog) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xAdminOptLogDo) First() (*model.XAdminOptLog, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminOptLog), nil
	}
}

func (x xAdminOptLogDo) Take() (*model.XAdminOptLog, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminOptLog), nil
	}
}

func (x xAdminOptLogDo) Last() (*model.XAdminOptLog, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminOptLog), nil
	}
}

func (x xAdminOptLogDo) Find() ([]*model.XAdminOptLog, error) {
	result, err := x.DO.Find()
	return result.([]*model.XAdminOptLog), err
}

func (x xAdminOptLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XAdminOptLog, err error) {
	buf := make([]*model.XAdminOptLog, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xAdminOptLogDo) FindInBatches(result *[]*model.XAdminOptLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xAdminOptLogDo) Attrs(attrs ...field.AssignExpr) IXAdminOptLogDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xAdminOptLogDo) Assign(attrs ...field.AssignExpr) IXAdminOptLogDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xAdminOptLogDo) Joins(fields ...field.RelationField) IXAdminOptLogDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xAdminOptLogDo) Preload(fields ...field.RelationField) IXAdminOptLogDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xAdminOptLogDo) FirstOrInit() (*model.XAdminOptLog, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminOptLog), nil
	}
}

func (x xAdminOptLogDo) FirstOrCreate() (*model.XAdminOptLog, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminOptLog), nil
	}
}

func (x xAdminOptLogDo) FindByPage(offset int, limit int) (result []*model.XAdminOptLog, count int64, err error) {
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

func (x xAdminOptLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xAdminOptLogDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xAdminOptLogDo) Delete(models ...*model.XAdminOptLog) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xAdminOptLogDo) withDO(do gen.Dao) *xAdminOptLogDo {
	x.DO = *do.(*gen.DO)
	return x
}
