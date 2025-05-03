package controllers

import (
	"math"
	"net/http"
	"os"

	"github.com/google/uuid"
	paymentModels "github.com/iamsabbiralam/stripe-with-go/pkg/payment/models"
	paymentRepositories "github.com/iamsabbiralam/stripe-with-go/pkg/payment/repositories"

	utils "github.com/iamsabbiralam/stripe-with-go/utils"

	stripe "github.com/stripe/stripe-go/v78"
	paymentIntent78 "github.com/stripe/stripe-go/v78/paymentintent"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PaymentControllers interface{
	Createproduct(ctx *gin.Context)
}

// paymentController struct implememts paymentController interface.
type PaymentController struct {
	paymentRepository paymentRepositories.PaymentService
	validator         *validator.Validate
	db                *gorm.DB
}

func NewPaymentController(db *gorm.DB) *PaymentController {
	return &PaymentController{
		paymentRepository: paymentRepositories.NewPaymentRepository(db),
		validator:         validator.New(),
		db:                db,
	}
}

func (pc *PaymentController) CreateProduct(ctx *gin.Context) {
	var orderRequest paymentModels.Payment
	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(
			http.StatusBadRequest,
			"No Body Request. Please provide the fields",
			err,
		))
		return
	}

	email := os.Getenv("CUSTOMER_EMAIL")
	totalAmount := orderRequest.Amount
	if totalAmount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment amount"})
		return
	}

	amountInCents := int64(math.Round(totalAmount * 100))
	getCustomer, searchErr := pc.paymentRepository.SearchCustomerOnStripe(ctx, email)
	if searchErr.Message != "" {
		ctx.JSON(searchErr.Status, gin.H{"error": searchErr.Message})
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountInCents),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Customer:     stripe.String(getCustomer),
		ReceiptEmail: stripe.String(email),
		Metadata: map[string]string{
			"user_id": "1",
			"city":    "Khulna",
			"address": "Khulna",
		},
	}

	pi, err := paymentIntent78.New(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	invoice, err := pc.paymentRepository.ConfirmPaymentIntent(pi.ID, email, getCustomer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderID := uuid.New()
	order := &paymentModels.Payment{
		ID:     orderID,
		Amount: totalAmount,
	}

	if err := pc.paymentRepository.CreateOne(order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"invoice": invoice,
		"message": "✅ Order created successfully after payment!",
	})
}
