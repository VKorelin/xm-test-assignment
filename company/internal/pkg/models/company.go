package models

import "github.com/gofrs/uuid"

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
	AmountOfEmployees uint32
	Registered        bool
	Type              CompanyType
}
