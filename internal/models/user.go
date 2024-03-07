package models

type User struct {
	ID                   uint   `sql:"id"`
	NationalityCountryID uint   `sql:"nationality_country_id"`
	CitizenshipCountryID uint   `sql:"citizenship_country_id"`
	GenderID             uint   `sql:"gender_id"`
	BudgetID             uint   `sql:"budget_id"`
	CitySearchID         uint   `sql:"city_search_id"`
	CountrySearchID      uint   `sql:"country_search_id"`
	FullName             string `sql:"full_name"`
	PhoneNumber          string `sql:"phone_number"`
	Email                string `sql:"email"`
	PasswordHash         string `sql:"password_hash"`
}
