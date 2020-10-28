// Package config implements the configuration parameters
package config

// ServiceConfig all the configuration properties for the shop service
type ServiceConfig struct {
	Address    string
	Port       int64
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

// GetServiceConfig returns the static config which is overridden by environment variables.
func GetServiceConfig() (*ServiceConfig, error) {
	serviceConfig := ServiceConfig{
		Address:    "localhost",
		Port:       8080,
		DBHost:     "localhost",
		DBPort:     5432,
		DBName:     "article_db",
		DBUser:     "article",
		DBPassword: "article",
	}
	return &serviceConfig, nil
}
