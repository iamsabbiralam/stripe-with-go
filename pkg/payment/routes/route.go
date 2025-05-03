package routes

import (
	paymentControllers "github.com/iamsabbiralam/stripe-with-go/pkg/payment/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentRoutes(router *gin.RouterGroup, db *gorm.DB) {
	paymentController := paymentControllers.NewPaymentController(db)
	paymentRoutes := router.Group("/payment")
	paymentRoutes.POST("/create", paymentController.CreateProduct)
}
