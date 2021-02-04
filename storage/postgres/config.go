package postgres

import (
	"fmt"
	"os"
)

// Config holds Postgres configuration and is used to build the DSN string,
// used to communicate with Postgres.
type Config struct {
	// The Postgres host that the server willl connect with. Default for
	// "postgres", to work with Kubernetes and Docker-compose.
	Host string

	// Postgres port that the server will conect with. Default for ":5432".
	Port string

	// Name of the database used by the server. By default, "titanic_dev" for dev
	// env, "titanic_prod" for prod env and "titanic_test" for test env. If you
	// decide to use a different name, make sure to update the script located at
	// "scripts/prepare_postgres.sh".
	Database string

	// User used to connect to the Postgres instance. "postgres" by default.
	User string

	// Password used to connect to the Postgres instance.
	Password string

	// SSLMode used to connect to the Postgres instance. Since the Postgres
	// instance is not accessible outside of the Kubernetes cluster, I'm using
	// "disable" for all envs.
	SSLMode string
}

// BuildConfig builds a valid Postgres config based on given environment.
// It has default configs, however their values might be overrided by
// envvars.
func BuildConfig(env string) Config {
	var cfg Config
	switch env {
	case "test":
		cfg = buildTestConfig()

	case "prod":
		cfg = buildProdConfig()

	default:
		cfg = buildDevConfig()
	}

	if os.Getenv("DB_HOST") != "" {
		cfg.Host = os.Getenv("DB_HOST")
	}

	if os.Getenv("DB_PORT") != "" {
		cfg.Port = os.Getenv("DB_PORT")
	}

	if os.Getenv("DB_SSLMODE") != "" {
		cfg.SSLMode = os.Getenv("DB_SSLMODE")
	}

	if os.Getenv("DB_USER") != "" {
		cfg.User = os.Getenv("DB_USER")
	}

	if os.Getenv("DB_PASSWORD") != "" {
		cfg.Password = os.Getenv("DB_PASSWORD")
	}

	return cfg
}

func buildDevConfig() Config {
	return Config{
		Host:     "postgres",
		Database: "titanic_dev",
		Port:     "5432",
		Password: "passwd",
		SSLMode:  "disable",
		User:     "postgres",
	}
}

func buildTestConfig() Config {
	return Config{
		Host:     "localhost",
		Database: "titanic_test",
		Port:     "5432",
		Password: "passwd",
		SSLMode:  "disable",
		User:     "postgres",
	}
}

func buildProdConfig() Config {
	return Config{
		Host:     "postgres",
		Database: "titanic_prod",
		Port:     "5432",
		Password: "passwd",
		SSLMode:  "disable",
		User:     "postgres",
	}
}

// BuildDSN builds a dsn string used to connect to Postgres.
func (c *Config) BuildDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
		c.SSLMode,
	)
}
