/*
 * Copyright (c) 2024 OceanBase.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package oceanbase

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/oceanbase/obshell/agent/repository/model/oceanbase"
)

func newUpgradePkgInfo(db *gorm.DB, opts ...gen.DOOption) upgradePkgInfo {
	_upgradePkgInfo := upgradePkgInfo{}

	_upgradePkgInfo.upgradePkgInfoDo.UseDB(db, opts...)
	_upgradePkgInfo.upgradePkgInfoDo.UseModel(&oceanbase.UpgradePkgInfo{})

	tableName := _upgradePkgInfo.upgradePkgInfoDo.TableName()
	_upgradePkgInfo.ALL = field.NewAsterisk(tableName)
	_upgradePkgInfo.PkgId = field.NewInt(tableName, "pkg_id")
	_upgradePkgInfo.Name = field.NewString(tableName, "name")
	_upgradePkgInfo.Version = field.NewString(tableName, "version")
	_upgradePkgInfo.ReleaseDistribution = field.NewString(tableName, "release_distribution")
	_upgradePkgInfo.Distribution = field.NewString(tableName, "distribution")
	_upgradePkgInfo.Release = field.NewString(tableName, "release")
	_upgradePkgInfo.Architecture = field.NewString(tableName, "architecture")
	_upgradePkgInfo.Size = field.NewUint64(tableName, "size")
	_upgradePkgInfo.PayloadSize = field.NewUint64(tableName, "payload_size")
	_upgradePkgInfo.ChunkCount = field.NewInt(tableName, "chunk_count")
	_upgradePkgInfo.Md5 = field.NewString(tableName, "md5")
	_upgradePkgInfo.UpgradeDepYaml = field.NewString(tableName, "upgrade_dep_yaml")
	_upgradePkgInfo.GmtModify = field.NewTime(tableName, "gmt_modify")

	_upgradePkgInfo.fillFieldMap()

	return _upgradePkgInfo
}

type upgradePkgInfo struct {
	upgradePkgInfoDo upgradePkgInfoDo

	ALL                 field.Asterisk
	PkgId               field.Int
	Name                field.String
	Version             field.String
	ReleaseDistribution field.String
	Distribution        field.String
	Release             field.String
	Architecture        field.String
	Size                field.Uint64
	PayloadSize         field.Uint64
	ChunkCount          field.Int
	Md5                 field.String
	UpgradeDepYaml      field.String
	GmtModify           field.Time

	fieldMap map[string]field.Expr
}

func (u upgradePkgInfo) Table(newTableName string) *upgradePkgInfo {
	u.upgradePkgInfoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u upgradePkgInfo) As(alias string) *upgradePkgInfo {
	u.upgradePkgInfoDo.DO = *(u.upgradePkgInfoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *upgradePkgInfo) updateTableName(table string) *upgradePkgInfo {
	u.ALL = field.NewAsterisk(table)
	u.PkgId = field.NewInt(table, "pkg_id")
	u.Name = field.NewString(table, "name")
	u.Version = field.NewString(table, "version")
	u.ReleaseDistribution = field.NewString(table, "release_distribution")
	u.Distribution = field.NewString(table, "distribution")
	u.Release = field.NewString(table, "release")
	u.Architecture = field.NewString(table, "architecture")
	u.Size = field.NewUint64(table, "size")
	u.PayloadSize = field.NewUint64(table, "payload_size")
	u.ChunkCount = field.NewInt(table, "chunk_count")
	u.Md5 = field.NewString(table, "md5")
	u.UpgradeDepYaml = field.NewString(table, "upgrade_dep_yaml")
	u.GmtModify = field.NewTime(table, "gmt_modify")

	u.fillFieldMap()

	return u
}

func (u *upgradePkgInfo) WithContext(ctx context.Context) *upgradePkgInfoDo {
	return u.upgradePkgInfoDo.WithContext(ctx)
}

func (u upgradePkgInfo) TableName() string { return u.upgradePkgInfoDo.TableName() }

func (u upgradePkgInfo) Alias() string { return u.upgradePkgInfoDo.Alias() }

func (u *upgradePkgInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *upgradePkgInfo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 13)
	u.fieldMap["pkg_id"] = u.PkgId
	u.fieldMap["name"] = u.Name
	u.fieldMap["version"] = u.Version
	u.fieldMap["release_distribution"] = u.ReleaseDistribution
	u.fieldMap["distribution"] = u.Distribution
	u.fieldMap["release"] = u.Release
	u.fieldMap["architecture"] = u.Architecture
	u.fieldMap["size"] = u.Size
	u.fieldMap["payload_size"] = u.PayloadSize
	u.fieldMap["chunk_count"] = u.ChunkCount
	u.fieldMap["md5"] = u.Md5
	u.fieldMap["upgrade_dep_yaml"] = u.UpgradeDepYaml
	u.fieldMap["gmt_modify"] = u.GmtModify
}

func (u upgradePkgInfo) clone(db *gorm.DB) upgradePkgInfo {
	u.upgradePkgInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u upgradePkgInfo) replaceDB(db *gorm.DB) upgradePkgInfo {
	u.upgradePkgInfoDo.ReplaceDB(db)
	return u
}

type upgradePkgInfoDo struct{ gen.DO }

func (u upgradePkgInfoDo) Debug() *upgradePkgInfoDo {
	return u.withDO(u.DO.Debug())
}

func (u upgradePkgInfoDo) WithContext(ctx context.Context) *upgradePkgInfoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u upgradePkgInfoDo) ReadDB() *upgradePkgInfoDo {
	return u.Clauses(dbresolver.Read)
}

func (u upgradePkgInfoDo) WriteDB() *upgradePkgInfoDo {
	return u.Clauses(dbresolver.Write)
}

func (u upgradePkgInfoDo) Session(config *gorm.Session) *upgradePkgInfoDo {
	return u.withDO(u.DO.Session(config))
}

func (u upgradePkgInfoDo) Clauses(conds ...clause.Expression) *upgradePkgInfoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u upgradePkgInfoDo) Returning(value interface{}, columns ...string) *upgradePkgInfoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u upgradePkgInfoDo) Not(conds ...gen.Condition) *upgradePkgInfoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u upgradePkgInfoDo) Or(conds ...gen.Condition) *upgradePkgInfoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u upgradePkgInfoDo) Select(conds ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u upgradePkgInfoDo) Where(conds ...gen.Condition) *upgradePkgInfoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u upgradePkgInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *upgradePkgInfoDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u upgradePkgInfoDo) Order(conds ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u upgradePkgInfoDo) Distinct(cols ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u upgradePkgInfoDo) Omit(cols ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u upgradePkgInfoDo) Join(table schema.Tabler, on ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u upgradePkgInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u upgradePkgInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u upgradePkgInfoDo) Group(cols ...field.Expr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u upgradePkgInfoDo) Having(conds ...gen.Condition) *upgradePkgInfoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u upgradePkgInfoDo) Limit(limit int) *upgradePkgInfoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u upgradePkgInfoDo) Offset(offset int) *upgradePkgInfoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u upgradePkgInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *upgradePkgInfoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u upgradePkgInfoDo) Unscoped() *upgradePkgInfoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u upgradePkgInfoDo) Create(values ...*oceanbase.UpgradePkgInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u upgradePkgInfoDo) CreateInBatches(values []*oceanbase.UpgradePkgInfo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u upgradePkgInfoDo) Save(values ...*oceanbase.UpgradePkgInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u upgradePkgInfoDo) First() (*oceanbase.UpgradePkgInfo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*oceanbase.UpgradePkgInfo), nil
	}
}

func (u upgradePkgInfoDo) Take() (*oceanbase.UpgradePkgInfo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*oceanbase.UpgradePkgInfo), nil
	}
}

func (u upgradePkgInfoDo) Last() (*oceanbase.UpgradePkgInfo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*oceanbase.UpgradePkgInfo), nil
	}
}

func (u upgradePkgInfoDo) Find() ([]*oceanbase.UpgradePkgInfo, error) {
	result, err := u.DO.Find()
	return result.([]*oceanbase.UpgradePkgInfo), err
}

func (u upgradePkgInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*oceanbase.UpgradePkgInfo, err error) {
	buf := make([]*oceanbase.UpgradePkgInfo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u upgradePkgInfoDo) FindInBatches(result *[]*oceanbase.UpgradePkgInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u upgradePkgInfoDo) Attrs(attrs ...field.AssignExpr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u upgradePkgInfoDo) Assign(attrs ...field.AssignExpr) *upgradePkgInfoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u upgradePkgInfoDo) Joins(fields ...field.RelationField) *upgradePkgInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u upgradePkgInfoDo) Preload(fields ...field.RelationField) *upgradePkgInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u upgradePkgInfoDo) FirstOrInit() (*oceanbase.UpgradePkgInfo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*oceanbase.UpgradePkgInfo), nil
	}
}

func (u upgradePkgInfoDo) FirstOrCreate() (*oceanbase.UpgradePkgInfo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*oceanbase.UpgradePkgInfo), nil
	}
}

func (u upgradePkgInfoDo) FindByPage(offset int, limit int) (result []*oceanbase.UpgradePkgInfo, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u upgradePkgInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u upgradePkgInfoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u upgradePkgInfoDo) Delete(models ...*oceanbase.UpgradePkgInfo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *upgradePkgInfoDo) withDO(do gen.Dao) *upgradePkgInfoDo {
	u.DO = *do.(*gen.DO)
	return u
}
