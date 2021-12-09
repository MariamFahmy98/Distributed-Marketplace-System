package controllers

import (
	"net/http"
	"strconv"

	"distributed-marketplace-system/db"
	"distributed-marketplace-system/errors"
	"distributed-marketplace-system/models"
	"distributed-marketplace-system/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct{}

func (ctrl UserController) Signup(c *gin.Context) {
	var input models.SignupInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db.DB.Find(&user, "email = ?", input.Email)

	if user.Email == input.Email {
		c.AbortWithStatusJSON(422, errors.ErrEmailAlreadyRegistered)
		return
	}

	bytePassword := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(422, errors.ErrUnprocessable)
		return
	}

	store := models.Store{Title: input.Name + " Store"}
	db.DB.Create(&store)

	user = models.User{
		Name:     input.Name,
		Email:    input.Email,
		StoreID:  store.ID,
		Password: string(hashedPassword),
	}
	db.DB.Create(&user)

	userId := strconv.FormatInt(user.ID, 10)
	token, _ := util.CreateToken(userId)

	c.IndentedJSON(http.StatusOK, gin.H{"token": token, "user": user.Serialize()})
}

func (ctrl UserController) Login(c *gin.Context) {
	var input models.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := db.DB.First(&user, "email = ?", input.Email)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrNotRegistered)
		return
	}

	bytePassword := []byte(input.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrIncorrectPassword)
		return
	}

	userId := strconv.FormatInt(user.ID, 10)
	token, err := util.CreateToken(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user.Serialize()})
}

func (ctrl UserController) GetAll(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	data := make([]map[string]interface{}, len(users))
	for i, user := range users {
		data[i] = user.PublicSerialize()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func (ctrl UserController) GetOne(c *gin.Context) {
	id := c.Param("id")

	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var user models.User
	result := db.DB.First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": user.PublicSerialize()})
}

func (ctrl UserController) GetProducts(c *gin.Context) {
	id := c.Param("id")

	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var user models.User
	result := db.DB.Preload("Products.User").First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	products := make([]map[string]interface{}, len(user.Products))
	for i, product := range user.Products {
		products[i] = product.Serialize()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": products})
}

func (ctrl UserController) GetSoldProducts(c *gin.Context) {
	id := c.Request.Header.Get("userId")

	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var user models.User
	result := db.DB.First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	var orders []models.Order
	result = db.DB.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()	
	}).Preload("Product.User").Find(&orders, "seller_id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	soldProducts := make([]map[string]interface{}, len(orders))
	for i, order := range orders {
		soldProducts[i] = order.Product.Serialize()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": soldProducts})
}

func (ctrl UserController) GetPurchasedProducts(c *gin.Context) {
	id := c.Request.Header.Get("userId")

	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var user models.User
	result := db.DB.First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	var orders []models.Order
	result = db.DB.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()	
	}).Preload("Product.User").Find(&orders, "buyer_id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	purchasedProducts := make([]map[string]interface{}, len(orders))
	for i, order := range orders {
		purchasedProducts[i] = order.Product.Serialize()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": purchasedProducts})
}

func (ctrl UserController) GetBalance(c *gin.Context) {
	id := c.Request.Header.Get("userId")
	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var user models.User
	result := db.DB.First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": user.Balance})
}

func (ctrl UserController) AddBalance(c *gin.Context) {
	id := c.Request.Header.Get("userId")
	var user models.User
	var input models.DepositInput

	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.First(&user, "id = ?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	db.DB.Create(&models.Transaction{
		UserID:        userId,
		Amount:        input.Amount,
		BalanceBefore: user.Balance,
		Type: 	       "Deposit",
	})

	db.DB.Model(&user).Where("id = ?", userId).Update("balance", input.Amount + user.Balance)

	c.JSON(http.StatusOK, gin.H{"success": true, "balance": user.Balance})
}

func (ctrl UserController) GetReportOnOrders(c *gin.Context) {
	id := c.Request.Header.Get("userId")

	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var user models.User
	result := db.DB.First(&user, "id=?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	var orders []models.Order
	result = db.DB.Joins("Seller").Joins("Buyer").Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()	
	}).Preload("Product.User").Find(&orders, "buyer_id = ? OR seller_id = ?", userId, userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	data := make([]map[string]interface{}, len(orders))
	for i, order := range orders {
		data[i] = order.Serialize()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func (ctrl UserController) GetReportOnTransactions(c *gin.Context) {
	id := c.Request.Header.Get("userId")

	userId, err := strconv.ParseInt(id, 10, 64)
	if userId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrInvalidParameter)
		return
	}

	var user models.User
	result := db.DB.First(&user, "id=?", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	var transactions []models.Transaction
	result = db.DB.Find(&transactions, "user_id", userId)
	if result.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.ErrNotFound)
		return
	}

	data := make([]map[string]interface{}, len(transactions))
	for i, transaction := range transactions {
		data[i] = transaction.Serialize()
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}
