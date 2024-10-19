package cmd

import (
	"fmt"
	"github.com/iugmali/fullcycle-products/adapters/cli"
	"github.com/spf13/cobra"
	"log"
)

var productId string
var productName string
var productPrice int64

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Manage products",
	Long:  `Manage products using the CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, "list", "", "", 0)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(res)
	},
}
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Manage products",
	Long:  `Manage products using the CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, "get", productId, "", 0)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(res)
	},
}
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a product",
	Long:  `Manage products using the CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, "create", "", productName, productPrice)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(res)
	},
}
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable product",
	Long:  `Enable product by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, "enable", productId, "", 0)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(res)
	},
}
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable product",
	Long:  `Disable product by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, "disable", productId, "", 0)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(res)
	},
}
var setpriceCmd = &cobra.Command{
	Use:   "setprice",
	Short: "Set product price",
	Long:  `Set a product price entering ID and new price.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, "setprice", productId, "", productPrice)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
	rootCmd.AddCommand(setpriceCmd)

	getCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")

	createCmd.Flags().StringVarP(&productName, "name", "n", "", "Product name")
	createCmd.Flags().Int64VarP(&productPrice, "price", "p", 0, "Product price")

	enableCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")

	disableCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")

	setpriceCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	setpriceCmd.Flags().Int64VarP(&productPrice, "price", "p", 0, "Product price")
}
