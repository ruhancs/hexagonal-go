package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// interface para definir um produto
type ProductInterface interface {
	IsValid() (bool, error)//retorna um bool ou error
	Enable() error// pode retornar um error ou vazio
	Disable() error // para desabilitar um produto pode retornar um erro ou vazio
	GetID() string
	GetName()string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLE = "enabled"
)

type Product struct {
	ID string `valid:"uuidv4"` // informa que id deve ser um uuid
	Name string `valid:"required"`
	Price float64 `valid:"float,optional"`
	Status string `valid:"required"`
}

func NewProduct() *Product {
	product := Product {
		ID: uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if(p.Status == "") {
		p.Status = DISABLED
	}
	
	if(p.Status != ENABLE && p.Status != DISABLED) {
		return false, errors.New("the status must be enabled or disabled")
	}	

	if(p.Price < 0) {
		return false, errors.New("the price must be greather or equal to 0")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if(p.Price > 0) {
		p.Status = ENABLE
		return nil // indica que nao retornou um erro
	}
	return errors.New("The must be greather tha 0")
}

func (p *Product) Disable() error {
	if(p.Price == 0) {
		p.Status = DISABLED
		return nil
	}
	return errors.New("The must be 0 to disable the product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

