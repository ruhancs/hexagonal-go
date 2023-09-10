package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/ruhancs/hexagonal-go/adapters/db"
	"github.com/ruhancs/hexagonal-go/application"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

//conexao com banco de daados
func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")//db em memoria

	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE PRODUCTS (
		"id" varchar(255),
		"name" varchar(255),
		"price" float,
		"status" string
	);`

	stmt,err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("abc", "Product", 10, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer DB.Close()
	productDb := db.NewProductDb(DB)
	product,err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer DB.Close()
	productDb := db.NewProductDb(DB)
	product := application.NewProduct()
	product.Name = "P1"
	product.Price = 10

	productResult,err := productDb.Save(product)
	
	require.Nil(t,err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())

	//update product
	product.Status = "enabled"
	productResult,err = productDb.Save(product)

	require.Nil(t,err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
}