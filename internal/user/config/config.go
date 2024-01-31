package config

type Config struct {
	Database   Database   `yaml:"Database"`
	HttpServer HttpServer `yaml:"HttpServer"`
}

type Database struct {
	Main    DbNode    `yaml:"Main"`
	Replica DbNode    `yaml:"Replica"`
	Redis   RedisNode `yaml:"Redis"`
}

type DbNode struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Port     int    `yaml:"Port"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
	SslMode  string `yaml:"SslMode"`
}

type RedisNode struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

type HttpServer struct {
	Port      int `yaml:"Port"`
	AdminPort int `yaml:"AdminPort"`
}
