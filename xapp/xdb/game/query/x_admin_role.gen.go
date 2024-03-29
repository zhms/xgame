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

func newXAdminRole(db *gorm.DB, opts ...gen.DOOption) xAdminRole {
	_xAdminRole := xAdminRole{}

	_xAdminRole.xAdminRoleDo.UseDB(db, opts...)
	_xAdminRole.xAdminRoleDo.UseModel(&model.XAdminRole{})

	tableName := _xAdminRole.xAdminRoleDo.TableName()
	_xAdminRole.ALL = field.NewAsterisk(tableName)
	_xAdminRole.ID = field.NewInt64(tableName, "id")
	_xAdminRole.SellerID = field.NewInt32(tableName, "seller_id")
	_xAdminRole.RoleName = field.NewString(tableName, "role_name")
	_xAdminRole.Parent = field.NewString(tableName, "parent")
	_xAdminRole.RoleData = field.NewString(tableName, "role_data")
	_xAdminRole.State = field.NewInt32(tableName, "state")
	_xAdminRole.Memo = field.NewString(tableName, "memo")
	_xAdminRole.CreateTime = field.NewTime(tableName, "create_time")

	_xAdminRole.fillFieldMap()

	return _xAdminRole
}

type xAdminRole struct {
	xAdminRoleDo

	ALL        field.Asterisk
	ID         field.Int64  // 自增Id
	SellerID   field.Int32  // 运营商
	RoleName   field.String // 角色名
	Parent     field.String // 上级角色
	RoleData   field.String // 权限数据
	State      field.Int32  // 状态 1开启,2关闭
	Memo       field.String // 备注
	CreateTime field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (x xAdminRole) Table(newTableName string) *xAdminRole {
	x.xAdminRoleDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xAdminRole) As(alias string) *xAdminRole {
	x.xAdminRoleDo.DO = *(x.xAdminRoleDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xAdminRole) updateTableName(table string) *xAdminRole {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.SellerID = field.NewInt32(table, "seller_id")
	x.RoleName = field.NewString(table, "role_name")
	x.Parent = field.NewString(table, "parent")
	x.RoleData = field.NewString(table, "role_data")
	x.State = field.NewInt32(table, "state")
	x.Memo = field.NewString(table, "memo")
	x.CreateTime = field.NewTime(table, "create_time")

	x.fillFieldMap()

	return x
}

func (x *xAdminRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xAdminRole) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 8)
	x.fieldMap["id"] = x.ID
	x.fieldMap["seller_id"] = x.SellerID
	x.fieldMap["role_name"] = x.RoleName
	x.fieldMap["parent"] = x.Parent
	x.fieldMap["role_data"] = x.RoleData
	x.fieldMap["state"] = x.State
	x.fieldMap["memo"] = x.Memo
	x.fieldMap["create_time"] = x.CreateTime
}

func (x xAdminRole) clone(db *gorm.DB) xAdminRole {
	x.xAdminRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xAdminRole) replaceDB(db *gorm.DB) xAdminRole {
	x.xAdminRoleDo.ReplaceDB(db)
	return x
}

type xAdminRoleDo struct{ gen.DO }

