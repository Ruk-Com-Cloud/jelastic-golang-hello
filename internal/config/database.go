package config

import "fmt"

func (d *DatabaseConfig) GetDSN() string {
	if d.URL != "" {
		return d.URL
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d.Host, d.User, d.Password, d.DBName, d.Port, d.SSLMode)
}

func (d *DatabaseConfig) IsURLConfigured() bool {
	return d.URL != ""
}

func (d *DatabaseConfig) GetConnectionInfo() map[string]string {
	if d.IsURLConfigured() {
		return map[string]string{
			"type": "url",
			"url":  d.URL,
		}
	}

	return map[string]string{
		"type":     "components",
		"host":     d.Host,
		"port":     d.Port,
		"user":     d.User,
		"database": d.DBName,
		"sslmode":  d.SSLMode,
	}
}