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

func newXUser(db *gorm.DB, opts ...gen.DOOption) xUser {
	_xUser := xUser{}

	_xUser.xUserDo.UseDB(db, opts...)
	_xUser.xUserDo.UseModel(&model.XUser{})

	tableName := _xUser.xUserDo.TableName()
	_xUser.ALL = field.NewAsterisk(tableName)
	_xUser.ID = field.NewInt32(tableName, "id")
	_xUser.SellerID = field.NewInt32(tableName, "seller_id")
	_xUser.Account = field.NewString(tableName, "account")
	_xUser.Password = field.NewString(tableName, "password")
	_xUser.IsVisitor = field.NewInt32(tableName, "is_visitor")
	_xUser.State = field.NewInt32(tableName, "state")
	_xUser.Agent = field.NewString(tableName, "agent")
	_xUser.Token = field.NewString(tableName, "token")
	_xUser.LoginIP = field.NewString(tableName, "login_ip")
	_xUser.LoginIPLocation = field.NewString(tableName, "login_ip_location")
	_xUser.LoginCount = field.NewInt32(tableName, "login_count")
	_xUser.LoginTime = field.NewTime(tableName, "login_time")
	_xUser.ChatState = field.NewInt32(tableName, "chat_state")
	_xUser.IsOnline = field.NewInt32(tableName, "is_online")
	_xUser.CreateTime = field.NewTime(tableName, "create_time")

	_xUser.fillFieldMap()

	return _xUser
}

type xUser struct {
	xUserDo

	ALL             field.Asterisk
	ID              field.Int32
	SellerID        field.Int32  // 运营商
	Account         field.String // 用户账号
	Password        field.String // 用户密码
	IsVisitor       field.Int32  // 是否是游客
	State           field.Int32  // 状态 1开启,2关闭
	Agent           field.String // 所属管理员
	Token           field.String // token
	LoginIP         field.String // 登录ip
	LoginIPLocation field.String // 登录ip地区
	LoginCount      field.Int32  // 登录次数
	LoginTime       field.Time   // 登录时间
	ChatState       field.Int32  // 禁言 1是,2否
	IsOnline        field.Int32  // 是否在线
	CreateTime      field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (x xUser) Table(newTableName string) *xUser {
	x.xUserDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xUser) As(alias string) *xUser {
	x.xUserDo.DO = *(x.xUserDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xUser) updateTableName(table string) *xUser {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt32(table, "id")
	x.SellerID = field.NewInt32(table, "seller_id")
	x.Account = field.NewString(table, "account")
	x.Password = field.NewString(table, "password")
	x.IsVisitor = field.NewInt32(table, "is_visitor")
	x.State = field.NewInt32(table, "state")
	x.Agent = field.NewString(table, "agent")
	x.Token = field.NewString(table, "token")
	x.LoginIP = field.NewString(table, "login_ip")
	x.LoginIPLocation = field.NewString(table, "login_ip_location")
	x.LoginCount = field.NewInt32(table, "login_count")
	x.LoginTime = field.NewTime(table, "login_time")
	x.ChatState = field.NewInt32(table, "chat_state")
	x.IsOnline = field.NewInt32(table, "is_online")
	x.CreateTime = field.NewTime(table, "create_time")

	x.fillFieldMap()

	return x
}

func (x *xUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xUser) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 15)
	x.fieldMap["id"] = x.ID
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["account"] = x.Account
	x.fieldMap["password"] = x.Password
	x.fieldMap["is_visitor"] = x.IsVisitor
	x.fieldMap["state"] = x.State
	x.fieldMap["agent"] = x.Agent
	x.fieldMap["token"] = x.Token
	x.fieldMap["login_ip"] = x.LoginIP
	x.fieldMap["login_ip_location"] = x.LoginIPLocation
	x.fieldMap["login_count"] = x.LoginCount
	x.fieldMap["login_time"] = x.LoginTime
	x.fieldMap["chat_state"] = x.ChatState
	x.fieldMap["is_online"] = x.IsOnline
	x.fieldMap["create_time"] = x.CreateTime
}

func (x xUser) clone(db *gorm.DB) xUser {
	x.xUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xUser) replaceDB(db *gorm.DB) xUser {
	x.xUserDo.ReplaceDB(db)
	return x
}

type xUserDo struct{ gen.DO }

