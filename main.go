package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ppob-service/app/config"
	_middleware "ppob-service/app/middleware"
	"ppob-service/app/routes"
	userDelivery "ppob-service/delivery/user"
	"ppob-service/drivers/mysql"
	productRepo "ppob-service/drivers/repository/product"
	transactionRepo "ppob-service/drivers/repository/transaction"
	userRepo "ppob-service/drivers/repository/user"
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

func dbMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&productRepo.Category{}, &productRepo.Product{}, &userRepo.User{}, &transactionRepo.Transaction{})
	if err != nil {
		log.Fatalln(err)
	}
	var users = []userRepo.User{{ID: 1, Role: "admin", Username: "admin", Password: "admin", Email: "admin@admin.com", PhoneNumber: "082135166117", Pin: 1234},
		{ID: 2, Role: "user", Username: "kuli", Password: "kuli", Email: "kuli@user.com", PhoneNumber: "0821313123", Pin: 1234},
		{ID: 3, Role: "user", Username: "kuli2", Password: "kuli2", Email: "kuli2@user.com", PhoneNumber: "0831231299", Pin: 1234},
	}
	db.Create(&users)
	var category = []productRepo.Category{{ID: 1, Name: "Voucher Belanja"},
		{ID: 2, Name: "Voucher Game"},
	}
	db.Create(&category)
	var products = []productRepo.Product{{ID: 1, Name: "Voucher Belanja Gesek 50K", Description: "Haha hihi", CategoryID: 1, Price: 50000, Stocks: 50, Discount: 0},
		{ID: 2, Name: "Voucher Belanja Gesek 100K", Description: "Haha hihi", CategoryID: 1, Price: 100000, Stocks: 50, Discount: 0},
	}
	db.Create(&products)

}

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

	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	e.Pre(middleware.RemoveTrailingSlash())

	// User
	userIRepo := userRepo.NewRepository(db)
	userIUsecase := userUsecase.NewUseCase(userIRepo, &jwt)
	userIDelivery := userDelivery.NewUserDelivery(userIUsecase)

	routesInit := routes.RouteControllerList{
		UserDelivery: *userIDelivery,
		JWTConfig:    jwt.Init(),
	}

	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(":1234"))
}
