package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	App      AppConfig
	HTTP     HTTPConfig
	GRPC     GRPCConfig
	Metrics  MetricsConfig
	Postgres PostgresConfig
	Kafka    KafkaConfig
	Auth     AuthConfig
}
type AppConfig struct {
	Env  string
	Name string //service name
}

type HTTPConfig struct {
	Port string
}

type GRPCConfig struct {
	Port string
}

type MetricsConfig struct {
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type KafkaConfig struct {
	Brokers  string
	ClientID string
}

type AuthConfig struct {
	JWTSecret        string
	DummyTokenSecret string //test token
}

func Load() (Config, error) {
	var cfg Config

	//App
	cfg.App.Env = getEnvOrDefault("APP_ENV", "local")
	cfg.App.Name = getEnvOrDefault("APP_NAME", "lastmile")

	//HTTP
	cfg.HTTP.Port = getEnvOrDefault("HTTP_PORT", "8080")

	//Postgres
	cfg.Postgres.Host = getEnv("POSTGRES_HOST")
	cfg.Postgres.Port = getEnvOrDefault("POSTGRES_PORT", "5433")
	cfg.Postgres.User = getEnv("POSTGRES_USER")
	cfg.Postgres.Password = getEnv("POSTGRES_PASSWORD")
	cfg.Postgres.DBName = getEnv("POSTGRES_DB")
	cfg.Postgres.SSLMode = getEnvOrDefault("POSTGRES_SSLMODE", "disable")

	//Kafka
	cfg.Kafka.Brokers = getEnvOrDefault("KAFKA_BROKERS", "")
	cfg.Kafka.ClientID = getEnvOrDefault("KAFKA_CLIENT_ID", cfg.App.Name)

	//Auth
	cfg.Auth.JWTSecret = getEnvOrDefault("JWT_SECRET", "")
	cfg.Auth.DummyTokenSecret = getEnvOrDefault("DUMMY_TOKEN_SECRET", "dummy-secret")

	//validation
	if err := validate(cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// building dsn for posgres
func (p PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.DBName,
		p.SSLMode,
	)
}

// building kafka
func (k KafkaConfig) BrokersSlice() []string {
	if k.Brokers == "" {
		return []string{}
	}
	s := strings.Split(k.Brokers, ",")

	// trim spaces
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}

	return s
}
func (a AppConfig) IsProd() bool {
	// to do
	return false
}

// helper for config
func getEnvOrDefault(key, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	return v
}
func getEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return ""
	}
	return v
}
func validate(cfg Config) error {
	if cfg.Postgres.Host == "" {
		return fmt.Errorf("POSTGRES_HOST is required")
	}
	if cfg.Postgres.User == "" {
		return fmt.Errorf("POSTGRES_USER is required")
	}
	if cfg.Postgres.DBName == "" {
		return fmt.Errorf("POSTGRES_DB is required")
	}

	// to do add : JWT_SECRET, Kafka brokers ...
	return nil
}
