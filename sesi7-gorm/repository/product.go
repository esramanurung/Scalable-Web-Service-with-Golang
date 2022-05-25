package repository
import (
	"sesi7-gorm/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(product *models.Product) error
	GetAllProduct() (*[]models.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) CreateProduct(product *models.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepo) GetAllProduct() (*[]models.Product, error) {
	var products []models.Product
	err := p.db.Find(&products).Error
	return &products, err
}