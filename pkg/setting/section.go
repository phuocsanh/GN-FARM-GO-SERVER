package setting
type Config struct {
	Server ServerSetting `mapstructure:"server"`	
	Mysql MySQLSetting `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
	JWT JWTSetting `mapstructure:"jwt"`
}

// JWT settings
type JWTSetting struct {
	TOKEN_HOUR_LIFESPAN string `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	API_SECRET_KEY string `mapstructure:"API_SECRET_KEY"`
	JWT_EXPIRATION string `mapstructure:"JWT_EXPRIRATION"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
type MySQLSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"maxIdleConns"`
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`

}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}

