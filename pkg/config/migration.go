package config

import (
	"database/sql/driver"

	paymentModels "github.com/iamsabbiralam/stripe-with-go/pkg/payment/models"

	"github.com/lib/pq"
)

// StringArray is a custom type to handle []string as JSONB
type StringArray []string

// Value marshals the array into a JSONB value for the database
func (s StringArray) Value() (driver.Value, error) {
	return pq.StringArray(s).Value()
}

// Scan implements the sql.Scanner interface for database deserialization
func (s *StringArray) Scan(src interface{}) error {
	return (*pq.StringArray)(s).Scan(src)
}

type Payment struct {
	paymentModels.Payment
}
