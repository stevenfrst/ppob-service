package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"math"
	"net/http"
	"ppob-service/app/config"
	_middleware "ppob-service/app/middleware"
	"ppob-service/app/routes"
	productDelivery "ppob-service/delivery/product"
	txDelivery "ppob-service/delivery/transaction"
	userDelivery "ppob-service/delivery/user"
	voucherDelivery "ppob-service/delivery/voucher"
	"ppob-service/drivers/email"
	payment "ppob-service/drivers/midtrans"
	"ppob-service/drivers/mysql"
	cache "ppob-service/drivers/redis"
	productRepo "ppob-service/drivers/repository/product"
	transactionRepo "ppob-service/drivers/repository/transaction"
	txRepository "ppob-service/drivers/repository/transaction"
	userRepo "ppob-service/drivers/repository/user"
	voucherRepo "ppob-service/drivers/repository/voucher"
	storagedriver "ppob-service/drivers/s3"
	"ppob-service/helpers/encrypt"
	productUsecase "ppob-service/usecase/product"
	txUsecase "ppob-service/usecase/transaction"
	userUsecase "ppob-service/usecase/user"
	voucherUsecase "ppob-service/usecase/voucher"
	"time"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func encryptMigration(password string) string {
	out, _ := encrypt.Hash(password)
	return out
}

func dbMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&voucherRepo.Voucher{}, &productRepo.Category{}, &productRepo.SubCategory{}, &productRepo.Product{}, &userRepo.User{}, &transactionRepo.Transaction{})
	if err != nil {
		log.Fatalln(err)
	}
	dummyParse, _ := time.Parse(time.RFC3339, "2022-01-05T00:00:00.00+07:00")
	log.Println(dummyParse)
	var voucher = voucherRepo.Voucher{
		ID:    1,
		Code:  "GESEKGESEK",
		Value: 10000,
		Valid: dummyParse,
	}
	db.Create(&voucher)
	var users = []userRepo.User{{ID: 1, Role: "admin", Username: "admin", Password: encryptMigration("admin"), Email: "admin@admin.com", PhoneNumber: "082135166117"},
		{ID: 2, Role: "user", Username: "kuli", Password: "kuli", Email: "kuli@user.com", PhoneNumber: "0821313123"},
		{ID: 3, Role: "user", Username: "kuli2", Password: "kuli2", Email: "kuli2@user.com", PhoneNumber: "0831231299"},
	}
	db.Create(&users)
	var category = []productRepo.Category{{ID: 1, Name: "Pulsa"},
		{ID: 2, Name: "Voucher Restoran"},
		{ID: 3, Name: "Tagihan PLN"},
	}
	db.Create(&category)
	var subcategory = []productRepo.SubCategory{
		{
			ID:       1,
			Name:     "Telkomsel",
			Tax:      1000,
			ImageURL: "",
		},
		{
			ID:       2,
			Name:     "Indosat",
			Tax:      1000,
			ImageURL: "",
		},
		{
			ID:       3,
			Name:     "Tri",
			Tax:      1000,
			ImageURL: "",
		},
		{
			ID:       4,
			Name:     "Xl",
			Tax:      1000,
			ImageURL: "",
		},
		{
			ID:       5,
			Name:     "KFC",
			Tax:      1000,
			ImageURL: "",
		},
		{
			ID:       6,
			Name:     "ANU",
			Tax:      1000,
			ImageURL: "",
		},
		{
			ID:       7,
			Name:     "PLN Prabayar",
			Tax:      2500,
			ImageURL: "",
		},
		{
			ID:       8,
			Name:     "PLN Token",
			Tax:      1000,
			ImageURL: "",
		},
	}
	db.Create(&subcategory)
	var products = []productRepo.Product{{ID: 1, Name: "Pulsa Telkomsel 10K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 10000, Stocks: 50, Sold: 6, SubCategoryID: 1},
		{ID: 2, Name: "Pulsa Telkomsel 20K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 20000, Stocks: 50, Sold: 3, SubCategoryID: 1},
		{ID: 3, Name: "Pulsa Telkomsel 25K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 25000, Stocks: 50, Sold: 10, SubCategoryID: 1},
		{ID: 4, Name: "Pulsa Telkomsel 50K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 50000, Stocks: 50, Sold: 99, SubCategoryID: 1},
		{ID: 5, Name: "Pulsa Telkomsel 100K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 100000, Stocks: 50, Sold: 1, SubCategoryID: 1},
		{ID: 6, Name: "Voucher KFC 100K", Description: "Voucher Restoran", CategoryID: 2, Price: 100000, Stocks: 50, Sold: 123, SubCategoryID: 5},
		{ID: 7, Name: "Voucher KFC 200K", Description: "Voucher Restoran", CategoryID: 2, Price: 200000, Stocks: 50, Sold: 90, SubCategoryID: 5},
		{ID: 8, Name: "Voucher ANU 50K", Description: "Voucher Restoran", CategoryID: 2, Price: 50000, Stocks: 50, Sold: 12, SubCategoryID: 6},
		{ID: 9, Name: "Voucher ANU 100K", Description: "Voucher Restoran", CategoryID: 2, Price: 100000, Stocks: 50, Sold: 43, SubCategoryID: 6},
		{ID: 10, Name: "Tagihan PLN 100K", Description: "Tagihan Listrik", CategoryID: 3, Price: 100000, Stocks: int(math.Inf(1)), SubCategoryID: 7},
		{ID: 11, Name: "Tagihan PLN 200K", Description: "Tagihan Listrik", CategoryID: 3, Price: 200000, Stocks: int(math.Inf(1)), SubCategoryID: 7},
		{ID: 12, Name: "Tagihan PLN 300K", Description: "Tagihan Listrik", CategoryID: 3, Price: 300000, Stocks: int(math.Inf(1)), SubCategoryID: 7},
		{ID: 14, Name: "Tagihan PLN 400K", Description: "Tagihan Listrik", CategoryID: 3, Price: 400000, Stocks: int(math.Inf(1)), SubCategoryID: 7},
	}
	//var products = productRepo.Product{ID: 14, Name: "Tagihan PLN 400K", Description: "Tagihan Listrik", CategoryID: 3, Price: 400000, Stocks: int(math.Inf(1)), SubCategoryID: 7}

	db.Create(&products)

}

// @title PPOB Service
// @version 1.0
// @description Dokumentasi Swagger e Slurr
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:1234
// @BasePath /
// @schemes http
func main() {
	getConfig := config.GetConfig()

	configPayment := payment.ConfigMidtrans{
		SERVER_KEY: getConfig.SERVER_KEY,
	}
	configPayment.SetupGlobalMidtransConfig()
	payment.InitializeSnapClient()

	configdb := mysql.ConfigDB{
		DB_Username: getConfig.DB_USERNAME,
		DB_Password: getConfig.DB_PASSWORD,
		DB_Host:     getConfig.DB_HOST,
		DB_Port:     getConfig.DB_PORT,
		DB_Database: getConfig.DB_NAME,
	}

	gmail := email.SmtpConfig{
		CONFIG_SMTP_HOST:       getConfig.CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:       getConfig.CONFIG_SMTP_PORT,
		CONFIG_SMTP_AUTH_EMAIL: getConfig.CONFIG_SMTP_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD:   getConfig.CONFIG_AUTH_PASSWORD,
		CONFIG_SENDER_NAME:     getConfig.CONFIG_SENDER_NAME,
	}

	dialer := email.NewGmailConfig(gmail)

	db := configdb.InitialDb()
	dbMigrate(db)

	jwt := _middleware.ConfigJWT{
		SecretJWT:       getConfig.JWT_SECRET,
		ExpiresDuration: int64(getConfig.JWT_EXPIRED),
	}

	configCache := cache.ConfigRedis{
		DB_Host: getConfig.REDIS_URL,
		DB_Port: "6379",
	}
	conn := configCache.InitRedis()

	s3Config := storagedriver.MinioService{
		Host:     getConfig.STORAGE_URL,
		Username: getConfig.STORAGE_ID,
		Secret:   getConfig.STORAGE_SECRET,
	}

	s3 := s3Config.NewClient()

	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	e.Pre(middleware.RemoveTrailingSlash())
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// User
	userIRepo := userRepo.NewRepository(db, conn)
	userIUsecase := userUsecase.NewUseCase(userIRepo, &jwt, *dialer)
	userIDelivery := userDelivery.NewUserDelivery(userIUsecase)

	// Product
	productIrepo := productRepo.NewRepository(db, conn)
	productIUsecase := productUsecase.NewUseCase(productIrepo, s3, getConfig.STORAGE_URL)
	productIdelivery := productDelivery.NewProductDelivery(productIUsecase)

	// VoucherDelivery
	voucherIRepo := voucherRepo.NewRepository(db)
	voucherIUseCase := voucherUsecase.NewUseCase(voucherIRepo)
	voucherIDelivery := voucherDelivery.NewProductDelivery(voucherIUseCase)

	// Transaction
	txIrepo := txRepository.NewRepository(db)
	txIUseCase := txUsecase.NewUseCase(txIrepo, &configPayment, *dialer)
	transactionDelivery := txDelivery.NewTransactionDelivery(txIUseCase)

	routesInit := routes.RouteControllerList{
		UserDelivery:        *userIDelivery,
		ProductDelivery:     *productIdelivery,
		VoucherDelivery:     *voucherIDelivery,
		TransactionDelivery: *transactionDelivery,
		JWTConfig:           jwt.Init(),
	}

	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(":1234"))
}
