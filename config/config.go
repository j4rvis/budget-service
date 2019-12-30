package config

// Config ...
type Config struct {
	ListenAddr     string
	LogLevel       string
	DbURL          string
	DbUser         string
	DbUserPassword string
	DbDatabase     string
}

// New ...
func New() *Config {
	config := Config{
		ListenAddr:     "localhost:8080",
		LogLevel:       "debug",
		DbURL:          "localhost:5432",
		DbUser:         "postgres",
		DbUserPassword: "mysecretpassword",
		DbDatabase:     "budgetdb",
	}

	return &config
}
