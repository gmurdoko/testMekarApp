package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DateBirth    string `json:"date_birth"`
	IdCardNumber string `json:"id_card_number"`
	Profession   string `json:"profession"`
	Education    string `json:"education"`
	IdProfession string `json:"id_profession"`
	IdEducation  string `json:"id_education"`
}

// type UserInput {
// 	Name         string `json:"name"`
// 	DateBirth    string `json:"date_birth"`
// 	IdCardNumber string `json:"id_card_number"`
// 	IdProfession   string `json:"profession"`
// 	IdEducation    string `json:"education"`
// }
