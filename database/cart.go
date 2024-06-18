package database

import "errors"

var (
	ErrorCantFindProduct    = errors.New("cant find the product")
	ErrorCantDecodeProducts = errors.New("Cant find the product")
	ErrorUserIdIsNotValid   = errors.New("User not valid")
	ErrorCantUpdateUser     = errors.New("Cant update user")
	ErrCantRemoveItemCart   = errors.New("Cant remove this item from cart")
	ErrCantGetItem          = errors.New("unable get item from cart")
	ErrCantBuyCartItem      = errors.New("cant buy cart item")
)

func AddProductToCart() {

}
func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
