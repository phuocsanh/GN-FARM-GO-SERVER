package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct{
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct{
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main(){
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
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Failed to unmarshal configuration: %v \n", err)
	}
	fmt.Println("Port configurations:: ", config.Server.Port)

	for _,db := range config.Databases {
		fmt.Printf("Database:: User: %s, Password: %s, Host: %s \n", db.User, db.Password, db.Host)
	}
}