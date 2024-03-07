package models

type User struct {
	ID                   uint   `sql:"id" json:"id"`
	NationalityCountryID uint   `sql:"nationality_countryId" json:"nationality_country_id,omitempty"`
	CitizenshipCountryID uint   `sql:"citizenship_countryId" json:"citizenship_country_id,omitempty"`
	GenderID             uint   `sql:"gender_id" json:"genderId,omitempty"`
	BudgetID             uint   `sql:"budget_id" json:"budgetId,omitempty"`
	CitySearchID         uint   `sql:"city_search_id" json:"citySearchId,omitempty"`
	CountrySearchID      uint   `sql:"country_search_id" json:"countrySearchId,omitempty"`
	FullName             string `sql:"full_name" json:"fullName,omitempty"`
	PhoneNumber          string `sql:"phone_number" json:"phoneNumber,omitempty"`
	Email                string `sql:"email" json:"email,omitempty"`
	PasswordHash         string `sql:"password_hash"`
}
