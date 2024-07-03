package database

import (
	"context"
	"ecommerce-golang/models"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrorCantFindProduct    = errors.New("cant find the product")
	ErrorCantDecodeProducts = errors.New("Cant find the product")
	ErrorUserIdIsNotValid   = errors.New("User not valid")
	ErrorCantUpdateUser     = errors.New("Cant update user")
	ErrCantRemoveItemCart   = errors.New("Cant remove this item from cart")
	ErrCantGetItem          = errors.New("unable get item from cart")
	ErrCantBuyCartItem      = errors.New("cant buy cart item")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	searchFromDB, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return ErrorCantFindProduct
	}
	var productCart []models.ProductUser

	err = searchFromDB.All(ctx, &productCart)
	if err != nil {
		log.Println(err)
		return ErrorCantDecodeProducts
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrorUserIdIsNotValid
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: productCart}}}}}}
	_, err = userCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return ErrorCantUpdateUser
	}
	return nil

}

func RemoveCartItem(ctx context.Context, userCollection, prodCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrorUserIdIsNotValid
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.M{"$pull": bson.M{"usercart": bson.M{"_id": productID}}}
	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return ErrCantRemoveItemCart
	}
	return nil
}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
