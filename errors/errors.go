package errors

import (
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound               = gin.H{"error": "404 Not Found"}
	ErrInvalidParameter       = gin.H{"error": "Invalid Parameter"}
	ErrUnauthorized           = gin.H{"error": "You are not authorized to perform this action"}
	ErrNotRegistered          = gin.H{"error": "You have not registered, please signup first"}
	ErrIncorrectPassword      = gin.H{"error": "Incorrect Password"}
	ErrInvalidToken           = gin.H{"error": "Invalid or expired token"}
	ErrEmailAlreadyRegistered = gin.H{"error": "Email is Already Registered"}
	ErrUnprocessable          = gin.H{"error": "422 Unprocessable"}
)
