package models

type Budget struct {
	ID         uint    `sql:"id"`
	CurrencyID uint    `sql:"currency_id"`
	MaxSum     float64 `sql:"max_sum"`
	MinSum     float64 `sql:"min_sum"`
}
