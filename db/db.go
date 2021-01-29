package db

import (
	"fmt"
	"log"
	"mycart/helpers"
	"mycart/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database - custom wrapper
type Database struct {
	sql *gorm.DB
}

var logger *log.Logger = helpers.GetLoggerInstace()

// CreateConnection - creates and returns a connection
func CreateConnection(dbName string) Database {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", dbName)), &gorm.Config{})
	if err != nil {
		logger.Println("Connection is not setup with ecommerce.db")
		logger.Fatalln(err)
	}

	// Create tables according to the models specified
	db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Cart{}, &models.Invoice{}, &models.User{})
	return Database{db}
}

// InsertRow - inserts row in a table
func (db *Database) InsertRow(table string, value interface{}) error {

	var result *gorm.DB
	switch value.(type) {

	case *models.Category:
		result = db.sql.Create(value.(*models.Category))
	case *models.Product:
		result = db.sql.Create(value.(*models.Product))
	case *models.User:
		result = db.sql.Create(value.(*models.User))
	case *models.Cart:
		result = db.sql.Create(value.(*models.Cart))
	case *models.Invoice:
		result = db.sql.Create(value.(*models.Invoice))
	}
	if result.Error != nil {
		return result.Error
	}
	logger.Println("Successfully added entry to", table)
	return nil
}

// Delete - delete categories or products
func (db *Database) Delete(model interface{}, ids []uint) error {

	res := db.sql.Where("id IN ?", ids).Delete(model)
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Affected rows:", res.RowsAffected)
	return nil
}

// DeleteFromCart - Deletes products from the cart using their ids
func (db *Database) DeleteFromCart(ids []uint) error {
	res := db.sql.Where("product_id IN ?", ids).Delete(&models.Cart{})
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Deleted products with ids", ids)
	return nil
}

// SaveBill - saves to bill to the table
func (db *Database) SaveBill(bill *models.Invoice) error {
	res := db.sql.Create(bill)
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Saved bill")
	return nil
}

// UpdateCart - add product to cart
func (db *Database) UpdateCart(cart *models.Cart) error {
	res := db.sql.Create(cart)
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Updated cart products")
	return nil
}

// DeleteFromUserCart - deletes user from the cart table
func (db *Database) DeleteFromUserCart(userID uint) error {
	res := db.sql.Where("user_id=?", userID).Delete(&models.Cart{})
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Removed User's cart")
	return nil
}

// GetUserDetails - Get user information.
func (db *Database) GetUserDetails(id uint, user *models.User) {
	res := db.sql.Where("id=?", id).Find(&user)
	if res.Error != nil {
		logger.Fatalln("Failed to fetch user details for id:", id, "\n", res.Error)
	}
	logger.Println("Fetched user details for id:", id)
}

// GetBills - gets all bills from table
func (db *Database) GetBills(bills *[]models.Invoice) {
	res := db.sql.Find(bills)
	if res.Error != nil {
		logger.Fatalln("Failed to get bills \n", res.Error)
	}
	logger.Println("Fetched bills")
}

// FetchCartProducts - fetch prodcuts in cart for a user
func (db *Database) FetchCartProducts(userID uint, products *[]models.Product) {
	cartProducts := []models.Cart{}
	res := db.sql.Where("user_id=?", userID).Find(&cartProducts)
	if res.Error != nil {
		logger.Fatalln("Failed to get cart products \n", res.Error)
	}
	prods := []models.Product{}
	for _, cp := range cartProducts {
		p := models.Product{}
		db.FetchProductDetails(cp.ProductID, &p)
		prods = append(prods, p)
	}
	*products = prods
	logger.Println("Fetched cart products")
}

// FetchProducts - fetch products for given ids
func (db *Database) FetchProducts(pids []uint, products *[]models.Product) {
	res := db.sql.Where("id IN ?", pids).Find(&products)
	if res.Error != nil {
		logger.Fatalln("Failed to get products \n", res.Error)
	}
	logger.Println("Fetched products")
}

// FetchAllCategories - fetch all the categories available
func (db *Database) FetchAllCategories(categories *[]models.Category) {
	res := db.sql.Find(&categories)
	if res.Error != nil {
		logger.Fatalln("Failed to fetch categories:\n", res.Error)
	}
	logger.Println("Fetched categories from `categories` table")
}

// FetchProductsOfCategory - fetch all the products ina given category
func (db *Database) FetchProductsOfCategory(categoryName string, products *[]models.Product) {
	res := db.sql.Find(&products, map[string]interface{}{"category_name": categoryName})
	if res.Error != nil {
		logger.Fatalln("Failed to fetch products for categoryName: ", categoryName, "\n", res.Error)
	}
	logger.Println("Fetched products from category", categoryName)
}

// FetchProductDetails - fetch the product details of a product
func (db *Database) FetchProductDetails(productID uint, product *models.Product) {
	res := db.sql.First(&product, map[string]interface{}{"id": productID})
	if res.Error != nil {
		logger.Fatalln("Failed to fetch product details for productID: ", productID, "\n", res.Error)
	}
	logger.Println("Product Details fetched for id:", productID)
}

// FetchCartDetails - fetches cart details of a user. userID is used to identify cart.
func (db *Database) FetchCartDetails(userID uint, cart *models.Cart) {
	res := db.sql.Where("user_id=?", userID).Find(&cart)
	if res.Error != nil {
		logger.Fatalln("Failed to cart details for a user \n", res.Error)
	}
	logger.Println("Fetched cart details for user:", userID)
}

// FetchParticularProducts - will fetch products according to their ids from the cart
func (db *Database) FetchParticularProducts(userID uint, pids []uint, cartProducts *[]models.Product) {

	cps := []models.Cart{}
	res := db.sql.Where("user_id=? AND product_id IN ?", userID, pids).Find(&cps)
	if res.Error != nil {
		logger.Fatalln("Failed to get cart products for productIDS", pids, "\n", res.Error)
	}
	prods := []models.Product{}
	for _, cp := range cps {
		p := models.Product{}
		db.FetchProductDetails(cp.ProductID, &p)
		prods = append(prods, p)
	}
	*cartProducts = prods
	logger.Println("Fetched products with ids", pids)
}
