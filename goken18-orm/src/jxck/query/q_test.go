package q

import (
	"jxck/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	db := NewDB()
	var actual string = db.Select("name", "age", "email").From("users").Query().String()
	var expected string = "SELECT name, age, email FROM users"
	assert.Equal(t, actual, expected)
}

func TestWhere(t *testing.T) {
	db := &DB{
		query: new(Query),
	}
	var actual string = db.Select("name").From("users").Where("id = ? and age > ?", "1", "20").Query().String()
	var expected string = "SELECT name FROM users WHERE id = 1 and age > 20"
	assert.Equal(t, actual, expected)
}

func TestOrderDesc(t *testing.T) {
	db := &DB{
		query: new(Query),
	}
	var actual string = db.Select("name").From("users").Where("id = 1").OrderDesc().Query().String()
	var expected string = "SELECT name FROM users WHERE id = 1 ORDER BY DESC"
	assert.Equal(t, actual, expected)
}

func TestJoin(t *testing.T) {
	db := &DB{
		query: new(Query),
	}
	var actual string = db.Select("name").From("users").Where("id = 1").Join("depts").On("users.dept_id = depts.id").OrderDesc().Query().String()
	var expected string = "SELECT name FROM users WHERE id = 1 JOIN depts ON users.dept_id = depts.id ORDER BY DESC"
	assert.Equal(t, actual, expected)
}
