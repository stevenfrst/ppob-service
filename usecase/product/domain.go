package product

type Domain struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	Category    string
	Price       int
	Stocks      int
	Tax         int
	SubCategory string
	Link        string
}

type CreateDomain struct {
	ID            uint `gorm:"primarykey"`
	Name          string
	Description   string
	CategoryID    uint
	Price         int
	Stocks        int
	SubCategoryID uint
}

type Category struct {
	ID   uint
	Name string
}

type SubCategory struct {
	ID       uint
	Name     string
	Tax      int
	ImageURL string
}

type IProductUsecase interface {
	GetTagihanPLN() (Domain, error)
	GetProduct(id int) ([]Domain, error)
	EditProduct(item Domain) error
	Delete(id int) error
	GetBestSellerCategory(id int) ([]Domain, error)
	Create(domain CreateDomain) error
	GetAll(offset, pageSize int) ([]Domain, error)
	GetAllCategory() []Category
	GetAllSubCategory() []SubCategory
	EditSubCategory(edit SubCategory) error
}

type IProductRepository interface {
	GetTagihanPLN(id int) (Domain, error)
	CountItem(category int) (int, error)
	GetProduct(id int) ([]Domain, error)
	EditProduct(item Domain) error
	Delete(id int) error
	GetBestSellerCategory(id, item int) (Domain, error)
	GetBestSellerCategorySQL(id int) ([]Domain, error)
	Create(input CreateDomain) error
	GetAllProduct() ([]Domain, error)
	GetAllProductPagination(offset, pageSize int) ([]Domain, error)
	GetAllCategory() []Category
	GetAllSubCategory() []SubCategory
	EditSubCategory(sub SubCategory) error
}
