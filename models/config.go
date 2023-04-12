package models

type Config struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"NAME"`
	Value string `json:"value" db:"VALUE"`
}
