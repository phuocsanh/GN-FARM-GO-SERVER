package initialize

import (
	"database/sql"
	"fmt"
	"go_ecommerce/global"
	"go_ecommerce/internal/common"
	"time"

	"gorm.io/gen"
)

// // checkErrorPanic logs the error and panics if the error is not nil
// func CheckErrorPanic(err error, errString string) {
// 	if err != nil {
// 		global.Logger.Error(errString, zap.Error(err))
// 		panic(err)
// 	}
// }


func InitMysqlC() {
	m := global.Config.Mysql
	// Build the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.Dbname)
	// Open the connection
	db, err := sql.Open("mysql", dsn)
	common.CheckErrorPanic(err, "Failed to initialize MySQL")

	global.Logger.Info("MySQLC Initialized Successfully")
	global.Mdbc = db


	// Set connection pool settings
	// A pool is a set of pre-maintained connections that improve performance.
	// genTableDAOC()

	// Run migrations
	// migrateTables()
}

// setPool sets the MySQL connection pool settings
func setPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	common.CheckErrorPanic(err, "Failed to get SQL DB from GORM")

	// Set connection pool configurations
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}

func genTableDAOC()  {
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

// migrateTables runs database migrations
// func migrateTablesC() {
// err:= global.Mdb.AutoMigrate(
// 	// &po.User{},
// 	// &po.Role{},
// 	&model.GoCrmUserV2{},
//  )
//  if err != nil {
// 	fmt.Println("Migration table failed", err)
//  }
// }
