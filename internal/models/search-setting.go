package models

import "time"

type SearchSetting struct {
	ID            uint64    `sql:"id"`
	UserID        uint64    `sql:"user_id"`
	NationalityID uint64    `sql:"nationality_id"`
	CountryID     uint64    `sql:"country_id"`
	CityID        uint64    `sql:"city_id"`
	CurrencyID    uint64    `sql:"currency_id"`
	MinSum        uint32    `sql:"min_sum"`
	MaxSum        uint32    `sql:"max_sum"`
	Gender        bool      `sql:"gender"`
	CreateDate    time.Time `sql:"create_date"`
	UpdateDate    time.Time `sql:"update_date"`
}
