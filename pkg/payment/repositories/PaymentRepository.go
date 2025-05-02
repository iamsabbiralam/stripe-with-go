package repositories

import (
	"context"
	"fmt"
	"os"
	"time"

	paymentModels "github.com/iamsabbiralam/stripe-with-go/pkg/payment/models"

	striped "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	paymentIntent72 "github.com/stripe/stripe-go/v72/paymentintent"
	stripe "github.com/stripe/stripe-go/v78"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	collection *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		collection: db.Table("payments"),
	}
}

func (pr *PaymentRepository) newSession(ctx context.Context) *gorm.DB {
	return pr.collection.Session(&gorm.Session{}).WithContext(ctx)
}

func (pr *PaymentRepository) CreateOne(payment *paymentModels.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db := pr.newSession(ctx)
	if err := db.Create(payment).Error; err != nil {
		return err
	}
	
	return nil
}

func (pr *PaymentRepository) SearchCustomerOnStripe(ctx context.Context, email string) (string, error) {
	striped.Key = os.Getenv("STRIPE_SECRET_KEY")
	if striped.Key == "" {
		return "", fmt.Errorf("API key is missing")
	}

	params := &striped.CustomerListParams{
		Email: striped.String(email),
	}

	iter := customer.List(params)
	for iter.Next() {
		c := iter.Customer()
		return c.ID, nil
	}

	if err := iter.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("Customer not found with email: %s", email)
}

func (pc *PaymentRepository) ConfirmPaymentIntent(paymentIntent, email, customerID string) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		return "", fmt.Errorf("API key is missing")
	}
	paymentMethod, err := pc.GetPaymentMethod(customerID)
	if err != nil {
		return "", err
	}

	if paymentMethod == "" {
		return "", fmt.Errorf("payment method is missing")
	}

	params := &striped.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String(paymentMethod),
		ReturnURL:     stripe.String("https://www.payment.com"),
		ReceiptEmail:  stripe.String(email),
	}
	payment, err := paymentIntent72.Confirm(paymentIntent, params)
	if err != nil {
		return "", err
	}

	return payment.Charges.Data[0].ReceiptURL, nil
}

func (pr *PaymentRepository) GetPaymentMethod(customerID string) (string, error) {
	striped.Key = os.Getenv("STRIPE_SECRET_KEY")
	if striped.Key == "" {
		return "", fmt.Errorf("API key is missing")
	}

	params := &striped.CustomerListPaymentMethodsParams{
		Customer: striped.String(customerID),
	}
	params.Limit = striped.Int64(3)
	result := customer.ListPaymentMethods(params)
	if len(result.PaymentMethodList().Data) == 0 {
		return "", fmt.Errorf("no payment method found")
	}

	return result.PaymentMethodList().Data[0].ID, nil
}
