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

func newXAdminLoginLog(db *gorm.DB, opts ...gen.DOOption) xAdminLoginLog {
	_xAdminLoginLog := xAdminLoginLog{}

	_xAdminLoginLog.xAdminLoginLogDo.UseDB(db, opts...)
	_xAdminLoginLog.xAdminLoginLogDo.UseModel(&order_model.XAdminLoginLog{})

	tableName := _xAdminLoginLog.xAdminLoginLogDo.TableName()
	_xAdminLoginLog.ALL = field.NewAsterisk(tableName)
	_xAdminLoginLog.ID = field.NewInt64(tableName, "id")
	_xAdminLoginLog.SellerID = field.NewInt32(tableName, "seller_id")
	_xAdminLoginLog.Account = field.NewString(tableName, "account")
	_xAdminLoginLog.Token = field.NewString(tableName, "token")
	_xAdminLoginLog.LoginIP = field.NewString(tableName, "login_ip")
	_xAdminLoginLog.LoginIPLocation = field.NewString(tableName, "login_ip_location")
	_xAdminLoginLog.Memo = field.NewString(tableName, "memo")
	_xAdminLoginLog.CreateTime = field.NewTime(tableName, "create_time")

	_xAdminLoginLog.fillFieldMap()

	return _xAdminLoginLog
}

type xAdminLoginLog struct {
	xAdminLoginLogDo

	ALL             field.Asterisk
	ID              field.Int64  // 自增Id
	SellerID        field.Int32  // 运营商
	Account         field.String // 账号
	Token           field.String // 登录的token
	LoginIP         field.String // 登录Ip
	LoginIPLocation field.String // 登录ip所在地区
	Memo            field.String // 备注
	CreateTime      field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (x xAdminLoginLog) Table(newTableName string) *xAdminLoginLog {
	x.xAdminLoginLogDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xAdminLoginLog) As(alias string) *xAdminLoginLog {
	x.xAdminLoginLogDo.DO = *(x.xAdminLoginLogDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xAdminLoginLog) updateTableName(table string) *xAdminLoginLog {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.SellerID = field.NewInt32(table, "seller_id")
	x.Account = field.NewString(table, "account")
	x.Token = field.NewString(table, "token")
	x.LoginIP = field.NewString(table, "login_ip")
	x.LoginIPLocation = field.NewString(table, "login_ip_location")
	x.Memo = field.NewString(table, "memo")
	x.CreateTime = field.NewTime(table, "create_time")

	x.fillFieldMap()

	return x
}

func (x *xAdminLoginLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xAdminLoginLog) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 8)
	x.fieldMap["id"] = x.ID
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["account"] = x.Account
	x.fieldMap["token"] = x.Token
	x.fieldMap["login_ip"] = x.LoginIP
	x.fieldMap["login_ip_location"] = x.LoginIPLocation
	x.fieldMap["memo"] = x.Memo
	x.fieldMap["create_time"] = x.CreateTime
}

func (x xAdminLoginLog) clone(db *gorm.DB) xAdminLoginLog {
	x.xAdminLoginLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xAdminLoginLog) replaceDB(db *gorm.DB) xAdminLoginLog {
	x.xAdminLoginLogDo.ReplaceDB(db)
	return x
}

type xAdminLoginLogDo struct{ gen.DO }

