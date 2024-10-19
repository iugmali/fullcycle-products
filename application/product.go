package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() int64
	SetPrice(price int64) error
}

type ProductServiceInterface interface {
	GetAll() ([]ProductInterface, error)
	Get(id string) (ProductInterface, error)
	SetPrice(product ProductInterface, price int64) (ProductInterface, error)
	Create(name string, price int64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	GetAll() ([]ProductInterface, error)
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string `valid:"uuidv4" json:"id"`
	Name   string `valid:"required" json:"name"`
	Price  int64  `valid:"int,optional" json:"price"`
	Status string `valid:"required" json:"status"`
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.NewString(),
		Status: DISABLED,
	}
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}
	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("the status must be enabled or disabled")
	}
	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price must be zero")
}

func (p *Product) GetID() string {
	return p.ID
}
func (p *Product) GetName() string {
	return p.Name
}
func (p *Product) SetName(name string) error {
	p.Name = name
	_, err := p.IsValid()
	if err != nil {
		return err
	}
	return nil
}
func (p *Product) GetStatus() string {
	return p.Status
}
func (p *Product) GetPrice() int64 {
	return p.Price
}
func (p *Product) SetPrice(price int64) error {
	p.Price = price
	_, err := p.IsValid()
	if err != nil {
		return err
	}
	return nil
}