type IXUserDo interface {
	gen.SubQuery
	Debug() IXUserDo
	WithContext(ctx context.Context) IXUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXUserDo
	WriteDB() IXUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXUserDo
	Not(conds ...gen.Condition) IXUserDo
	Or(conds ...gen.Condition) IXUserDo
	Select(conds ...field.Expr) IXUserDo
	Where(conds ...gen.Condition) IXUserDo
	Order(conds ...field.Expr) IXUserDo
	Distinct(cols ...field.Expr) IXUserDo
	Omit(cols ...field.Expr) IXUserDo
	Join(table schema.Tabler, on ...field.Expr) IXUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXUserDo
	Group(cols ...field.Expr) IXUserDo
	Having(conds ...gen.Condition) IXUserDo
	Limit(limit int) IXUserDo
	Offset(offset int) IXUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXUserDo
	Unscoped() IXUserDo
	Create(values ...*model.XUser) error
	CreateInBatches(values []*model.XUser, batchSize int) error
	Save(values ...*model.XUser) error
	First() (*model.XUser, error)
	Take() (*model.XUser, error)
	Last() (*model.XUser, error)
	Find() ([]*model.XUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XUser, err error)
	FindInBatches(result *[]*model.XUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXUserDo
	Assign(attrs ...field.AssignExpr) IXUserDo
	Joins(fields ...field.RelationField) IXUserDo
	Preload(fields ...field.RelationField) IXUserDo
	FirstOrInit() (*model.XUser, error)
	FirstOrCreate() (*model.XUser, error)
	FindByPage(offset int, limit int) (result []*model.XUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xUserDo) Debug() IXUserDo {
	return x.withDO(x.DO.Debug())
}

func (x xUserDo) WithContext(ctx context.Context) IXUserDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xUserDo) ReadDB() IXUserDo {
	return x.Clauses(dbresolver.Read)
}

func (x xUserDo) WriteDB() IXUserDo {
	return x.Clauses(dbresolver.Write)
}

func (x xUserDo) Session(config *gorm.Session) IXUserDo {
	return x.withDO(x.DO.Session(config))
}

func (x xUserDo) Clauses(conds ...clause.Expression) IXUserDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xUserDo) Returning(value interface{}, columns ...string) IXUserDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xUserDo) Not(conds ...gen.Condition) IXUserDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xUserDo) Or(conds ...gen.Condition) IXUserDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xUserDo) Select(conds ...field.Expr) IXUserDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xUserDo) Where(conds ...gen.Condition) IXUserDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xUserDo) Order(conds ...field.Expr) IXUserDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xUserDo) Distinct(cols ...field.Expr) IXUserDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xUserDo) Omit(cols ...field.Expr) IXUserDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xUserDo) Join(table schema.Tabler, on ...field.Expr) IXUserDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXUserDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IXUserDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xUserDo) Group(cols ...field.Expr) IXUserDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xUserDo) Having(conds ...gen.Condition) IXUserDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xUserDo) Limit(limit int) IXUserDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xUserDo) Offset(offset int) IXUserDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXUserDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xUserDo) Unscoped() IXUserDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xUserDo) Create(values ...*model.XUser) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xUserDo) CreateInBatches(values []*model.XUser, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xUserDo) Save(values ...*model.XUser) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xUserDo) First() (*model.XUser, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XUser), nil
	}
}

func (x xUserDo) Take() (*model.XUser, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XUser), nil
	}
}

func (x xUserDo) Last() (*model.XUser, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XUser), nil
	}
}

func (x xUserDo) Find() ([]*model.XUser, error) {
	result, err := x.DO.Find()
	return result.([]*model.XUser), err
}

func (x xUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XUser, err error) {
	buf := make([]*model.XUser, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xUserDo) FindInBatches(result *[]*model.XUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xUserDo) Attrs(attrs ...field.AssignExpr) IXUserDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xUserDo) Assign(attrs ...field.AssignExpr) IXUserDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xUserDo) Joins(fields ...field.RelationField) IXUserDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xUserDo) Preload(fields ...field.RelationField) IXUserDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xUserDo) FirstOrInit() (*model.XUser, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XUser), nil
	}
}

func (x xUserDo) FirstOrCreate() (*model.XUser, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XUser), nil
	}
}

func (x xUserDo) FindByPage(offset int, limit int) (result []*model.XUser, count int64, err error) {
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

func (x xUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xUserDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xUserDo) Delete(models ...*model.XUser) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xUserDo) withDO(do gen.Dao) *xUserDo {
	x.DO = *do.(*gen.DO)
	return x
}
