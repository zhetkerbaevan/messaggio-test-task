package config

import "os"

type Config struct {
	DBUser     			string
	DBPassword 			string
	DBHost     			string
	DBPort     			string
	DBName     			string
	KafkaURL 			string
	KafkaClientCert 	string
	KafkaClientCertKey 	string
	KafkaTrustedCert 	string
	Port 				string
}

var Envs = initConfig() //Singleton

func initConfig() Config {
	return Config{
		DBUser: getEnv("DB_USER", "u1r5302n6qj669"),
		DBPassword: getEnv("DB_PASSWORD", "p303b395a6b7dcbf8f71fe72f15f07e36e817358a3c4062a1b189993a1ea67ec7"),
		DBHost: getEnv("DB_HOST", "c724r43q8jp5nk.cluster-czz5s0kz4scl.eu-west-1.rds.amazonaws.com"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBName: getEnv("DB_NAME", "ddegrb2pgcrnb4"),
		KafkaURL : getEnv("KAFKA_URL", "ec2-54-77-1-249.eu-west-1.compute.amazonaws.com:9096"),
		KafkaClientCert : getEnv("KAFKA_CLIENT_CERT", ""),
		KafkaClientCertKey : getEnv("KAFKA_CLIENT_CERT_KEY", ""),
		KafkaTrustedCert : getEnv("KAFKA_TRUSTED_CERT", ""),
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