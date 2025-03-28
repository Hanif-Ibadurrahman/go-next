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
	IDUser    int
	Name      string
	Username  string
	Password  string
	Level     int
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin *time.Time
	IsActive  int
	DeletedAt *time.Time
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
	IDUser        int           `json:"id_user"`
	Name          string        `json:"name"`
	Username      string        `json:"username"`
	Level         int           `json:"level"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	LastLogin     *time.Time    `json:"last_login"`
	IsActive      int           `json:"is_active"`
	AccountNumber string        `json:"account_number"`
	Balance       float64       `json:"balance"`
	Pockets       []Pocket      `json:"pockets"`
	TermDeposits  []TermDeposit `json:"term_deposits"`
}
type BankAccount struct {
	ID            int           `json:"id"`
	CustomerID    int           `json:"customer_id"`
	AccountNumber string        `json:"account_number"`
	Balance       float64       `json:"balance"`
	CreatedAt     time.Time     `json:"created_at"`
	Pockets       []Pocket      `json:"pockets"`
	TermDeposits  []TermDeposit `json:"term_deposits"`
}

type Pocket struct {
	ID            int       `json:"id"`
	BankAccountID int       `json:"bank_account_id"`
	PocketName    string    `json:"pocket_name"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}

type TermDeposit struct {
	ID            int       `json:"id"`
	BankAccountID int       `json:"bank_account_id"`
	Amount        float64   `json:"amount"`
	InterestRate  float64   `json:"interest_rate"`
	TermMonths    int       `json:"term_months"`
	StartDate     time.Time `json:"start_date"`
	MaturityDate  time.Time `json:"maturity_date"`
	Status        string    `json:"status"`
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
