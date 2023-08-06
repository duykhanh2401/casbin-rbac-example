package models

type APIData struct {
	API    string   `json:"api"`
	Role   []string `json:"role"`
	Method []string `json:"method"`
}
