package modelapp

import "time"

type Users struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `json:"name"`
	Email        string         `gorm:"unique" json:"email"`
	Password     string         `json:"password"`
	Transactions []Transactions `gorm:"foreignKey:UserID"`
}

type Transactions struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	Amount      float64   `gorm:"not null;default:0" json:"amount"`
	Currency    string    `gorm:"not null" json:"currency"`
	Transaction string    `json:"type"`
	Category    string    `json:"category"`
	CreateDate  time.Time `json:"date"`
	Description string    `json:"description"`
	AmountPaid  float64   `gorm:"not null;default:0" json:"amount_paid"`
}

type ExchangeResult struct {
	Currency     string       `json:"currency"`
	Rate         float64      `json:"rate"`
	Amount       float64      `json:"amount"`
	Date         time.Time    `json:"date"`
	Transactions Transactions `json:"transactions"`
}
