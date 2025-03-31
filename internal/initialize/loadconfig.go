package initialize

import (
	"fmt"
	"go_ecommerce/global"

	"github.com/spf13/viper"
)
func LoadConfig(){

	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config file
	viper.SetConfigName("local") // tên file
	viper.SetConfigType("yaml") 

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration: %w \n", err))
	}

	// read server configuration
	
	fmt.Println("Server port:: ", viper.GetInt("server.port"))
	fmt.Println("Server security.jwt.key:: ", viper.GetString("security.jwt.key"))

	// configure stucture
	
	if err = viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Failed to unmarshal configuration: %v \n", err)
	}
}