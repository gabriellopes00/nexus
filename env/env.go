package env

import (
	"errors"
	"nexus/utils"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// block chain mining difficulty
	BLOCKCHAIN_DIFFICULTY = 0

	// network port
	NETWORK_PORT = 0

	// network host
	NETWORK_HOST = ""

	// path to database files
	DATABASE_PATH = ""

	// database file name
	DATABASE_FILE = ""
)

// ErrLoadingEnvVars represents an error on loading env vars
var ErrLoadingEnvVars = errors.New("error while loading environment variables")

func init() {
	err := godotenv.Load()
	utils.HandleException(err)

	BLOCKCHAIN_DIFFICULTY, err = strconv.Atoi(os.Getenv("BLOCKCHAIN_DIFFICULTY"))
	utils.HandleException(err)

	NETWORK_PORT, err = strconv.Atoi(os.Getenv("NETWORK_PORT"))
	utils.HandleException(err)

	NETWORK_HOST = os.Getenv("NETWORK_HOST")
	DATABASE_PATH = os.Getenv("DATABASE_PATH")
	DATABASE_FILE = os.Getenv("DATABASE_FILE")
}
