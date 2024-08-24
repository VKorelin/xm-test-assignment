package models

import "github.com/google/uuid"

type CompanyType int

const (
	Corporations CompanyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
)

type Company struct {
	Id                uuid.UUID
	Name              string
	Description       string
	AmountOfEmployees int
	Registered        bool
	Type              CompanyType
}
