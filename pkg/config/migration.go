package config

import (
	"database/sql/driver"

	paymentModels "github.com/iamsabbiralam/stripe-with-go/pkg/payment/models"

	"github.com/lib/pq"
)

type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	return pq.StringArray(s).Value()
}

func (s *StringArray) Scan(src interface{}) error {
	return (*pq.StringArray)(s).Scan(src)
}

type Payment struct {
	paymentModels.Payment
}