type IXAdminRoleDo interface {
	gen.SubQuery
	Debug() IXAdminRoleDo
	WithContext(ctx context.Context) IXAdminRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXAdminRoleDo
	WriteDB() IXAdminRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXAdminRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXAdminRoleDo
	Not(conds ...gen.Condition) IXAdminRoleDo
	Or(conds ...gen.Condition) IXAdminRoleDo
	Select(conds ...field.Expr) IXAdminRoleDo
	Where(conds ...gen.Condition) IXAdminRoleDo
	Order(conds ...field.Expr) IXAdminRoleDo
	Distinct(cols ...field.Expr) IXAdminRoleDo
	Omit(cols ...field.Expr) IXAdminRoleDo
	Join(table schema.Tabler, on ...field.Expr) IXAdminRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXAdminRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXAdminRoleDo
	Group(cols ...field.Expr) IXAdminRoleDo
	Having(conds ...gen.Condition) IXAdminRoleDo
	Limit(limit int) IXAdminRoleDo
	Offset(offset int) IXAdminRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXAdminRoleDo
	Unscoped() IXAdminRoleDo
	Create(values ...*model.XAdminRole) error
	CreateInBatches(values []*model.XAdminRole, batchSize int) error
	Save(values ...*model.XAdminRole) error
	First() (*model.XAdminRole, error)
	Take() (*model.XAdminRole, error)
	Last() (*model.XAdminRole, error)
	Find() ([]*model.XAdminRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XAdminRole, err error)
	FindInBatches(result *[]*model.XAdminRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XAdminRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXAdminRoleDo
	Assign(attrs ...field.AssignExpr) IXAdminRoleDo
	Joins(fields ...field.RelationField) IXAdminRoleDo
	Preload(fields ...field.RelationField) IXAdminRoleDo
	FirstOrInit() (*model.XAdminRole, error)
	FirstOrCreate() (*model.XAdminRole, error)
	FindByPage(offset int, limit int) (result []*model.XAdminRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXAdminRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xAdminRoleDo) Debug() IXAdminRoleDo {
	return x.withDO(x.DO.Debug())
}

func (x xAdminRoleDo) WithContext(ctx context.Context) IXAdminRoleDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xAdminRoleDo) ReadDB() IXAdminRoleDo {
	return x.Clauses(dbresolver.Read)
}

func (x xAdminRoleDo) WriteDB() IXAdminRoleDo {
	return x.Clauses(dbresolver.Write)
}

func (x xAdminRoleDo) Session(config *gorm.Session) IXAdminRoleDo {
	return x.withDO(x.DO.Session(config))
}

func (x xAdminRoleDo) Clauses(conds ...clause.Expression) IXAdminRoleDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xAdminRoleDo) Returning(value interface{}, columns ...string) IXAdminRoleDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xAdminRoleDo) Not(conds ...gen.Condition) IXAdminRoleDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xAdminRoleDo) Or(conds ...gen.Condition) IXAdminRoleDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xAdminRoleDo) Select(conds ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xAdminRoleDo) Where(conds ...gen.Condition) IXAdminRoleDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xAdminRoleDo) Order(conds ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xAdminRoleDo) Distinct(cols ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xAdminRoleDo) Omit(cols ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xAdminRoleDo) Join(table schema.Tabler, on ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xAdminRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xAdminRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xAdminRoleDo) Group(cols ...field.Expr) IXAdminRoleDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xAdminRoleDo) Having(conds ...gen.Condition) IXAdminRoleDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xAdminRoleDo) Limit(limit int) IXAdminRoleDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xAdminRoleDo) Offset(offset int) IXAdminRoleDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xAdminRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXAdminRoleDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xAdminRoleDo) Unscoped() IXAdminRoleDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xAdminRoleDo) Create(values ...*model.XAdminRole) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xAdminRoleDo) CreateInBatches(values []*model.XAdminRole, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xAdminRoleDo) Save(values ...*model.XAdminRole) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xAdminRoleDo) First() (*model.XAdminRole, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminRole), nil
	}
}

func (x xAdminRoleDo) Take() (*model.XAdminRole, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminRole), nil
	}
}

func (x xAdminRoleDo) Last() (*model.XAdminRole, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminRole), nil
	}
}

func (x xAdminRoleDo) Find() ([]*model.XAdminRole, error) {
	result, err := x.DO.Find()
	return result.([]*model.XAdminRole), err
}

func (x xAdminRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XAdminRole, err error) {
	buf := make([]*model.XAdminRole, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xAdminRoleDo) FindInBatches(result *[]*model.XAdminRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xAdminRoleDo) Attrs(attrs ...field.AssignExpr) IXAdminRoleDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xAdminRoleDo) Assign(attrs ...field.AssignExpr) IXAdminRoleDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xAdminRoleDo) Joins(fields ...field.RelationField) IXAdminRoleDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xAdminRoleDo) Preload(fields ...field.RelationField) IXAdminRoleDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xAdminRoleDo) FirstOrInit() (*model.XAdminRole, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminRole), nil
	}
}

func (x xAdminRoleDo) FirstOrCreate() (*model.XAdminRole, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XAdminRole), nil
	}
}

func (x xAdminRoleDo) FindByPage(offset int, limit int) (result []*model.XAdminRole, count int64, err error) {
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

func (x xAdminRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xAdminRoleDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xAdminRoleDo) Delete(models ...*model.XAdminRole) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xAdminRoleDo) withDO(do gen.Dao) *xAdminRoleDo {
	x.DO = *do.(*gen.DO)
	return x
}
