package models

type RoleData struct {
	Role   string   `json:"role"`
	API    []string `json:"api"`
	Method []string `json:"method"`
}

type Role struct {
	Role string `json:"role"`
}
