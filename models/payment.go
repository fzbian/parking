package models

import "github.com/shopspring/decimal"

type Payment struct {
	ID          int64           `db:"id"`
	VehicleID   int64           `db:"vehicle_id"`
	PaymentType string          `db:"payment_type"`
	Amount      decimal.Decimal `db:"amount"`
}
