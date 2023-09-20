// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Env is an object representing the database table.
type Env struct {
	ID         string        `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  time.Time     `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	TeamID     string        `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`
	Dockerfile string        `boil:"dockerfile" json:"dockerfile" toml:"dockerfile" yaml:"dockerfile"`
	Status     EnvStatusEnum `boil:"status" json:"status" toml:"status" yaml:"status"`
	Public     bool          `boil:"public" json:"public" toml:"public" yaml:"public"`

	R *envR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L envL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EnvColumns = struct {
	ID         string
	CreatedAt  string
	TeamID     string
	Dockerfile string
	Status     string
	Public     string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	TeamID:     "team_id",
	Dockerfile: "dockerfile",
	Status:     "status",
	Public:     "public",
}

var EnvTableColumns = struct {
	ID         string
	CreatedAt  string
	TeamID     string
	Dockerfile string
	Status     string
	Public     string
}{
	ID:         "envs.id",
	CreatedAt:  "envs.created_at",
	TeamID:     "envs.team_id",
	Dockerfile: "envs.dockerfile",
	Status:     "envs.status",
	Public:     "envs.public",
}

// Generated where

type whereHelperEnvStatusEnum struct{ field string }

func (w whereHelperEnvStatusEnum) EQ(x EnvStatusEnum) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperEnvStatusEnum) NEQ(x EnvStatusEnum) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperEnvStatusEnum) LT(x EnvStatusEnum) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperEnvStatusEnum) LTE(x EnvStatusEnum) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperEnvStatusEnum) GT(x EnvStatusEnum) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperEnvStatusEnum) GTE(x EnvStatusEnum) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperEnvStatusEnum) IN(slice []EnvStatusEnum) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperEnvStatusEnum) NIN(slice []EnvStatusEnum) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var EnvWhere = struct {
	ID         whereHelperstring
	CreatedAt  whereHelpertime_Time
	TeamID     whereHelperstring
	Dockerfile whereHelperstring
	Status     whereHelperEnvStatusEnum
	Public     whereHelperbool
}{
	ID:         whereHelperstring{field: "\"envs\".\"id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"envs\".\"created_at\""},
	TeamID:     whereHelperstring{field: "\"envs\".\"team_id\""},
	Dockerfile: whereHelperstring{field: "\"envs\".\"dockerfile\""},
	Status:     whereHelperEnvStatusEnum{field: "\"envs\".\"status\""},
	Public:     whereHelperbool{field: "\"envs\".\"public\""},
}

// EnvRels is where relationship names are stored.
var EnvRels = struct {
	Team string
}{
	Team: "Team",
}

// envR is where relationships are stored.
type envR struct {
	Team *Team `boil:"Team" json:"Team" toml:"Team" yaml:"Team"`
}

// NewStruct creates a new relationship struct
func (*envR) NewStruct() *envR {
	return &envR{}
}

func (r *envR) GetTeam() *Team {
	if r == nil {
		return nil
	}
	return r.Team
}

// envL is where Load methods for each relationship are stored.
type envL struct{}

var (
	envAllColumns            = []string{"id", "created_at", "team_id", "dockerfile", "status", "public"}
	envColumnsWithoutDefault = []string{"id", "team_id", "dockerfile"}
	envColumnsWithDefault    = []string{"created_at", "status", "public"}
	envPrimaryKeyColumns     = []string{"id"}
	envGeneratedColumns      = []string{}
)

type (
	// EnvSlice is an alias for a slice of pointers to Env.
	// This should almost always be used instead of []Env.
	EnvSlice []*Env
	// EnvHook is the signature for custom Env hook methods
	EnvHook func(boil.Executor, *Env) error

	envQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	envType                 = reflect.TypeOf(&Env{})
	envMapping              = queries.MakeStructMapping(envType)
	envPrimaryKeyMapping, _ = queries.BindMapping(envType, envMapping, envPrimaryKeyColumns)
	envInsertCacheMut       sync.RWMutex
	envInsertCache          = make(map[string]insertCache)
	envUpdateCacheMut       sync.RWMutex
	envUpdateCache          = make(map[string]updateCache)
	envUpsertCacheMut       sync.RWMutex
	envUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var envAfterSelectHooks []EnvHook

var envBeforeInsertHooks []EnvHook
var envAfterInsertHooks []EnvHook

var envBeforeUpdateHooks []EnvHook
var envAfterUpdateHooks []EnvHook

var envBeforeDeleteHooks []EnvHook
var envAfterDeleteHooks []EnvHook

var envBeforeUpsertHooks []EnvHook
var envAfterUpsertHooks []EnvHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Env) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range envAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Env) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range envBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Env) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range envAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Env) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range envBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Env) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range envAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Env) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range envBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Env) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range envAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Env) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range envBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Env) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range envAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEnvHook registers your hook function for all future operations.
func AddEnvHook(hookPoint boil.HookPoint, envHook EnvHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		envAfterSelectHooks = append(envAfterSelectHooks, envHook)
	case boil.BeforeInsertHook:
		envBeforeInsertHooks = append(envBeforeInsertHooks, envHook)
	case boil.AfterInsertHook:
		envAfterInsertHooks = append(envAfterInsertHooks, envHook)
	case boil.BeforeUpdateHook:
		envBeforeUpdateHooks = append(envBeforeUpdateHooks, envHook)
	case boil.AfterUpdateHook:
		envAfterUpdateHooks = append(envAfterUpdateHooks, envHook)
	case boil.BeforeDeleteHook:
		envBeforeDeleteHooks = append(envBeforeDeleteHooks, envHook)
	case boil.AfterDeleteHook:
		envAfterDeleteHooks = append(envAfterDeleteHooks, envHook)
	case boil.BeforeUpsertHook:
		envBeforeUpsertHooks = append(envBeforeUpsertHooks, envHook)
	case boil.AfterUpsertHook:
		envAfterUpsertHooks = append(envAfterUpsertHooks, envHook)
	}
}

