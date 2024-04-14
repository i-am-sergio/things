package models

type Role string

const (
	RoleAdmin      Role = "ADMIN"
	RoleUser       Role = "USER"
	RoleEnterprise Role = "ENTERPRISE"
)
