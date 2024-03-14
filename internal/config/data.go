package config

type AppConfig struct {
	Server ServerConfig `yaml:"Server"`
	DB     DBConfig     `yaml:"DB"`
}

type ServerConfig struct {
	Port uint `yaml:"Port"`
}

type Postgres struct {
	ConnString string `yaml:"ConnString"`
}

type DBConfig struct {
	Postgres Postgres `yaml:"Postgres"`
}
