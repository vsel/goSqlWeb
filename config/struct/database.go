package config

// DatabaseConfiguration is server variables
type DatabaseConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSL      string
}
