package helpers

import (
	"fmt"
	"mycart/models"
	"strings"
)

// PrintProducts - prints the output in terminal in formatted way
func PrintProducts(products []models.Product) {

	var strBuilder strings.Builder

	strBuilder.WriteString("Product Name  | Price \n")

	for _, product := range products {
		strBuilder.WriteString(fmt.Sprintf("%13s | %3f \n", product.Name, product.Price))
	}

	fmt.Println(strBuilder.String())
}

// PrintCategories - prints the output in terminal in formatted way
func PrintCategories(categories []models.Category) {

	var strBuilder strings.Builder

	strBuilder.WriteString(" Category Name  \n")

	for _, category := range categories {
		strBuilder.WriteString(fmt.Sprintf("%14s \n", category.Name))
	}

	fmt.Println(strBuilder.String())
}

// PrintBill - prints details of a bill
func PrintBill(final, discount, total float64, cartProducts []models.Product) {
	var bill strings.Builder
	bill.WriteString("Sr.No  |  Name  |  Price \n")
	for i, product := range cartProducts {
		bill.WriteString(fmt.Sprintf("%6d | %6s |%6f\n", i, product.Name, product.Price))
	}
	bill.WriteString("\n-------------------------------------------------------\n")
	bill.WriteString(fmt.Sprintf("Total Amount:    %f\n", total))
	bill.WriteString(fmt.Sprintf("Discount Amount: %f\n", discount))
	bill.WriteString(fmt.Sprintf("Final Amount:    %f\n", final))
	fmt.Println(bill.String())
}

// PrintBillSummary - prints summary of all bills
func PrintBillSummary(bills []models.Invoice) {
	var summary strings.Builder

	summary.WriteString("UserID  |  Total    | Discount | Final\n")
	for _, b := range bills {
		summary.WriteString(fmt.Sprintf("%7d | %7f | %7f | %7f \n", b.UserID, b.TotalAmount, b.Discount, b.FinalAmount))
	}
	fmt.Println(summary.String())
}

// PrintCart - prints cart details
func PrintCart(userID uint, products []models.Product) {
	cartDetails := fmt.Sprintf("Cart details %d", userID)

	fmt.Println(cartDetails)
	PrintProducts(products)
}
