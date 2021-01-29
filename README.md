# CLI Application

A simple Ecommerce CLI app

## Functionalities

### For Users

- Listing of categories
- Listing of products in a specific category.
- Viewing a product's details.
- Add product to cart.
- Buying products from the cart.
- Simple Discount calculation.
- Bill display.

### For Admin

- Adding new categories and products.
- Viewing details of user cart.
- Display of all the bills summary generated .

## Installation

- Go version greater than 1.13
- Clone the app
- Navigate to the root directory of the app.
- Execute 
  - `go mod download` 
  - `go install`
- Check if app successfully installed by `mycart --help`

## Examples

- See help `mycart --help`

- Add user and admin

    ```
    // To add a normal user
    mycart add --user username

    // To add an admin user
    mycart add --admin username

    ```

- Activate a user (login with a user) (Functionalities will be restricted if user is not activated)

    ```
    mycart activate --user userID
    ```

- Add Product and Categories (Admin Restricted)

  ```
  // Add product
  mycart add --product CategoryName ProductName Price

  // Add Category
  mycart add --category CategoryName
  ```

- List Categories

    ```
    mycart list --categories
    ```

- List Products of a specific Category

    ```
    mycart list --products [--categoryName | -c ] CatName
    ```

- List Items in a cart

    ```
    mycart list --cart-items
    ```
- List Bills Summary (Admin)

    ```
    mycart list --bills
    ```

- Delete Products (Admin)

    ```
    mycart delete --products ProductID1,ProductID2
    ```

- Delete Categories (Admin)

    ```
    mycart delete --categories CategoryID1,CategoryID2
    ```

- Add Product to Cart

    ```
      mycart update-cart --add-product ProductID
    ```

- Remove Products from Cart

    ```
      mycart update-cart --remove-products ProductID1,ProductID2,ProductID3
    ```

- Buy products - ProductID's must be from the cart only. Will generate Bill.

    ```
    mycart checkout
    ```


## Miscellenous

- **credentials.json** - Saves the current activated user
- **ecommerce.db** - is the SQLite db file
