package models

type Language struct {
	// 言語
	BaseModel
	Name string `json:"name"`
	Code int    `json:"code"`
}

func (Language) TableName() string {
	return "language"
}
