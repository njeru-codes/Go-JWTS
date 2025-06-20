package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_DSN string
	JWT_SECRET   string
}

func checkRequiredEnvs(requiredEnvs []string) {

	missing := []string{}

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			missing = append(missing, env)
		}
	}
	if len(missing) > 0 {
		log.Fatalf("Missing reuired environment variables: %v", missing)
	}
}

func LoadEnv() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Print("warning: no .env detected,(relying on exported env variaavles)")
	}
	requiredEnvs := []string{
		"DATABASE_DSN",
		"JWT_SECRET",
	}

	checkRequiredEnvs(requiredEnvs)
	return &Config{
		DATABASE_DSN: os.Getenv("DATABASE_DSN"),
		JWT_SECRET:   os.Getenv("JWT_SECRET"),
	}
}
