package cmd

import (
	"mycart/helpers"
	"mycart/models"

	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Command to buy products.",
	Long:  `Once you have all the products feel free to buy them. Do take the bill.`,
	Run: func(cmd *cobra.Command, args []string) {
		cartProducts := []models.Product{}
		sqlDb.FetchCartProducts(user.ID, &cartProducts)

		var totalAmount, finalAmount, discountedAmount float64
		for _, product := range cartProducts {
			totalAmount += product.Price
		}
		if totalAmount > 10000 {
			discountedAmount = totalAmount - 500
		} else {
			discountedAmount = 0
		}
		finalAmount = totalAmount - discountedAmount

		helpers.PrintBill(finalAmount, discountedAmount, totalAmount, cartProducts)
		bill := models.Invoice{}
		bill.TotalAmount = totalAmount
		bill.FinalAmount = finalAmount
		bill.Discount = discountedAmount
		bill.UserID = user.ID

		// bill.Products =
		err := sqlDb.SaveBill(&bill)
		if err != nil {
			logger.Fatalln("Failed to save bill \n", err)
		}
		err = sqlDb.DeleteFromUserCart(user.ID)
		if err != nil {
			logger.Fatalln("Failed to remove user's cart\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