// One returns a single env record from the query.
func (q envQuery) One(exec boil.Executor) (*Env, error) {
	o := &Env{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for envs")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Env records from the query.
func (q envQuery) All(exec boil.Executor) (EnvSlice, error) {
	var o []*Env

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Env slice")
	}

	if len(envAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Env records in the query.
func (q envQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count envs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q envQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if envs exists")
	}

	return count > 0, nil
}

// Team pointed to by the foreign key.
func (o *Env) Team(mods ...qm.QueryMod) teamQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TeamID),
	}

	queryMods = append(queryMods, mods...)

	return Teams(queryMods...)
}

// LoadTeam allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (envL) LoadTeam(e boil.Executor, singular bool, maybeEnv interface{}, mods queries.Applicator) error {
	var slice []*Env
	var object *Env

	if singular {
		var ok bool
		object, ok = maybeEnv.(*Env)
		if !ok {
			object = new(Env)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEnv)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEnv))
			}
		}
	} else {
		s, ok := maybeEnv.(*[]*Env)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEnv)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEnv))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &envR{}
		}
		args = append(args, object.TeamID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &envR{}
			}

			for _, a := range args {
				if a == obj.TeamID {
					continue Outer
				}
			}

			args = append(args, obj.TeamID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`teams`),
		qm.WhereIn(`teams.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Team")
	}

	var resultSlice []*Team
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Team")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for teams")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for teams")
	}

	if len(teamAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Team = foreign
		if foreign.R == nil {
			foreign.R = &teamR{}
		}
		foreign.R.Envs = append(foreign.R.Envs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TeamID == foreign.ID {
				local.R.Team = foreign
				if foreign.R == nil {
					foreign.R = &teamR{}
				}
				foreign.R.Envs = append(foreign.R.Envs, local)
				break
			}
		}
	}

	return nil
}

// SetTeam of the env to the related item.
// Sets o.R.Team to related.
// Adds o to related.R.Envs.
func (o *Env) SetTeam(exec boil.Executor, insert bool, related *Team) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"envs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"team_id"}),
		strmangle.WhereClause("\"", "\"", 2, envPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TeamID = related.ID
	if o.R == nil {
		o.R = &envR{
			Team: related,
		}
	} else {
		o.R.Team = related
	}

	if related.R == nil {
		related.R = &teamR{
			Envs: EnvSlice{o},
		}
	} else {
		related.R.Envs = append(related.R.Envs, o)
	}

	return nil
}

// Envs retrieves all the records using an executor.
func Envs(mods ...qm.QueryMod) envQuery {
	mods = append(mods, qm.From("\"envs\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"envs\".*"})
	}

	return envQuery{q}
}

// FindEnv retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEnv(exec boil.Executor, iD string, selectCols ...string) (*Env, error) {
	envObj := &Env{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"envs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, envObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from envs")
	}

	if err = envObj.doAfterSelectHooks(exec); err != nil {
		return envObj, err
	}

	return envObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Env) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no envs provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(envColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	envInsertCacheMut.RLock()
	cache, cached := envInsertCache[key]
	envInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			envAllColumns,
			envColumnsWithDefault,
			envColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(envType, envMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(envType, envMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"envs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"envs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into envs")
	}

	if !cached {
		envInsertCacheMut.Lock()
		envInsertCache[key] = cache
		envInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the Env.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Env) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	envUpdateCacheMut.RLock()
	cache, cached := envUpdateCache[key]
	envUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			envAllColumns,
			envPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update envs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"envs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, envPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(envType, envMapping, append(wl, envPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update envs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for envs")
	}

	if !cached {
		envUpdateCacheMut.Lock()
		envUpdateCache[key] = cache
		envUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q envQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for envs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for envs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EnvSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), envPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"envs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, envPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in env slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all env")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Env) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no envs provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(envColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	envUpsertCacheMut.RLock()
	cache, cached := envUpsertCache[key]
	envUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			envAllColumns,
			envColumnsWithDefault,
			envColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			envAllColumns,
			envPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert envs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(envPrimaryKeyColumns))
			copy(conflict, envPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"envs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(envType, envMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(envType, envMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert envs")
	}

	if !cached {
		envUpsertCacheMut.Lock()
		envUpsertCache[key] = cache
		envUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single Env record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Env) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Env provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), envPrimaryKeyMapping)
	sql := "DELETE FROM \"envs\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from envs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for envs")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q envQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no envQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from envs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for envs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EnvSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(envBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), envPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"envs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, envPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from env slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for envs")
	}

	if len(envAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Env) Reload(exec boil.Executor) error {
	ret, err := FindEnv(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EnvSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EnvSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), envPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"envs\".* FROM \"envs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, envPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EnvSlice")
	}

	*o = slice

	return nil
}

// EnvExists checks if the Env row exists.
func EnvExists(exec boil.Executor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"envs\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if envs exists")
	}

	return exists, nil
}

// Exists checks if the Env row exists.
func (o *Env) Exists(exec boil.Executor) (bool, error) {
	return EnvExists(exec, o.ID)
}