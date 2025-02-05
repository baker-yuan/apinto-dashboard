package store

import (
	"context"
	"go/ast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"reflect"
)

type Table interface {
	schema.Tabler
	IdValue() int
}
type DBInfo struct {
	Addr string
	User string
	DB   string
}
type IDB interface {
	DB(ctx context.Context) *gorm.DB
	IsTxCtx(ctx context.Context) bool
	Info() DBInfo
}

type IBaseStore[T any] interface {
	IDB
	Get(ctx context.Context, id int) (*T, error)
	Save(ctx context.Context, t *T) error
	UpdateByUnique(ctx context.Context, t *T, uniques []string) error
	Delete(ctx context.Context, id ...int) (int, error)
	UpdateWhere(ctx context.Context, t *T, m map[string]interface{}) (int, error)
	Update(ctx context.Context, t *T) (int, error)
	DeleteWhere(ctx context.Context, m map[string]interface{}) (int, error)
	Insert(ctx context.Context, t ...*T) error
	List(ctx context.Context, m map[string]interface{}) ([]*T, error)
	ListQuery(ctx context.Context, sql string, args []interface{}, order string) ([]*T, error)
	First(ctx context.Context, m map[string]interface{}) (*T, error)
	FirstQuery(ctx context.Context, sql string, args []interface{}, order string) (*T, error)
	ListPage(ctx context.Context, sql string, pageNum, pageSize int, args []interface{}, order string) ([]*T, int, error)
	Transaction(ctx context.Context, f func(txCtx context.Context) error) error
}

var TxContextKey = struct{}{}

type BaseStore[T any] struct {
	IDB
	UniqueList []string
	TargetType *T
}

func CreateStore[T any](db IDB) *BaseStore[T] {
	b := &BaseStore[T]{
		IDB:        db,
		TargetType: new(T),
	}
	modelType := reflect.TypeOf(new(T)).Elem()
	for i := 0; i < modelType.NumField(); i++ {
		if fieldStruct := modelType.Field(i); ast.IsExported(fieldStruct.Name) {
			tagSetting := schema.ParseTagSetting(fieldStruct.Tag.Get("gorm"), ";")
			if _, ok := tagSetting["UNIQUEINDEX"]; ok {
				b.UniqueList = append(b.UniqueList, tagSetting["COLUMN"])
			}
		}
	}

	return b
}

func (b *BaseStore[T]) Get(ctx context.Context, id int) (*T, error) {
	value := new(T)
	err := b.DB(ctx).First(value, id).Error
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *BaseStore[T]) Save(ctx context.Context, t *T) error {

	var v interface{} = t
	if table, ok := v.(Table); ok {

		if table.IdValue() != 0 {
			return b.DB(ctx).Save(t).Error
		}
		//没查到主键ID的数据 看看有没有唯一索引 有唯一索引 用唯一索引更新所有字段
		if len(b.UniqueList) > 0 {
			return b.UpdateByUnique(ctx, t, b.UniqueList)
		}
	}
	return b.Insert(ctx, t)
}

func (b *BaseStore[T]) UpdateByUnique(ctx context.Context, t *T, uniques []string) error {
	columns := make([]clause.Column, 0, len(uniques))
	for _, unique := range uniques {
		columns = append(columns, clause.Column{
			Name: unique,
		})
	}
	return b.DB(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		UpdateAll: true,
	}).Create(t).Error
}

func (b *BaseStore[T]) Delete(ctx context.Context, id ...int) (int, error) {

	result := b.DB(ctx).Delete(b.TargetType, id)

	return int(result.RowsAffected), result.Error
}

func (b *BaseStore[T]) UpdateWhere(ctx context.Context, t *T, m map[string]interface{}) (int, error) {

	result := b.DB(ctx).Model(t).Updates(m)

	return int(result.RowsAffected), result.Error
}

func (b *BaseStore[T]) Update(ctx context.Context, t *T) (int, error) {

	result := b.DB(ctx).Updates(t)

	return int(result.RowsAffected), result.Error
}
func (b *BaseStore[T]) DeleteWhere(ctx context.Context, m map[string]interface{}) (int, error) {
	if len(m) == 0 {
		return 0, gorm.ErrMissingWhereClause
	}
	result := b.DB(ctx).Where(m).Delete(b.TargetType)

	return int(result.RowsAffected), result.Error
}

func (b *BaseStore[T]) Insert(ctx context.Context, t ...*T) error {

	return b.DB(ctx).Create(t).Error
}

func (b *BaseStore[T]) List(ctx context.Context, m map[string]interface{}) ([]*T, error) {
	list := make([]*T, 0)

	err := b.DB(ctx).Where(m).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
func (b *BaseStore[T]) ListQuery(ctx context.Context, where string, args []interface{}, order string) ([]*T, error) {
	list := make([]*T, 0)
	db := b.DB(ctx)
	db = db.Where(where, args...)
	if order != "" {
		db = db.Order(order)
	}
	err := db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (b *BaseStore[T]) First(ctx context.Context, m map[string]interface{}) (*T, error) {
	value := new(T)
	db := b.DB(ctx)

	err := db.Where(m).First(value).Error
	if err != nil {
		return nil, err
	}

	return value, nil
}
func (b *BaseStore[T]) FirstQuery(ctx context.Context, where string, args []interface{}, order string) (*T, error) {
	value := new(T)
	db := b.DB(ctx)
	if order != "" {
		db = db.Order(order)
	}
	err := db.Where(where, args...).Take(value).Error
	if err != nil {
		return nil, err
	}

	return value, nil
}
func (b *BaseStore[T]) ListPage(ctx context.Context, where string, pageNum, pageSize int, args []interface{}, order string) ([]*T, int, error) {
	list := make([]*T, 0, pageSize)
	db := b.DB(ctx).Where(where, args...)
	if order != "" {
		db = db.Order(order)
	}
	count := int64(0)
	err := db.Model(list).Count(&count).Limit(pageSize).Offset(PageIndex(pageNum, pageSize)).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, int(count), nil
}

// Transaction 执行事务
func (b *BaseStore[T]) Transaction(ctx context.Context, f func(context.Context) error) error {
	return b.DB(ctx).Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, TxContextKey, tx)
		return f(txCtx)
	})
}
