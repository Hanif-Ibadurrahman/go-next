package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type QuerySearch struct {
	Q string `json:"q"`
}

type RequestCreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
}

type RequestUpdateUser struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
}

type UpdateUser struct {
	Name     *string
	Password *string
	IDUser   int
}

type ResponseCreateUser struct {
	IDUser   int    `json:"id_user"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Level    int    `json:"level"`
}

type User struct {
	IDUser       int
	Name         string
	Username     string
	Password     string
	Level        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastLogin    *time.Time
	IsActive     int
	DeletedAt    *time.Time
	BankAccounts []BankAccount
}

type FlatUser struct {
	IDUser        int
	Name          string
	Username      string
	Level         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	LastLogin     *time.Time
	IsActive      int
	AccountNumber string
	Balance       float64
}

type UserDetail struct {
	IDUser        int
	Name          string
	Username      string
	Level         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	LastLogin     *time.Time
	IsActive      int
	AccountNumber string
	Balance       float64
	Pockets       []Pocket
	TermDeposits  []TermDeposit
}
type BankAccount struct {
	ID            int
	CustomerID    int
	AccountNumber string
	Balance       float64
	CreatedAt     time.Time
	Pockets       []Pocket
	TermDeposits  []TermDeposit
}

type Pocket struct {
	ID            int
	BankAccountID int
	PocketName    string
	Balance       float64
	CreatedAt     time.Time
}

type TermDeposit struct {
	ID            int
	BankAccountID int
	Amount        float64
	InterestRate  float64
	TermMonths    int
	StartDate     time.Time
	MaturityDate  time.Time
	Status        string
}

func (v *QuerySearch) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.Q, validation.Length(0, 100), is.Alphanumeric),
	)
}

func (v *RequestCreateUser) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.Username, validation.Required),
		validation.Field(&v.Password, validation.Required),
		validation.Field(&v.Name, validation.Required),
		validation.Field(&v.Level, validation.Required),
	)
}

func (v *RequestUpdateUser) Validate() error {
	return validation.ValidateStruct(v,
		validation.Field(&v.Password),
		validation.Field(&v.Name),
	)
}
