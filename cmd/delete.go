package cmd

import (
	"mycart/models"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Allowing to delete products and categories",
	Run: func(cmd *cobra.Command, args []string) {
		if user.IsAdmin {
			categories, err := cmd.Flags().GetUintSlice("categories")
			if err != nil {
				logger.Fatalln("Failed to get `categories` flag \n", err)
			}

			if len(categories) > 0 {
				err = sqlDb.Delete(&models.Category{}, categories)

				if err != nil {
					logger.Fatalln("Failed to delete rows \n", err)
				}
			}

			products, err := cmd.Flags().GetUintSlice("products")

			if err != nil {
				logger.Fatalln("Failed to get `products` flag \n", err)
			}
			if len(products) > 0 {
				err = sqlDb.Delete(&models.Product{}, products)
				if err != nil {
					logger.Fatalln("Failed to delete rows \n", err)
				}
			}
		}
	},
}

func init() {
	if user.IsAdmin {
		rootCmd.AddCommand(deleteCmd)

		// admin level
		deleteCmd.Flags().UintSlice("categories", []uint{}, "delete --categories 1,2")
		deleteCmd.Flags().UintSlice("products", []uint{}, "delete --products 1,2,3")
	}
}
