package golanggorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:123456789@tcp(localhost:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values(?,?)", "1", "Fandi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?,?)", "2", "Agung").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?,?)", "3", "Alam").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?,?)", "4", "Afifah").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

// Raw SQL with Scan
func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Fandi", sample.Name)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

// SQL Rows with Manual Mapping
func TestSQLRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}

	assert.Equal(t, 4, len(samples))
}

// Gorm ScanRows
func TestScanRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 4, len(samples))
}
