package cmd

import (
	"fmt"
	server2 "github.com/iugmali/fullcycle-products/adapters/web/server"
	"github.com/spf13/cobra"
	"os"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "serve",
	Short: "Webserver",
	Long: `

Webserver that runs on port 9000 and serves the product service.

Available routes:
	
GET /products
GET /products/{id}
POST /products
PATCH /products/{id}/enable
PATCH /products/{id}/disable
PATCH /products/{id}/setprice/{price}

	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebserver()
		server.Service = &productService
		fmt.Println(`Webserver running on port ` + os.Getenv("PORT"))
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
