package env

import (
	"errors"
	"nexus/utils"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	BLOCKCHAIN_MINING_DIFFICULTY = 0
	NETWORK_PORT                 = 0
	NETWORK_HOST                 = ""
)

var ErrLoadingEnvVars = errors.New("error while loading environment variables")

func init() {
	err := godotenv.Load()
	utils.HandleException(err)

	BLOCKCHAIN_MINING_DIFFICULTY, err = strconv.Atoi(os.Getenv("BLOCKCHAIN_MINING_DIFFICULTY"))
	utils.HandleException(err)

	NETWORK_PORT, err = strconv.Atoi(os.Getenv("NETWORK_PORT"))
	utils.HandleException(err)

	NETWORK_HOST = os.Getenv("NETWORK_HOST")
}
