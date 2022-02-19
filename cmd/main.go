package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/majoo_test/soal_1/configs"
	"github.com/majoo_test/soal_1/pkg/handler"
	"github.com/majoo_test/soal_1/pkg/repository"
	"github.com/majoo_test/soal_1/pkg/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	os.Setenv("TZ", configs.TZ)
}

var (
	db              = initDB()
	userRepo        = repository.NewUserRepository(db)
	merchantRepo    = repository.NewMerchantRepository(db)
	outletRepo      = repository.NewOutletRepository(db)
	transactionRepo = repository.NewTransactionRepository(db)
	authService     = service.NewAuthService(userRepo)
	merchantService = service.NewMerchantService(merchantRepo, transactionRepo)
	outletService   = service.NewOutletService(outletRepo, merchantRepo, transactionRepo)
	authHandler     = handler.NewAuthHandler(authService)
	merchantHandler = handler.NewMerchantHandler(authService, merchantService)
	outletHandler   = handler.NewOutletHandler(authService, outletService)
)

func main() {
	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Authorization, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Post("/login", authHandler.Login)

	app.Get("/merchant/omzet", merchantHandler.GetOmzetReport)

	app.Get("/outlet/omzet", outletHandler.GetOmzetReport)

	app.Listen(":8100")
}

func initDB() *gorm.DB {

	dsn := configs.GetMySqlDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
