package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customer Model
type Customer struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	Fullname    string    `gorm:"type:varchar(255);not null"`
	Username    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password    string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"type:varchar(13);not null"`
	Address     string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Orders []Order `gorm:"foreignKey:CustomerID"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}

// Barista Model
type Barista struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	Fullname  string    `gorm:"type:varchar(255);not null"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (b *Barista) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

// Category Model
type Category struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Products []Product `gorm:"foreignKey:CategoryID"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}

// Product Model
type Product struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Slug        string    `gorm:"type:varchar(255);uniqueIndex;not-null"`
	CategoryID  string    `gorm:"type:char(36);not null"`
	Description string    `gorm:"type:text;not null"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Image       string    `gorm:"type:varchar(255)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

type Cart struct {
	ID         string    `gorm:"type:char(36);primaryKey"`
	CustomerID string    `gorm:"type:char(36);not null"`
	ProductID  string    `gorm:"type:char(36);not null"`
	Quantity   int       `gorm:"type:int;not null"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Customer Customer `gorm:"foreignKey:CustomerID;references:ID;constraint:OnDelete:CASCADE"`
	Product  Product  `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}

func (p *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

// OrderStatus Enum
type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Processed OrderStatus = "Processed"
	Completed OrderStatus = "Completed"
	Cancelled OrderStatus = "Cancelled"
)

// Order Model
type Order struct {
	ID            string      `gorm:"type:char(36);primaryKey"`
	OrderCode     string      `gorm:"text;not null"`
	CustomerID    string      `gorm:"type:char(36);not null"`
	Address       string      `gorm:"text;not null"`
	PhoneNumber   string      `gorm:"char(13);not null"`
	PaymentMethod string      `gorm:"type:varchar(255)"`
	TotalPrice    float64     `gorm:"type:decimal(10,2);not null"`
	Status        OrderStatus `gorm:"type:varchar(20);not null"`
	AdminFee      float64     `gorm:"type:decimal(10,2);not null"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"`

	Customer     Customer      `gorm:"foreignKey:CustomerID;references:ID"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

// OrderDetail Model
type OrderDetail struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	OrderID   string    `gorm:"type:char(36);not null"`
	ProductID string    `gorm:"type:char(36);not null"`
	Quantity  int       `gorm:"type:int;not null"`
	SubTotal  float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Order   Order   `gorm:"foreignKey:OrderID;references:ID"`
	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}

func (od *OrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	od.ID = uuid.New().String()
	return
}
