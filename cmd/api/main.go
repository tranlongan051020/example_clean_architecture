package main

import (
	"log"

	"clean-architecture/config"
	"clean-architecture/db"
	"clean-architecture/docs"
	userService "clean-architecture/internal/user/service"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	loadConfig()
	initSwagger()
	// db := loadMySQLDB()
	db := loadPosgresDB()

	r := gin.Default()

	userService.StartUserService(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Fatalf("Server run error: %v", err)
	}
}

func loadConfig() {
	cfgFile, err := config.LoadConfig("dev")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	if err := config.ParseConfig(cfgFile); err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
}

func loadPosgresDB() *gorm.DB {
	db, err := db.PgConnection()
	if err != nil {
		log.Fatalf("Postgresql init error: %s", err)
	} else {
		log.Println("Postgres connected")
	}
	return db
}

func loadMySQLDB() *gorm.DB {
	db, err := db.MysqlConnection()
	if err != nil {
		log.Fatalf("MySQL init error: %s", err)
	} else {
		log.Println("MySQL connected")
	}
	return db
}

func initSwagger() {
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
