package cmd

import (
	"fmt"
	"mycart/db"
	"mycart/helpers"
	"mycart/models"
	"strconv"

	"github.com/spf13/cobra"
)

var sqlDb db.Database = db.CreateConnection("ecommerce")
var user models.User = helpers.GetUser()

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Use it to add users, products, categories",
	Long: `With the help of flags you will be able to add either users or products or categories.
	To add products and categories user must be an admin.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if user.IsAdmin {
			isProd, err := cmd.Flags().GetBool("product")
			if err != nil {
				logger.Fatalln("Failed to get `product` flag \n", err)
			}
			// Add product
			if isProd {
				product := models.Product{}
				fmt.Println(args, len(args))
				if len(args) != 3 {
					logger.Fatalln("Not enough arguments passed.")
				}
				// category := models.Category{Name: args[0]}
				product.CategoryName = args[0]
				product.Name = args[1]
				price, err := strconv.Atoi(args[2])
				if err != nil {
					logger.Fatalln("Failed to convert price to int \n", err)
				}
				product.Price = float64(price)
				err = sqlDb.InsertRow("products", &product)
				if err != nil {
					logger.Println("Failed to Write to the table")
					logger.Fatalln(err)
				}
			}
			isCat, err := cmd.Flags().GetBool("category")
			if err != nil {
				logger.Fatalln("Failed to get `category` flag \n", err)
			}
			// Add category
			if isCat {
				category := models.Category{}
				if len(args) != 1 {
					logger.Fatalln("Not enough arguments passed")
				}
				category.Name = args[0]
				sqlDb.InsertRow("categories", &category)
			}
		}

		isUser, err := cmd.Flags().GetBool("user")
		if err != nil {
			logger.Fatalln("Failed to get `user` flag \n", err)
		}
		// Add normal user
		if isUser {
			user := models.User{}
			user.Name = args[0]
			sqlDb.InsertRow("users", &user)
		}

		// Add admin user
		isAdmin, err := cmd.Flags().GetBool("admin")
		if err != nil {
			logger.Fatalln("Failed to get `admin` flag \n", err)
		}
		if isAdmin {
			user := models.User{}
			user.Name = args[0]
			user.IsAdmin = true
			sqlDb.InsertRow("users", &user)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().Bool("user", false, "add --user Username")
	addCmd.Flags().Bool("admin", false, "add --admin Username")

	// Admin level
	addCmd.Flags().Bool("product", false, "add --product CategoryName ProductName Price")
	addCmd.Flags().Bool("category", false, "add --category CategoryName")
}
