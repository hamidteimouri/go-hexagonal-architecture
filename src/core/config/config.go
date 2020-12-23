package config

import "os"

func Env(key string) string {
	return os.Getenv(key)
}

func IsDebugMode() bool {
	if Env("APP_DEBUG") == "true" {
		return true
	}
	return false
}

func GetEnvironment() string {
	if Env("APP_ENV") == "production" {
		return "production"
	}
	return "development"
}
