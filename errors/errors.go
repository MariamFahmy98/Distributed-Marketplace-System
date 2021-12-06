package errors

import (
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = gin.H{"error": "404 Not Found"}

	ErrBadRequest = gin.H{"error": "400 Bad Request"}

	ErrUserNotFound = gin.H{"error": "User Not Found"}

	ErrProductNotFound = gin.H{"error": "Product Not Found"}

	ErrStoreNotFound = gin.H{"error": "Store Not Found"}

	ErrInvalidParameter = gin.H{"error": "Missing or Invalid Parameter"}

	ErrUnauthorized = gin.H{"error": "You are not authorized to perform this action"}

	ErrNotRegistered = gin.H{"error": "You have not registered, please signup first"}

	ErrIncorrectPassword = gin.H{"error": "Incorrect Password"}

	ErrInvalidToken = gin.H{"error": "Invalid or expired token"}

	ErrEmailExists = gin.H{"error": "Email is Already Exists"}

	ErrUnprocessable = gin.H{"error": "422 Unprocessable"}

	ErrEmailAlreadyRegistered = gin.H{"error": "Email is Already Registered"}

	ErrBalanceNotEnough = gin.H{"error": "Balance is not enough"}

	ErrNotForSales = gin.H{"error": "This product is not for sales"}

  ErrCannotBuyYourProduct = gin.H{"error": "You cannot buy your own product"}
)
