package initialize

import (
	"fmt"
	"go_ecommerce/global"
	"go_ecommerce/internal/common"
	"go_ecommerce/internal/model"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// // checkErrorPanic logs the error and panics if the error is not nil
// func CheckErrorPanic(err error, errString string) {
// 	if err != nil {
// 		global.Logger.Error(errString, zap.Error(err))
// 		panic(err)
// 	}
// }


func InitMysql() {
	m := global.Config.Mysql
	// Build the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.Dbname)
	// Open the connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	common.CheckErrorPanic(err, "Failed to initialize MySQL")

	global.Logger.Info("MySQL Initialized Successfully")
	global.Mdb = db


	// Set connection pool settings
	// A pool is a set of pre-maintained connections that improve performance.
	setPool()
	// genTableDAO()

	// Run migrations for products tables if needed
	migrateProductTables()
}

// setPool sets the MySQL connection pool settings
func setPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	common.CheckErrorPanic(err, "Failed to get SQL DB from GORM")

	// Set connection pool configurations
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}

func genTableDAO()  {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
	  })
	
	  // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	  g.UseDB(global.Mdb) // reuse your gorm db
		// g.GenerateAllTable()
		 g.GenerateModel("go_crm_user")

		// fmt.Println("go_crm_user_v2",genner )
	  	// Generate basic type-safe DAO API for struct `model.User` following conventions
		//   g.ApplyBasic(model.User{})
	
	  	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
		//   g.ApplyInterface(func(Querier){}, model.User{}, model.Company{})
	
	 	// Generate the code
	  g.Execute()
}

// migrateProductTables runs database migrations for product related tables
func migrateProductTables() {
	err := global.Mdb.AutoMigrate(
		&model.ProductModel{},
		&model.MushroomModel{},
		&model.VegetableModel{},
		&model.BonsaiModel{},
	)
	if err != nil {
		fmt.Println("Migration products tables failed", err)
		global.Logger.Error("Migration products tables failed", zap.Error(err))
	} else {
		global.Logger.Info("Migration products tables successful")
	}
}
