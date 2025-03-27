package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	configuration configurations
	mutex         sync.Once
)

type configurations struct {
	HTTPPort          string
	HTTPServerTimeOut int
	RateLimit         int
	PostgreDBUser     string
	PostgreDBPassword string
	PostgreDBHost     string
	PostgreDBName     string
	PostgreDBPort     string
	SecretKey         string
}

func goDotVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func getIntEnv(key string) int {
	val := goDotVariable(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Error converting %s to int", key)
	}
	return ret
}

func GetConfig() configurations {
	mutex.Do(func() {
		configuration = newConfig()
	})

	return configuration
}

func newConfig() configurations {
	var cfg configurations
	cfg.HTTPPort = goDotVariable("HTTP_PORT")
	cfg.HTTPServerTimeOut = getIntEnv("HTTP_SERVER_TIMEOUT")
	cfg.RateLimit = getIntEnv("RATE_LIMIT")
	cfg.PostgreDBUser = goDotVariable("POSTGRE_DB_USER")
	cfg.PostgreDBPassword = goDotVariable("POSTGRE_DB_PASSWORD")
	cfg.PostgreDBHost = goDotVariable("POSTGRE_DB_HOST")
	cfg.PostgreDBName = goDotVariable("POSTGRE_DB_NAME")
	cfg.PostgreDBPort = goDotVariable("POSTGRE_DB_PORT")
	cfg.SecretKey = goDotVariable("SECRET_KEY")
	return cfg
}
