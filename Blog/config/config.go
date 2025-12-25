package config

type Config struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port int
}

type DB struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}
