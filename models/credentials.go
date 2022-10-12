package models

import (
	"time"

	"gorm.io/gorm"
)

type UserCredentials struct {
	gorm.Model
	UserID         uint             `json:"userId" gorm:"primary_key"`
	FullName       string           `json:"full_name"`
	TimeRegistered time.Time        `json:"time_registered"`
	Email          string           `json:"email"`
	Credentials    []CredentialsIDs `json:"credentials" gorm:"foreignKey:CredID	"`
}

type CredentialsIDs struct {
	CredID            uint   `json:"credId" gorm:"primary_key"`
	BVN               string `json:"bvn"`
	NationalID        string `json:"national_id"`
	BankAccountNumber string `json:"account_number"`
}
