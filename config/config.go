package config

type Config struct {
	Http     Http     `yaml:"http" mapstructure:"http"`
	Database Database `yaml:"database" mapstructure:"database"`
	Secret   Security `yaml:"secret" mapstructure:"secret"`
}

type Http struct {
	Port int `yaml:"port" mapstructure:"port"`
}

type Security struct {
	Pepper string `yaml:"pepper" mapstructure:"pepper"`
}

type Database struct {
	Driver                  string `yaml:"driver", mapstructure:"driver"`
	Host                    string `yaml:"host", mapstructure:"host"`
	Port                    string `yaml:"port" mapstructure:"port"`
	User                    string `yaml:"user", mapstructure:"user"`
	Password                string `yaml:"password", mapstructure:"password"`
	DbName                  string `yaml:"dbName", mapstructure:"dbName"`
	CorsEnabled             bool   `yaml:"CorsEnabled" mapstructure:"CorsEnabled"`
	MigrationsPath          string `yaml:"migrationsPath" mapstructure:"migrationsPath"`
	ConnectionMaxLifetimeMs int    `yaml:"connectionMaxLifetimeMs" mapstructure:"connectionMaxLifetimeMs"`
	MaxIdleConnections      int    `yaml:"maxIdleConnections" mapstructure:"maxIdleConnections"`
	MaxOpenConnection       int    `yaml:"maxOpenConnection" mapstructure:"maxOpenConnection"`
}
