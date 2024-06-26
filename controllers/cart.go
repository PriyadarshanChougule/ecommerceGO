package controllers

import (
	"context"
	"ecommerce-golang/database"
	"ecommerce-golang/models"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {

	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ProductQueryID := c.Query("id")
		if ProductQueryID == "" {
			log.Println("Product id is ")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(ProductQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "successfully added to cart")

	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ProductQueryID := c.Query("id")
		if ProductQueryID == "" {
			log.Println("Product id is ")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(ProductQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "successfully removed item from cart")

	}
}

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var filledCart models.User
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: user_id}}).Decode(&filledCart)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(404, "not found")
			return
		}

		filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{
			primitive.E{Key: "_id", Value: "$_id"},
			{Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: usert_id}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter, unwind, group})

		if err != nil {
			log.Println(err)
		}

		var list []bson.M
		if err = pointCursor.All(ctx, &list); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		for _, json := range list {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledCart.UserCart)
		}

		ctx.Done()
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			log.Panic("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userID is empty."))
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "successfully placed the order")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ProductQueryID := c.Query("id")
		if ProductQueryID == "" {
			log.Println("Product id is ")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product id is empty"))
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(ProductQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "successfully Placed the order")
	}
}