type IXAdminLoginLogDo interface {
	gen.SubQuery
	Debug() IXAdminLoginLogDo
	WithContext(ctx context.Context) IXAdminLoginLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXAdminLoginLogDo
	WriteDB() IXAdminLoginLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXAdminLoginLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXAdminLoginLogDo
	Not(conds ...gen.Condition) IXAdminLoginLogDo
	Or(conds ...gen.Condition) IXAdminLoginLogDo
	Select(conds ...field.Expr) IXAdminLoginLogDo
	Where(conds ...gen.Condition) IXAdminLoginLogDo
	Order(conds ...field.Expr) IXAdminLoginLogDo
	Distinct(cols ...field.Expr) IXAdminLoginLogDo
	Omit(cols ...field.Expr) IXAdminLoginLogDo
	Join(table schema.Tabler, on ...field.Expr) IXAdminLoginLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXAdminLoginLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXAdminLoginLogDo
	Group(cols ...field.Expr) IXAdminLoginLogDo
	Having(conds ...gen.Condition) IXAdminLoginLogDo
	Limit(limit int) IXAdminLoginLogDo
	Offset(offset int) IXAdminLoginLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXAdminLoginLogDo
	Unscoped() IXAdminLoginLogDo
	Create(values ...*order_model.XAdminLoginLog) error
	CreateInBatches(values []*order_model.XAdminLoginLog, batchSize int) error
	Save(values ...*order_model.XAdminLoginLog) error
	First() (*order_model.XAdminLoginLog, error)
	Take() (*order_model.XAdminLoginLog, error)
	Last() (*order_model.XAdminLoginLog, error)
	Find() ([]*order_model.XAdminLoginLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*order_model.XAdminLoginLog, err error)
	FindInBatches(result *[]*order_model.XAdminLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*order_model.XAdminLoginLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXAdminLoginLogDo
	Assign(attrs ...field.AssignExpr) IXAdminLoginLogDo
	Joins(fields ...field.RelationField) IXAdminLoginLogDo
	Preload(fields ...field.RelationField) IXAdminLoginLogDo
	FirstOrInit() (*order_model.XAdminLoginLog, error)
	FirstOrCreate() (*order_model.XAdminLoginLog, error)
	FindByPage(offset int, limit int) (result []*order_model.XAdminLoginLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXAdminLoginLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xAdminLoginLogDo) Debug() IXAdminLoginLogDo {
	return x.withDO(x.DO.Debug())
}

func (x xAdminLoginLogDo) WithContext(ctx context.Context) IXAdminLoginLogDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xAdminLoginLogDo) ReadDB() IXAdminLoginLogDo {
	return x.Clauses(dbresolver.Read)
}

func (x xAdminLoginLogDo) WriteDB() IXAdminLoginLogDo {
	return x.Clauses(dbresolver.Write)
}

func (x xAdminLoginLogDo) Session(config *gorm.Session) IXAdminLoginLogDo {
	return x.withDO(x.DO.Session(config))
}

func (x xAdminLoginLogDo) Clauses(conds ...clause.Expression) IXAdminLoginLogDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xAdminLoginLogDo) Returning(value interface{}, columns ...string) IXAdminLoginLogDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xAdminLoginLogDo) Not(conds ...gen.Condition) IXAdminLoginLogDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xAdminLoginLogDo) Or(conds ...gen.Condition) IXAdminLoginLogDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xAdminLoginLogDo) Select(conds ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xAdminLoginLogDo) Where(conds ...gen.Condition) IXAdminLoginLogDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xAdminLoginLogDo) Order(conds ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xAdminLoginLogDo) Distinct(cols ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xAdminLoginLogDo) Omit(cols ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xAdminLoginLogDo) Join(table schema.Tabler, on ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xAdminLoginLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xAdminLoginLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xAdminLoginLogDo) Group(cols ...field.Expr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xAdminLoginLogDo) Having(conds ...gen.Condition) IXAdminLoginLogDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xAdminLoginLogDo) Limit(limit int) IXAdminLoginLogDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xAdminLoginLogDo) Offset(offset int) IXAdminLoginLogDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xAdminLoginLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXAdminLoginLogDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xAdminLoginLogDo) Unscoped() IXAdminLoginLogDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xAdminLoginLogDo) Create(values ...*order_model.XAdminLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xAdminLoginLogDo) CreateInBatches(values []*order_model.XAdminLoginLog, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xAdminLoginLogDo) Save(values ...*order_model.XAdminLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xAdminLoginLogDo) First() (*order_model.XAdminLoginLog, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XAdminLoginLog), nil
	}
}

func (x xAdminLoginLogDo) Take() (*order_model.XAdminLoginLog, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XAdminLoginLog), nil
	}
}

func (x xAdminLoginLogDo) Last() (*order_model.XAdminLoginLog, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XAdminLoginLog), nil
	}
}

func (x xAdminLoginLogDo) Find() ([]*order_model.XAdminLoginLog, error) {
	result, err := x.DO.Find()
	return result.([]*order_model.XAdminLoginLog), err
}

func (x xAdminLoginLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*order_model.XAdminLoginLog, err error) {
	buf := make([]*order_model.XAdminLoginLog, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xAdminLoginLogDo) FindInBatches(result *[]*order_model.XAdminLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xAdminLoginLogDo) Attrs(attrs ...field.AssignExpr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xAdminLoginLogDo) Assign(attrs ...field.AssignExpr) IXAdminLoginLogDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xAdminLoginLogDo) Joins(fields ...field.RelationField) IXAdminLoginLogDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xAdminLoginLogDo) Preload(fields ...field.RelationField) IXAdminLoginLogDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xAdminLoginLogDo) FirstOrInit() (*order_model.XAdminLoginLog, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XAdminLoginLog), nil
	}
}

func (x xAdminLoginLogDo) FirstOrCreate() (*order_model.XAdminLoginLog, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*order_model.XAdminLoginLog), nil
	}
}

func (x xAdminLoginLogDo) FindByPage(offset int, limit int) (result []*order_model.XAdminLoginLog, count int64, err error) {
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

func (x xAdminLoginLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xAdminLoginLogDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xAdminLoginLogDo) Delete(models ...*order_model.XAdminLoginLog) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xAdminLoginLogDo) withDO(do gen.Dao) *xAdminLoginLogDo {
	x.DO = *do.(*gen.DO)
	return x
}