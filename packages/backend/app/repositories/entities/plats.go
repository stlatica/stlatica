// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entities

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/stlatica/stlatica/packages/backend/app/domains/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PlatBase is an object representing the database table.
type PlatBase struct { // ulid
	PlatID types.PlatID `boil:"plat_id" json:"plat_id" toml:"plat_id" yaml:"plat_id"`
	// user_id
	UserID types.UserID `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// body text
	Content string `boil:"content" json:"content" toml:"content" yaml:"content"`
	// Unix time
	CreatedAt types.UnixTime `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// Unix time
	UpdatedAt types.UnixTime `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *platR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L platL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PlatBaseColumns = struct {
	PlatID    string
	UserID    string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	PlatID:    "plat_id",
	UserID:    "user_id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var PlatBaseTableColumns = struct {
	PlatID    string
	UserID    string
	Content   string
	CreatedAt string
	UpdatedAt string
}{
	PlatID:    "plats.plat_id",
	UserID:    "plats.user_id",
	Content:   "plats.content",
	CreatedAt: "plats.created_at",
	UpdatedAt: "plats.updated_at",
}

// Generated where

type whereHelpertypes_PlatID struct{ field string }

func (w whereHelpertypes_PlatID) EQ(x types.PlatID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_PlatID) NEQ(x types.PlatID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_PlatID) LT(x types.PlatID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_PlatID) LTE(x types.PlatID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_PlatID) GT(x types.PlatID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_PlatID) GTE(x types.PlatID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertypes_UserID struct{ field string }

func (w whereHelpertypes_UserID) EQ(x types.UserID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_UserID) NEQ(x types.UserID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_UserID) LT(x types.UserID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_UserID) LTE(x types.UserID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_UserID) GT(x types.UserID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_UserID) GTE(x types.UserID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_UnixTime struct{ field string }

func (w whereHelpertypes_UnixTime) EQ(x types.UnixTime) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_UnixTime) NEQ(x types.UnixTime) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_UnixTime) LT(x types.UnixTime) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_UnixTime) LTE(x types.UnixTime) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_UnixTime) GT(x types.UnixTime) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_UnixTime) GTE(x types.UnixTime) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PlatBaseWhere = struct {
	PlatID    whereHelpertypes_PlatID
	UserID    whereHelpertypes_UserID
	Content   whereHelperstring
	CreatedAt whereHelpertypes_UnixTime
	UpdatedAt whereHelpertypes_UnixTime
}{
	PlatID:    whereHelpertypes_PlatID{field: "`plats`.`plat_id`"},
	UserID:    whereHelpertypes_UserID{field: "`plats`.`user_id`"},
	Content:   whereHelperstring{field: "`plats`.`content`"},
	CreatedAt: whereHelpertypes_UnixTime{field: "`plats`.`created_at`"},
	UpdatedAt: whereHelpertypes_UnixTime{field: "`plats`.`updated_at`"},
}

// PlatBaseRels is where relationship names are stored.
var PlatBaseRels = struct {
	User string
}{
	User: "User",
}

// platR is where relationships are stored.
type platR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*platR) NewStruct() *platR {
	return &platR{}
}

func (r *platR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// platL is where Load methods for each relationship are stored.
type platL struct{}

var (
	platAllColumns            = []string{"plat_id", "user_id", "content", "created_at", "updated_at"}
	platColumnsWithoutDefault = []string{"plat_id", "user_id", "content", "created_at", "updated_at"}
	platColumnsWithDefault    = []string{}
	platPrimaryKeyColumns     = []string{"plat_id"}
	platGeneratedColumns      = []string{}
)

type (
	// PlatBaseSlice is an alias for a slice of pointers to PlatBase.
	// This should almost always be used instead of []PlatBase.
	PlatBaseSlice []*PlatBase

	platQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	platType                 = reflect.TypeOf(&PlatBase{})
	platMapping              = queries.MakeStructMapping(platType)
	platPrimaryKeyMapping, _ = queries.BindMapping(platType, platMapping, platPrimaryKeyColumns)
	platInsertCacheMut       sync.RWMutex
	platInsertCache          = make(map[string]insertCache)
	platUpdateCacheMut       sync.RWMutex
	platUpdateCache          = make(map[string]updateCache)
	platUpsertCacheMut       sync.RWMutex
	platUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single plat record from the query.
func (q platQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PlatBase, error) {
	o := &PlatBase{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entities: failed to execute a one query for plats")
	}

	return o, nil
}

// All returns all PlatBase records from the query.
func (q platQuery) All(ctx context.Context, exec boil.ContextExecutor) (PlatBaseSlice, error) {
	var o []*PlatBase

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entities: failed to assign all query results to PlatBase slice")
	}

	return o, nil
}

// Count returns the count of all PlatBase records in the query.
func (q platQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entities: failed to count plats rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q platQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entities: failed to check if plats exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *PlatBase) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`user_id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (platL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePlatBase interface{}, mods queries.Applicator) error {
	var slice []*PlatBase
	var object *PlatBase

	if singular {
		var ok bool
		object, ok = maybePlatBase.(*PlatBase)
		if !ok {
			object = new(PlatBase)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePlatBase)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePlatBase))
			}
		}
	} else {
		s, ok := maybePlatBase.(*[]*PlatBase)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePlatBase)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePlatBase))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &platR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &platR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Plats = append(foreign.R.Plats, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.UserID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Plats = append(foreign.R.Plats, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the plat to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Plats.
func (o *PlatBase) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `plats` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, platPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.PlatID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.UserID)
	if o.R == nil {
		o.R = &platR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Plats: PlatBaseSlice{o},
		}
	} else {
		related.R.Plats = append(related.R.Plats, o)
	}

	return nil
}

