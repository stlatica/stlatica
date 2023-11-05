// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Actor is an object representing the database table.
type Actor struct { // ulid
	ActorID types.ActorID `boil:"actor_id" json:"actor_id" toml:"actor_id" yaml:"actor_id"`
	// name
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`
	// Unix time
	CreatedAt types.UnixTime `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *actorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L actorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ActorColumns = struct {
	ActorID   string
	Name      string
	CreatedAt string
}{
	ActorID:   "actor_id",
	Name:      "name",
	CreatedAt: "created_at",
}

var ActorTableColumns = struct {
	ActorID   string
	Name      string
	CreatedAt string
}{
	ActorID:   "actors.actor_id",
	Name:      "actors.name",
	CreatedAt: "actors.created_at",
}

// Generated where

type whereHelpertypes_ActorID struct{ field string }

func (w whereHelpertypes_ActorID) EQ(x types.ActorID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_ActorID) NEQ(x types.ActorID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_ActorID) LT(x types.ActorID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_ActorID) LTE(x types.ActorID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_ActorID) GT(x types.ActorID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_ActorID) GTE(x types.ActorID) qm.QueryMod {
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

var ActorWhere = struct {
	ActorID   whereHelpertypes_ActorID
	Name      whereHelperstring
	CreatedAt whereHelpertypes_UnixTime
}{
	ActorID:   whereHelpertypes_ActorID{field: "`actors`.`actor_id`"},
	Name:      whereHelperstring{field: "`actors`.`name`"},
	CreatedAt: whereHelpertypes_UnixTime{field: "`actors`.`created_at`"},
}

// ActorRels is where relationship names are stored.
var ActorRels = struct {
}{}

// actorR is where relationships are stored.
type actorR struct {
}

// NewStruct creates a new relationship struct
func (*actorR) NewStruct() *actorR {
	return &actorR{}
}

// actorL is where Load methods for each relationship are stored.
type actorL struct{}

var (
	actorAllColumns            = []string{"actor_id", "name", "created_at"}
	actorColumnsWithoutDefault = []string{"actor_id", "name", "created_at"}
	actorColumnsWithDefault    = []string{}
	actorPrimaryKeyColumns     = []string{"actor_id"}
	actorGeneratedColumns      = []string{}
)

type (
	// ActorSlice is an alias for a slice of pointers to Actor.
	// This should almost always be used instead of []Actor.
	ActorSlice []*Actor

	actorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	actorType                 = reflect.TypeOf(&Actor{})
	actorMapping              = queries.MakeStructMapping(actorType)
	actorPrimaryKeyMapping, _ = queries.BindMapping(actorType, actorMapping, actorPrimaryKeyColumns)
	actorInsertCacheMut       sync.RWMutex
	actorInsertCache          = make(map[string]insertCache)
	actorUpdateCacheMut       sync.RWMutex
	actorUpdateCache          = make(map[string]updateCache)
	actorUpsertCacheMut       sync.RWMutex
	actorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single actor record from the query.
func (q actorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Actor, error) {
	o := &Actor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entities: failed to execute a one query for actors")
	}

	return o, nil
}

// All returns all Actor records from the query.
func (q actorQuery) All(ctx context.Context, exec boil.ContextExecutor) (ActorSlice, error) {
	var o []*Actor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entities: failed to assign all query results to Actor slice")
	}

	return o, nil
}

// Count returns the count of all Actor records in the query.
func (q actorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entities: failed to count actors rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q actorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entities: failed to check if actors exists")
	}

	return count > 0, nil
}

// Actors retrieves all the records using an executor.
func Actors(mods ...qm.QueryMod) actorQuery {
	mods = append(mods, qm.From("`actors`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`actors`.*"})
	}

	return actorQuery{q}
}

// FindActor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindActor(ctx context.Context, exec boil.ContextExecutor, actorID types.ActorID, selectCols ...string) (*Actor, error) {
	actorObj := &Actor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `actors` where `actor_id`=?", sel,
	)

	q := queries.Raw(query, actorID)

	err := q.Bind(ctx, exec, actorObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entities: unable to select from actors")
	}

	return actorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Actor) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entities: no actors provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(actorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	actorInsertCacheMut.RLock()
	cache, cached := actorInsertCache[key]
	actorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			actorAllColumns,
			actorColumnsWithDefault,
			actorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(actorType, actorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(actorType, actorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `actors` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `actors` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `actors` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, actorPrimaryKeyColumns))
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
		return errors.Wrap(err, "entities: unable to insert into actors")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ActorID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to populate default values for actors")
	}

CacheNoHooks:
	if !cached {
		actorInsertCacheMut.Lock()
		actorInsertCache[key] = cache
		actorInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Actor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Actor) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	key := makeCacheKey(columns, nil)
	actorUpdateCacheMut.RLock()
	cache, cached := actorUpdateCache[key]
	actorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			actorAllColumns,
			actorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("entities: unable to update actors, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `actors` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, actorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(actorType, actorMapping, append(wl, actorPrimaryKeyColumns...))
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
		return errors.Wrap(err, "entities: unable to update actors row")
	}

	if !cached {
		actorUpdateCacheMut.Lock()
		actorUpdateCache[key] = cache
		actorUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAll updates all rows with the specified column values.
func (q actorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "entities: unable to update all for actors")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ActorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), actorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `actors` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, actorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to update all in actor slice")
	}

	return nil
}

// Delete deletes a single Actor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Actor) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("entities: no Actor provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), actorPrimaryKeyMapping)
	sql := "DELETE FROM `actors` WHERE `actor_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to delete from actors")
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q actorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("entities: no actorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "entities: unable to delete all from actors")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ActorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), actorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `actors` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, actorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "entities: unable to delete all from actor slice")
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Actor) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindActor(ctx, exec, o.ActorID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ActorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ActorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), actorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `actors`.* FROM `actors` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, actorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entities: unable to reload all in ActorSlice")
	}

	*o = slice

	return nil
}

// ActorExists checks if the Actor row exists.
func ActorExists(ctx context.Context, exec boil.ContextExecutor, actorID types.ActorID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `actors` where `actor_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, actorID)
	}
	row := exec.QueryRowContext(ctx, sql, actorID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entities: unable to check if actors exists")
	}

	return exists, nil
}

// Exists checks if the Actor row exists.
func (o *Actor) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ActorExists(ctx, exec, o.ActorID)
}
