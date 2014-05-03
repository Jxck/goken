package q

import (
	"fmt"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type STATEMENT string

const (
	SELECT STATEMENT = "SELECT"
	FROM             = "FROM"
	WHERE            = "WHERE"
	ORDER            = "ORDER BY"
	JOIN             = "JOIN"
	ON               = "ON"
)

const (
	COMMA = ","
	SPACE = " "
	ASC   = "ASC"
	DESC  = "DESC"
)

func NewDB() *DB {
	return &DB{
		query: new(Query),
	}
}

type DB struct {
	query *Query
}

type Query string

func (q *Query) Start(stmt STATEMENT, literal string) {
	*q += Query(fmt.Sprintf("%s %s", stmt, literal))
}

func (q *Query) Append(stmt STATEMENT, literal string) {
	*q += Query(fmt.Sprintf(" %s %s", stmt, literal))
}
func (q *Query) String() string {
	return string(*q)
}

func (db *DB) Select(str ...string) *DB {
	s := strings.Join(str, COMMA+SPACE)
	db.query.Start(SELECT, s)
	return db
}

func (db *DB) From(str string) *DB {
	db.query.Append(FROM, str)
	return db
}

func escape(value string) string {
	// TODO: implement me
	return value
}

func (db *DB) Where(str string, values ...string) *DB {
	for _, v := range values {
		str = strings.Replace(str, "?", escape(v), 1)
	}
	db.query.Append(WHERE, str)
	return db
}

func (db *DB) Join(str string) *DB {
	db.query.Append(JOIN, str)
	return db
}

func (db *DB) On(str string) *DB {
	db.query.Append(ON, str)
	return db
}

func (db *DB) OrderAsc() *DB {
	db.query.Append(ORDER, ASC)
	return db
}

func (db *DB) OrderDesc() *DB {
	db.query.Append(ORDER, DESC)
	return db
}

func (db *DB) Query() *Query {
	return db.query
}
