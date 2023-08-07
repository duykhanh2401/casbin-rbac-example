package models

type CategoryData struct {
	Name   string   `json:"category"`
	Role   []string `json:"role"`
	Method []string `json:"method"`
}
