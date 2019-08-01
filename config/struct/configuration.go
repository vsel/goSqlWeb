package config

// Configuration Links all sub configurations"
type Configuration struct {
	Database   DatabaseConfiguration
	HTTPServer HTTPServerConfiguration
}
