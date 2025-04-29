package config

import (
	"os"

	"github.com/stripe/stripe-go/v78"
)

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}
