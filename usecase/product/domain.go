package product

type Domain struct {
	ID          uint
	Name        string
	Description string
	Category    string
	Price       int
	Stocks      int
	Discount    int
}

type IProductUsecase interface {
	GetTagihanPLN() (Domain, error)
	GetProduct(id int) ([]Domain, error)
	EditProduct(item Domain) error
	Delete(id int) error
	GetBestSellerCategory(id int) ([]Domain, error)
}

type IProductRepository interface {
	GetTagihanPLN(id int) (Domain, error)
	CountItem(category int) (int, error)
	GetProduct(id int) ([]Domain, error)
	EditProduct(item Domain) error
	Delete(id int) error
	GetBestSellerCategory(id, item int) (Domain, error)
	GetBestSellerCategorySQL(id int) ([]Domain, error)
}
