package config

type AppConfig struct {
	Server    ServerConfig `yaml:"Server"`
	DB        DBConfig     `yaml:"DB"`
	LocalHash LocalHash    `yaml:"LocalHash"`
}

type ServerConfig struct {
	Port uint `yaml:"Port"`
}

type Postgres struct {
	ConnString                 string `yaml:"ConnString"`
	MaxCons                    uint16 `yaml:"MaxCons"`
	MinCons                    uint16 `yaml:"MinCons"`
	MaxConnLifetimeInMinutes   uint32 `yaml:"MaxConnLifetimeInMinutes"`
	MaxConnIdleTimeInMinutes   uint32 `yaml:"MaxConnIdleTimeInMinutes"`
	HealthCheckPeriodInSeconds uint32 `yaml:"HealthCheckPeriodInSeconds"`
	ConnectTimeoutInSeconds    uint32 `yaml:"ConnectTimeoutInSeconds"`
}

type DBConfig struct {
	Postgres Postgres `yaml:"Postgres"`
}

type LocalHash struct {
	Salt string `yaml:"Salt"`
}
