package config

import "os"

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	KafkaBrokers string
}

var Envs = initConfig() //Singleton

func initConfig() Config {
	return Config{
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5433"),
		DBName: getEnv("DB_NAME", "messaggio_test_task"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
	}
}

func getEnv(key, fallback string) string {
	//Look for env variable by key
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback //Return default
}