// Code generated by ent, DO NOT EDIT.

package teamapikey

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the teamapikey type in the database.
	Label = "team_api_key"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "api_key"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTeamID holds the string denoting the team_id field in the database.
	FieldTeamID = "team_id"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// TeamFieldID holds the string denoting the ID field of the Team.
	TeamFieldID = "id"
	// Table holds the table name of the teamapikey in the database.
	Table = "team_api_keys"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "teams"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "teams"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "team_api_key_team"
)

// Columns holds all SQL columns for teamapikey fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldTeamID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the TeamApiKey queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTeamID orders the results by the team_id field.
func ByTeamID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamID, opts...).ToFunc()
}

// ByTeamCount orders the results by team count.
func ByTeamCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamStep(), opts...)
	}
}

// ByTeam orders the results by team terms.
func ByTeam(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, TeamFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TeamTable, TeamColumn),
	)
}
