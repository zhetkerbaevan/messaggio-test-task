package config

import "os"

type Config struct {
	DBUrl     			string
	KafkaBrokers 		string
	Port 				string
}

var Envs = initConfig() //Singleton

func initConfig() Config {
	return Config{
		DBUrl: getEnv("DB_URL", "postgresql://postgres:AlBUVkGGqvapNtjYsJmnFYCAeVkTpEOG@monorail.proxy.rlwy.net:29387/railway"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
		Port: getEnv("PORT", ":9006"),
	}
}

func getEnv(key, fallback string) string {
	//Look for env variable by key
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback //Return default
}