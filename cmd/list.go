package cmd

import (
	"fmt"
	"mycart/helpers"
	"mycart/models"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [--categories | --products --category=CategoryName | --cart-items | --bills]",
	Short: "Will list either the categories, product of a specified category, cart items, bills",
	Run: func(cmd *cobra.Command, args []string) {

		isCat, err := cmd.Flags().GetBool("categories")
		if err != nil {
			logger.Fatalln("Failed to get `categories` flag\n", err)
		}
		if isCat {
			categories := []models.Category{}
			sqlDb.FetchAllCategories(&categories)
			helpers.PrintCategories(categories)
		}
		isProd, err := cmd.Flags().GetBool("products")
		if err != nil {
			logger.Fatalln("Failed to get `products` flag\n", err)
		}
		if isProd {
			catName, err := cmd.Flags().GetString("categoryName")
			if err != nil {
				fmt.Println("er", err)
			}
			if catName == "" {
				logger.Fatalln("Specify category name")
			}
			products := []models.Product{}
			sqlDb.FetchProductsOfCategory(catName, &products)
			fmt.Println("Category:", catName)
			helpers.PrintProducts(products)
		}

		isCartItems, err := cmd.Flags().GetBool("cart-items")

		if err != nil {
			logger.Fatalln("Failed to get `cart-items` flag\n", err)
		}
		if isCartItems {
			cart := models.Cart{}
			sqlDb.FetchCartDetails(user.ID, &cart)
			cartProducts := []models.Product{}
			sqlDb.FetchCartProducts(cart.UserID, &cartProducts)
			helpers.PrintCart(user.ID, cartProducts)
		}

		if user.IsAdmin {
			isBills, err := cmd.Flags().GetBool("bills")
			if err != nil {
				logger.Fatalln("Failed to get flag `bills`\n", err)
			}
			if isBills {
				bills := []models.Invoice{}
				sqlDb.GetBills(&bills)
				helpers.PrintBillSummary(bills)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().Bool("categories", false, "List categories")
	listCmd.Flags().Bool("products", false, "List products")
	listCmd.Flags().StringP("categoryName", "c", "", "Category name for which products are to be listed")
	listCmd.Flags().Bool("cart-items", false, "List cart items")

	// admin level
	listCmd.Flags().Bool("bills", false, "List all the bills")

}
