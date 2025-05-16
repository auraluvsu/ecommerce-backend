package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Create new user
type User struct {
	Name string // This should be kinda obvious but its the username
	Pswd string // raw password string that will be hashed with bcrypt
	Id   string // Unique ID for every user
	usrBskt	Basket // Each user will have a dedicated basket, auto set to nil
}

// Create new product
type Product struct {
	Id 		string // Product ID
	Quant	int // Quantity of product. If product count is 0, user will not be able to checkout with the item in their basket
}

// Create new Basket
type Basket struct {
	Items		[]Product // Array of Products. Will be initialised as either an empty array or nil, whichever one is easier
	NetPrice	float64 // Net price before tax will be calculated
	GrossPrice	float64
	BasketUser	User
}