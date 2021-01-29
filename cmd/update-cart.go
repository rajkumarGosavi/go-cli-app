package cmd

import (
	"mycart/models"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update-cart",
	Short: "Command to add or delete products from cart",
	Run: func(cmd *cobra.Command, args []string) {
		if user.Name == "" {
			logger.Fatalln("User is not active please activate again")
		}
		addProduct, err := cmd.Flags().GetUint("add-product")
		if err != nil {
			logger.Fatalln("Failed to get `add-product` flag\n", err)
		}

		if addProduct != uint(0) {
			product := models.Product{}
			sqlDb.FetchProductDetails(addProduct, &product)

			cartData := models.Cart{}
			sqlDb.FetchCartDetails(user.ID, &cartData)
			cartData.CartProducts = append(cartData.CartProducts, product)
			final := models.Cart{UserID: user.ID, ProductID: product.ID}
			err = sqlDb.UpdateCart(&final)
			if err != nil {
				logger.Fatalln("Failed to add cart products \n", err)
			}
		}

		productIds, err := cmd.Flags().GetUintSlice("remove-products")
		if err != nil {
			logger.Fatalln("Failed to get `remove-products` flag\n", err)
		}
		if len(productIds) > 0 {
			err = sqlDb.DeleteFromCart(productIds)
			if err != nil {
				logger.Fatalln("Failed to delete products from cart with productIDS", productIds, "\n", err)
			}
		}
	},
}

func init() {

	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().Uint("add-product", 0, "update-cart --add-product ProductID ")
	updateCmd.Flags().UintSlice("remove-products", []uint{}, "update-cart --remove-products=productID1,productID2")

}
