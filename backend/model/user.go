package model

type User struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Poll           Poll   `js33on:"poll"`
	SelectedOption string `json:"selectedOption"`
}