// Plats retrieves all the records using an executor.
func Plats(mods ...qm.QueryMod) platQuery {
	mods = append(mods, qm.From("`plats`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`plats`.*"})
	}

	return platQuery{q}
}

// FindPlatBase retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPlatBase(ctx context.Context, exec boil.ContextExecutor, platID types.PlatID, selectCols ...string) (*PlatBase, error) {
	platObj := &PlatBase{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `plats` where `plat_id`=?", sel,
	)

	q := queries.Raw(query, platID)

	err := q.Bind(ctx, exec, platObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entities: unable to select from plats")
	}

	return platObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PlatBase) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entities: no plats provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(platColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	platInsertCacheMut.RLock()
	cache, cached := platInsertCache[key]
	platInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			platAllColumns,
			platColumnsWithDefault,
			platColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(platType, platMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(platType, platMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `plats` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `plats` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `plats` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, platPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "entities: unable to insert into plats")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PlatID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to populate default values for plats")
	}

CacheNoHooks:
	if !cached {
		platInsertCacheMut.Lock()
		platInsertCache[key] = cache
		platInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PlatBase.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PlatBase) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	platUpdateCacheMut.RLock()
	cache, cached := platUpdateCache[key]
	platUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			platAllColumns,
			platPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("entities: unable to update plats, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `plats` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, platPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(platType, platMapping, append(wl, platPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to update plats row")
	}

	if !cached {
		platUpdateCacheMut.Lock()
		platUpdateCache[key] = cache
		platUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q platQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "entities: unable to update all for plats")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PlatBaseSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("entities: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), platPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `plats` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, platPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to update all in plat slice")
	}

	return nil
}

// Delete deletes a single PlatBase record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PlatBase) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("entities: no PlatBase provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), platPrimaryKeyMapping)
	sql := "DELETE FROM `plats` WHERE `plat_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to delete from plats")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q platQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("entities: no platQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "entities: unable to delete all from plats")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PlatBaseSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), platPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `plats` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, platPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to delete all from plat slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PlatBase) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPlatBase(ctx, exec, o.PlatID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PlatBaseSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PlatBaseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), platPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `plats`.* FROM `plats` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, platPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entities: unable to reload all in PlatBaseSlice")
	}

	*o = slice

	return nil
}

// PlatBaseExists checks if the PlatBase row exists.
func PlatBaseExists(ctx context.Context, exec boil.ContextExecutor, platID types.PlatID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `plats` where `plat_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, platID)
	}
	row := exec.QueryRowContext(ctx, sql, platID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entities: unable to check if plats exists")
	}

	return exists, nil
}

// Exists checks if the PlatBase row exists.
func (o *PlatBase) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PlatBaseExists(ctx, exec, o.PlatID)
}
