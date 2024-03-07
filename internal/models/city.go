package models

type City struct {
	ID         uint   `sql:"id"`
	CountryID  uint   `sql:"country_id"`
	Name       string `sql:"name"`
	FlagImgUrl string `sql:"flag_img_url"`
}
