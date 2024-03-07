package models

type Country struct {
	ID         uint   `sql:"id"`
	Name       string `sql:"name"`
	FlagImgUrl string `sql:"flag_img_url"`
}
