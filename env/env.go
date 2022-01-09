package env

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	BLOCKCHAIN_MINING_DIFFICULTY = 0
)

var ErrLoadingEnvVars = errors.New("error while loading environment variables")

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(ErrLoadingEnvVars)
	}

	BLOCKCHAIN_MINING_DIFFICULTY, err = strconv.Atoi(os.Getenv("BLOCKCHAIN_MINING_DIFFICULTY"))
	if err != nil {
		log.Fatalln(err)
	}
}
