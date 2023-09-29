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

// TeamAPIKey is an object representing the database table.
type TeamAPIKey struct {
	APIKey    string    `boil:"api_key" json:"api_key" toml:"api_key" yaml:"api_key"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	TeamID    string    `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`

	R *teamAPIKeyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L teamAPIKeyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TeamAPIKeyColumns = struct {
	APIKey    string
	CreatedAt string
	TeamID    string
}{
	APIKey:    "api_key",
	CreatedAt: "created_at",
	TeamID:    "team_id",
}

var TeamAPIKeyTableColumns = struct {
	APIKey    string
	CreatedAt string
	TeamID    string
}{
	APIKey:    "team_api_keys.api_key",
	CreatedAt: "team_api_keys.created_at",
	TeamID:    "team_api_keys.team_id",
}

// Generated where

var TeamAPIKeyWhere = struct {
	APIKey    whereHelperstring
	CreatedAt whereHelpertime_Time
	TeamID    whereHelperstring
}{
	APIKey:    whereHelperstring{field: "\"team_api_keys\".\"api_key\""},
	CreatedAt: whereHelpertime_Time{field: "\"team_api_keys\".\"created_at\""},
	TeamID:    whereHelperstring{field: "\"team_api_keys\".\"team_id\""},
}

// TeamAPIKeyRels is where relationship names are stored.
var TeamAPIKeyRels = struct {
	Team string
}{
	Team: "Team",
}

// teamAPIKeyR is where relationships are stored.
type teamAPIKeyR struct {
	Team *Team `boil:"Team" json:"Team" toml:"Team" yaml:"Team"`
}

// NewStruct creates a new relationship struct
func (*teamAPIKeyR) NewStruct() *teamAPIKeyR {
	return &teamAPIKeyR{}
}

func (r *teamAPIKeyR) GetTeam() *Team {
	if r == nil {
		return nil
	}
	return r.Team
}

// teamAPIKeyL is where Load methods for each relationship are stored.
type teamAPIKeyL struct{}

var (
	teamAPIKeyAllColumns            = []string{"api_key", "created_at", "team_id"}
	teamAPIKeyColumnsWithoutDefault = []string{"api_key", "team_id"}
	teamAPIKeyColumnsWithDefault    = []string{"created_at"}
	teamAPIKeyPrimaryKeyColumns     = []string{"api_key"}
	teamAPIKeyGeneratedColumns      = []string{}
)

type (
	// TeamAPIKeySlice is an alias for a slice of pointers to TeamAPIKey.
	// This should almost always be used instead of []TeamAPIKey.
	TeamAPIKeySlice []*TeamAPIKey
	// TeamAPIKeyHook is the signature for custom TeamAPIKey hook methods
	TeamAPIKeyHook func(boil.Executor, *TeamAPIKey) error

	teamAPIKeyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	teamAPIKeyType                 = reflect.TypeOf(&TeamAPIKey{})
	teamAPIKeyMapping              = queries.MakeStructMapping(teamAPIKeyType)
	teamAPIKeyPrimaryKeyMapping, _ = queries.BindMapping(teamAPIKeyType, teamAPIKeyMapping, teamAPIKeyPrimaryKeyColumns)
	teamAPIKeyInsertCacheMut       sync.RWMutex
	teamAPIKeyInsertCache          = make(map[string]insertCache)
	teamAPIKeyUpdateCacheMut       sync.RWMutex
	teamAPIKeyUpdateCache          = make(map[string]updateCache)
	teamAPIKeyUpsertCacheMut       sync.RWMutex
	teamAPIKeyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var teamAPIKeyAfterSelectHooks []TeamAPIKeyHook

var teamAPIKeyBeforeInsertHooks []TeamAPIKeyHook
var teamAPIKeyAfterInsertHooks []TeamAPIKeyHook

var teamAPIKeyBeforeUpdateHooks []TeamAPIKeyHook
var teamAPIKeyAfterUpdateHooks []TeamAPIKeyHook

var teamAPIKeyBeforeDeleteHooks []TeamAPIKeyHook
var teamAPIKeyAfterDeleteHooks []TeamAPIKeyHook

var teamAPIKeyBeforeUpsertHooks []TeamAPIKeyHook
var teamAPIKeyAfterUpsertHooks []TeamAPIKeyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TeamAPIKey) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TeamAPIKey) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TeamAPIKey) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TeamAPIKey) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TeamAPIKey) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TeamAPIKey) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TeamAPIKey) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TeamAPIKey) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TeamAPIKey) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range teamAPIKeyAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTeamAPIKeyHook registers your hook function for all future operations.
func AddTeamAPIKeyHook(hookPoint boil.HookPoint, teamAPIKeyHook TeamAPIKeyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		teamAPIKeyAfterSelectHooks = append(teamAPIKeyAfterSelectHooks, teamAPIKeyHook)
	case boil.BeforeInsertHook:
		teamAPIKeyBeforeInsertHooks = append(teamAPIKeyBeforeInsertHooks, teamAPIKeyHook)
	case boil.AfterInsertHook:
		teamAPIKeyAfterInsertHooks = append(teamAPIKeyAfterInsertHooks, teamAPIKeyHook)
	case boil.BeforeUpdateHook:
		teamAPIKeyBeforeUpdateHooks = append(teamAPIKeyBeforeUpdateHooks, teamAPIKeyHook)
	case boil.AfterUpdateHook:
		teamAPIKeyAfterUpdateHooks = append(teamAPIKeyAfterUpdateHooks, teamAPIKeyHook)
	case boil.BeforeDeleteHook:
		teamAPIKeyBeforeDeleteHooks = append(teamAPIKeyBeforeDeleteHooks, teamAPIKeyHook)
	case boil.AfterDeleteHook:
		teamAPIKeyAfterDeleteHooks = append(teamAPIKeyAfterDeleteHooks, teamAPIKeyHook)
	case boil.BeforeUpsertHook:
		teamAPIKeyBeforeUpsertHooks = append(teamAPIKeyBeforeUpsertHooks, teamAPIKeyHook)
	case boil.AfterUpsertHook:
		teamAPIKeyAfterUpsertHooks = append(teamAPIKeyAfterUpsertHooks, teamAPIKeyHook)
	}
}

// One returns a single teamAPIKey record from the query.
func (q teamAPIKeyQuery) One(exec boil.Executor) (*TeamAPIKey, error) {
	o := &TeamAPIKey{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for team_api_keys")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TeamAPIKey records from the query.
func (q teamAPIKeyQuery) All(exec boil.Executor) (TeamAPIKeySlice, error) {
	var o []*TeamAPIKey

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TeamAPIKey slice")
	}

	if len(teamAPIKeyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TeamAPIKey records in the query.
func (q teamAPIKeyQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count team_api_keys rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q teamAPIKeyQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if team_api_keys exists")
	}

	return count > 0, nil
}

// Team pointed to by the foreign key.
func (o *TeamAPIKey) Team(mods ...qm.QueryMod) teamQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TeamID),
	}

	queryMods = append(queryMods, mods...)

	return Teams(queryMods...)
}

// LoadTeam allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (teamAPIKeyL) LoadTeam(e boil.Executor, singular bool, maybeTeamAPIKey interface{}, mods queries.Applicator) error {
	var slice []*TeamAPIKey
	var object *TeamAPIKey

	if singular {
		var ok bool
		object, ok = maybeTeamAPIKey.(*TeamAPIKey)
		if !ok {
			object = new(TeamAPIKey)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTeamAPIKey)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTeamAPIKey))
			}
		}
	} else {
		s, ok := maybeTeamAPIKey.(*[]*TeamAPIKey)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTeamAPIKey)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTeamAPIKey))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &teamAPIKeyR{}
		}
		args = append(args, object.TeamID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &teamAPIKeyR{}
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
		foreign.R.TeamAPIKeys = append(foreign.R.TeamAPIKeys, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TeamID == foreign.ID {
				local.R.Team = foreign
				if foreign.R == nil {
					foreign.R = &teamR{}
				}
				foreign.R.TeamAPIKeys = append(foreign.R.TeamAPIKeys, local)
				break
			}
		}
	}

	return nil
}

// SetTeam of the teamAPIKey to the related item.
// Sets o.R.Team to related.
// Adds o to related.R.TeamAPIKeys.
func (o *TeamAPIKey) SetTeam(exec boil.Executor, insert bool, related *Team) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"team_api_keys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"team_id"}),
		strmangle.WhereClause("\"", "\"", 2, teamAPIKeyPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.APIKey}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TeamID = related.ID
	if o.R == nil {
		o.R = &teamAPIKeyR{
			Team: related,
		}
	} else {
		o.R.Team = related
	}

	if related.R == nil {
		related.R = &teamR{
			TeamAPIKeys: TeamAPIKeySlice{o},
		}
	} else {
		related.R.TeamAPIKeys = append(related.R.TeamAPIKeys, o)
	}

	return nil
}

// TeamAPIKeys retrieves all the records using an executor.
func TeamAPIKeys(mods ...qm.QueryMod) teamAPIKeyQuery {
	mods = append(mods, qm.From("\"team_api_keys\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"team_api_keys\".*"})
	}

	return teamAPIKeyQuery{q}
}

// FindTeamAPIKey retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTeamAPIKey(exec boil.Executor, aPIKey string, selectCols ...string) (*TeamAPIKey, error) {
	teamAPIKeyObj := &TeamAPIKey{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"team_api_keys\" where \"api_key\"=$1", sel,
	)

	q := queries.Raw(query, aPIKey)

	err := q.Bind(nil, exec, teamAPIKeyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from team_api_keys")
	}

	if err = teamAPIKeyObj.doAfterSelectHooks(exec); err != nil {
		return teamAPIKeyObj, err
	}

	return teamAPIKeyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TeamAPIKey) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no team_api_keys provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(teamAPIKeyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	teamAPIKeyInsertCacheMut.RLock()
	cache, cached := teamAPIKeyInsertCache[key]
	teamAPIKeyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			teamAPIKeyAllColumns,
			teamAPIKeyColumnsWithDefault,
			teamAPIKeyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(teamAPIKeyType, teamAPIKeyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(teamAPIKeyType, teamAPIKeyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"team_api_keys\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"team_api_keys\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into team_api_keys")
	}

	if !cached {
		teamAPIKeyInsertCacheMut.Lock()
		teamAPIKeyInsertCache[key] = cache
		teamAPIKeyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the TeamAPIKey.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TeamAPIKey) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	teamAPIKeyUpdateCacheMut.RLock()
	cache, cached := teamAPIKeyUpdateCache[key]
	teamAPIKeyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			teamAPIKeyAllColumns,
			teamAPIKeyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update team_api_keys, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"team_api_keys\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, teamAPIKeyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(teamAPIKeyType, teamAPIKeyMapping, append(wl, teamAPIKeyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update team_api_keys row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for team_api_keys")
	}

	if !cached {
		teamAPIKeyUpdateCacheMut.Lock()
		teamAPIKeyUpdateCache[key] = cache
		teamAPIKeyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q teamAPIKeyQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for team_api_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for team_api_keys")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TeamAPIKeySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamAPIKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"team_api_keys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, teamAPIKeyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in teamAPIKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all teamAPIKey")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TeamAPIKey) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no team_api_keys provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(teamAPIKeyColumnsWithDefault, o)

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

	teamAPIKeyUpsertCacheMut.RLock()
	cache, cached := teamAPIKeyUpsertCache[key]
	teamAPIKeyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			teamAPIKeyAllColumns,
			teamAPIKeyColumnsWithDefault,
			teamAPIKeyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			teamAPIKeyAllColumns,
			teamAPIKeyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert team_api_keys, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(teamAPIKeyPrimaryKeyColumns))
			copy(conflict, teamAPIKeyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"team_api_keys\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(teamAPIKeyType, teamAPIKeyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(teamAPIKeyType, teamAPIKeyMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert team_api_keys")
	}

	if !cached {
		teamAPIKeyUpsertCacheMut.Lock()
		teamAPIKeyUpsertCache[key] = cache
		teamAPIKeyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single TeamAPIKey record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TeamAPIKey) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TeamAPIKey provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), teamAPIKeyPrimaryKeyMapping)
	sql := "DELETE FROM \"team_api_keys\" WHERE \"api_key\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from team_api_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for team_api_keys")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q teamAPIKeyQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no teamAPIKeyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from team_api_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for team_api_keys")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TeamAPIKeySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(teamAPIKeyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamAPIKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"team_api_keys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teamAPIKeyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from teamAPIKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for team_api_keys")
	}

	if len(teamAPIKeyAfterDeleteHooks) != 0 {
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
func (o *TeamAPIKey) Reload(exec boil.Executor) error {
	ret, err := FindTeamAPIKey(exec, o.APIKey)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TeamAPIKeySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TeamAPIKeySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamAPIKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"team_api_keys\".* FROM \"team_api_keys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teamAPIKeyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TeamAPIKeySlice")
	}

	*o = slice

	return nil
}

// TeamAPIKeyExists checks if the TeamAPIKey row exists.
func TeamAPIKeyExists(exec boil.Executor, aPIKey string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"team_api_keys\" where \"api_key\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, aPIKey)
	}
	row := exec.QueryRow(sql, aPIKey)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if team_api_keys exists")
	}

	return exists, nil
}

// Exists checks if the TeamAPIKey row exists.
func (o *TeamAPIKey) Exists(exec boil.Executor) (bool, error) {
	return TeamAPIKeyExists(exec, o.APIKey)
}