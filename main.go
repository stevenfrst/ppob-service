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
	userDelivery "ppob-service/delivery/user"
	"ppob-service/drivers/mysql"
	cache "ppob-service/drivers/redis"
	productRepo "ppob-service/drivers/repository/product"
	transactionRepo "ppob-service/drivers/repository/transaction"
	userRepo "ppob-service/drivers/repository/user"
	"ppob-service/helpers/encrypt"
	productUsecase "ppob-service/usecase/product"
	userUsecase "ppob-service/usecase/user"
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
	err := db.AutoMigrate(&productRepo.Category{}, &productRepo.Product{}, &userRepo.User{}, &transactionRepo.Transaction{})
	if err != nil {
		log.Fatalln(err)
	}
	var users = []userRepo.User{{ID: 1, Role: "admin", Username: "admin", Password: encryptMigration("admin"), Email: "admin@admin.com", PhoneNumber: "082135166117", Pin: 1234},
		{ID: 2, Role: "user", Username: "kuli", Password: "kuli", Email: "kuli@user.com", PhoneNumber: "0821313123", Pin: 1234},
		{ID: 3, Role: "user", Username: "kuli2", Password: "kuli2", Email: "kuli2@user.com", PhoneNumber: "0831231299", Pin: 1234},
	}
	db.Create(&users)
	var category = []productRepo.Category{{ID: 1, Name: "Pulsa"},
		{ID: 2, Name: "Voucher Restoran"},
		{ID: 3, Name: "Tagihan PLN"},
	}
	db.Create(&category)
	var products = []productRepo.Product{{ID: 1, Name: "Pulsa Telkomsel 10K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 10000, Stocks: 50, Discount: 0, Sold: 6},
		{ID: 2, Name: "Pulsa Telkomsel 20K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 20000, Stocks: 50, Discount: 0, Sold: 3},
		{ID: 3, Name: "Pulsa Telkomsel 25K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 25000, Stocks: 50, Discount: 0, Sold: 10},
		{ID: 4, Name: "Pulsa Telkomsel 50K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 50000, Stocks: 50, Discount: 0, Sold: 99},
		{ID: 5, Name: "Pulsa Telkomsel 100K", Description: "Pulsa Telkomsel", CategoryID: 1, Price: 100000, Stocks: 50, Discount: 0, Sold: 1},
		{ID: 6, Name: "Voucher KFC 100K", Description: "Voucher Restoran", CategoryID: 2, Price: 100000, Stocks: 50, Discount: 0, Sold: 123},
		{ID: 7, Name: "Voucher KFC 200K", Description: "Voucher Restoran", CategoryID: 2, Price: 200000, Stocks: 50, Discount: 0, Sold: 90},
		{ID: 8, Name: "Voucher ANU 50K", Description: "Voucher Restoran", CategoryID: 2, Price: 50000, Stocks: 50, Discount: 0, Sold: 12},
		{ID: 9, Name: "Voucher ANU 100K", Description: "Voucher Restoran", CategoryID: 2, Price: 100000, Stocks: 50, Discount: 0, Sold: 43},
		{ID: 10, Name: "Tagihan PLN 100K", Description: "Voucher Restoran", CategoryID: 3, Price: 100000, Stocks: int(math.Inf(1)), Discount: 0},
		{ID: 11, Name: "Tagihan PLN 200K", Description: "Voucher Restoran", CategoryID: 3, Price: 200000, Stocks: int(math.Inf(1)), Discount: 0},
		{ID: 12, Name: "Tagihan PLN 300K", Description: "Voucher Restoran", CategoryID: 3, Price: 300000, Stocks: int(math.Inf(1)), Discount: 0},
	}
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
	configdb := mysql.ConfigDB{
		DB_Username: getConfig.DB_USERNAME,
		DB_Password: getConfig.DB_PASSWORD,
		DB_Host:     getConfig.DB_HOST,
		DB_Port:     getConfig.DB_PORT,
		DB_Database: getConfig.DB_NAME,
	}
	db := configdb.InitialDb()
	dbMigrate(db)

	jwt := _middleware.ConfigJWT{
		SecretJWT:       getConfig.JWT_SECRET,
		ExpiresDuration: int64(getConfig.JWT_EXPIRED),
	}

	configCache := cache.ConfigRedis{
		DB_Host: "localhost",
		DB_Port: "6379",
	}
	conn := configCache.InitRedis()

	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// User
	userIRepo := userRepo.NewRepository(db)
	userIUsecase := userUsecase.NewUseCase(userIRepo, &jwt)
	userIDelivery := userDelivery.NewUserDelivery(userIUsecase)

	// Product
	productIrepo := productRepo.NewRepository(db, conn)
	productIUsecase := productUsecase.NewUseCase(productIrepo)
	productIdelivery := productDelivery.NewProductDelivery(productIUsecase)

	routesInit := routes.RouteControllerList{
		UserDelivery:    *userIDelivery,
		ProductDelivery: *productIdelivery,
		JWTConfig:       jwt.Init(),
	}

	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(":1234"))
}
