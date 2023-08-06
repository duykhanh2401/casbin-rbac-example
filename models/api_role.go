package models

type APIRole struct {
	API  string   `json:"api"`
	Role []string `json:"role"`
}
