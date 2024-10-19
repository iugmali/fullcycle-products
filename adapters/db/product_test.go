package db_test

import (
	"database/sql"
	"github.com/iugmali/fullcycle-products/adapters/db"
	"github.com/iugmali/fullcycle-products/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
    				"id" string,
    				"name" string,
    				"price" integer,
    				"status" string
    			);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "product 1", 0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "product 1", product.GetName())
	require.Equal(t, int64(0), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Get_NotFound(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("not_found")
	require.Nil(t, product)
	require.NotNil(t, err)
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "product 2"
	product.Price = 1500
	_, err := productDb.Save(product)
	require.Nil(t, err)
	product.Status = application.ENABLED
	_, err = productDb.Save(product)
	require.Nil(t, err)
}
