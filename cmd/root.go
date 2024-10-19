package cmd

import (
	"database/sql"
	dbInfra "github.com/iugmali/fullcycle-products/adapters/db"
	"github.com/iugmali/fullcycle-products/application"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var db, _ = sql.Open("sqlite3", "db.sqlite")

var productDb = dbInfra.NewProductDb(db)
var productService = application.ProductService{Persistence: productDb}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "products",
	Short: "CLI products management application",
	Long: `This application is a CLI to manage products. 

	It can create, enable or disable, set price and get product.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome to FullCycle Products")
	},
}

func Execute() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
    				"id" string,
    				"name" string,
    				"price" integer,
    				"status" string
    			)`)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
