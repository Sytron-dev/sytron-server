package database

import (
	"os"

	"github.com/joho/godotenv"
)

// Load the .env file first before fe==getting environment variables
var _ = godotenv.Load(".env")

// Database constants from the .env file
var MONGODB_CONNECTION_STRING = os.Getenv("MONGODB_CONNECTION_STRING")
var DATABASE_NAME = os.Getenv("PRIMARY_DATABASE")
var CONFIGS_DATABASE_NAME = os.Getenv("CONFIGS_DATABASE")
